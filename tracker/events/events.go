// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package events

import (
	"context"
	"math"
	"math/big"
	"sync"
	"time"

	"github.com/bluele/gcache"
	ethereum_t "github.com/cryptoriums/packages/ethereum"
	"github.com/cryptoriums/packages/events"
	"github.com/cryptoriums/packages/logging"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/pkg/errors"
)

const (
	ComponentName = "trackerEvent"
	defaultDelay  = 10 * time.Second
)

type Config struct {
	LogLevel string
}

type Contract interface {
	Abi() abi.ABI
	WatchLogs(opts *bind.WatchOpts, name string, query ...[]interface{}) (chan types.Log, event.Subscription, error)
	Addr() common.Address
}

type TrackerEvents struct {
	logger       log.Logger
	ctx          context.Context
	stop         context.CancelFunc
	client       ethereum_t.EthClient
	mtx          sync.Mutex
	cacheSentTXs gcache.Cache
	contract     Contract
	lookBack     time.Duration
	eventName    string
	dstChan      chan types.Log

	reorgWaitPeriod  time.Duration
	reorgWaitPending map[string]context.CancelFunc
}

func New(
	ctx context.Context,
	logger log.Logger,
	cfg Config,
	client ethereum_t.EthClient,
	contract Contract,
	lookBack time.Duration,
	eventName string,
	reorgWaitPeriod time.Duration,
) (*TrackerEvents, chan types.Log, error) {
	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, nil, errors.Wrap(err, "apply filter logger")
	}
	logger = log.With(logger, "component", ComponentName)
	ctx, stop := context.WithCancel(ctx)

	dstChan := make(chan types.Log)
	return &TrackerEvents{
		client:           client,
		ctx:              ctx,
		stop:             stop,
		logger:           logger,
		contract:         contract,
		lookBack:         lookBack,
		eventName:        eventName,
		dstChan:          dstChan,
		reorgWaitPeriod:  reorgWaitPeriod,
		reorgWaitPending: make(map[string]context.CancelFunc),
		// To be on the safe side create the cache 2 times bigger then the expected block count during the reorg wait.
		cacheSentTXs: gcache.New(int(math.Max(50, 2*ethereum_t.BlocksPerSecond*reorgWaitPeriod.Seconds()))).LRU().Build(),
	}, dstChan, nil
}
func (self *TrackerEvents) Start() error {
	level.Info(self.logger).Log("msg", "starting",
		"monitorAddr", self.contract.Addr(),
		"lookBack", self.lookBack,
		"eventName", self.eventName,
		"reorgWaitPeriod", self.reorgWaitPeriod,
	)
	if self.lookBack != 0 {
		ctx, cncl := context.WithTimeout(self.ctx, time.Minute)
		defer cncl()
		blockNums := ethereum_t.BlocksPerMinute * self.lookBack.Minutes()

		headerNow, err := self.client.HeaderByNumber(ctx, nil)
		if err != nil {
			return errors.Wrap(err, "get latest eth block header")
		}
		fromBlock := headerNow.Number.Int64() - int64(blockNums)
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(fromBlock),
			ToBlock:   nil,
			Addresses: []common.Address{self.contract.Addr()},
			Topics:    [][]common.Hash{{self.contract.Abi().Events[self.eventName].ID}},
		}

		logs, err := self.client.FilterLogs(self.ctx, query)
		if err != nil {
			return errors.Wrap(err, "getting historical logs")
		}
		for _, log := range logs {
			select {
			case self.dstChan <- log:
			case <-self.ctx.Done():
				return nil
			}
		}

	}

	// Initial subscription.
	src, subs := self.waitSubscribe()
	defer func() {
		if subs != nil {
			subs.Unsubscribe()
		}
	}()

	ctx, cncl := context.WithCancel(self.ctx)
	go self.listen(ctx, src)

	for {
		select {
		case <-self.ctx.Done():
			cncl()
			return nil
		case err := <-subs.Err():
			level.Error(self.logger).Log("msg", "subscription failed will try to resubscribe", "err", err)
			src, subs = self.waitSubscribe()
			cncl()
			ctx, cncl = context.WithCancel(self.ctx)
			go self.listen(ctx, src)
		}
	}
}

func (self *TrackerEvents) listen(ctx context.Context, src chan types.Log) {
	level.Info(self.logger).Log("msg", "starting new subs listener")

	for {
		select {
		case <-ctx.Done():
			level.Info(self.logger).Log("msg", "subscription listener canceled")
			return
		case event := <-src:
			hash := events.HashFromLog(event)
			level.Debug(self.logger).Log("msg", "new event received", "hash", hash)

			if event.Removed {
				self.cancelPending(hash)
				level.Debug(self.logger).Log("msg", "canceling event due to reorg", "hash", hash)
				continue
			}
			ctx, cncl := context.WithCancel(self.ctx)
			self.addPending(hash, cncl)

			go func(ctxReorg context.Context, event types.Log, hash string) {
				waitReorg := time.NewTicker(self.reorgWaitPeriod)
				defer waitReorg.Stop()

				select {
				case <-waitReorg.C:
					// With short reorg wait it is possible to try and send the same TX twice so this check mitigates that.
					_, err := self.cacheSentTXs.Get(hash)
					if err != gcache.KeyNotFoundError {
						level.Info(self.logger).Log("msg", "skipping event that has already been sent", "id", hash)
						return
					}
					if err := self.cacheSentTXs.Set(hash, true); err != nil {
						level.Error(self.logger).Log("msg", "adding tx event cache", "err", err)
					}

					self.cancelPending(hash) // Cleanup the ctx.
					level.Debug(self.logger).Log("msg", "sending event", "hash", hash)
					select {
					case self.dstChan <- event:
						return
					case <-self.ctx.Done():
						return
					}
				case <-ctxReorg.Done():
					level.Debug(self.logger).Log("msg", "canceled due to reorg", "hash", hash)
					return
				}

			}(ctx, event, hash)
		}
	}
}

func (self *TrackerEvents) Stop() {
	self.stop()
}

func (self *TrackerEvents) waitSubscribe() (chan types.Log, event.Subscription) {
	ticker := time.NewTicker(1)
	defer ticker.Stop()
	var resetTicker sync.Once

	for {
		select {
		case <-self.ctx.Done():
			return nil, &NoopSubs{} // To avoid panics in the caller.
		case <-ticker.C:
			resetTicker.Do(func() { ticker.Reset(defaultDelay) })
		}

		opts := &bind.WatchOpts{
			Context: self.ctx,
		}
		src, subs, err := self.contract.WatchLogs(opts, self.eventName)
		if err != nil {
			level.Error(self.logger).Log("msg", "subscription to events failed", "err", err)
			continue
		}
		level.Info(self.logger).Log("msg", "subscription created", "eventName", self.eventName)
		return src, subs
	}
}

func (self *TrackerEvents) addPending(hash string, cncl context.CancelFunc) {
	self.mtx.Lock()
	defer self.mtx.Unlock()
	self.reorgWaitPending[hash] = cncl
}

func (self *TrackerEvents) cancelPending(hash string) {
	self.mtx.Lock()
	defer self.mtx.Unlock()

	if cncl, ok := self.reorgWaitPending[hash]; ok {
		cncl()
		delete(self.reorgWaitPending, hash)
	}
}

type NoopSubs struct{}

func (self *NoopSubs) Unsubscribe()      {}
func (self *NoopSubs) Err() <-chan error { return nil }
