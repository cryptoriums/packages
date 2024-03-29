// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package events

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"sync"
	"time"

	"github.com/bluele/gcache"
	"github.com/cryptoriums/packages/client"
	"github.com/cryptoriums/packages/client/events"
	"github.com/cryptoriums/packages/constants"
	"github.com/cryptoriums/packages/logging"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/pkg/errors"
)

const (
	ComponentName = "trackerEvent"
)

type Config struct {
	LogLevel string
}

type TrackerEvents struct {
	logger       log.Logger
	ctx          context.Context
	stop         context.CancelFunc
	client       client.EthClient
	mtx          sync.Mutex
	cacheSentTXs gcache.Cache
	addrs        []common.Address
	fromBlock    *big.Int
	eventQuery   [][]interface{}
	dstChan      chan types.Log

	reorgWaitPeriod  time.Duration
	reorgWaitPending map[string]context.CancelFunc
}

func New(
	ctx context.Context,
	logger log.Logger,
	cfg Config,
	client client.EthClient,
	addrs []common.Address,
	fromBlock uint64,
	lookBack time.Duration,
	eventQuery [][]interface{},
	reorgWaitPeriod time.Duration,
) (*TrackerEvents, chan types.Log, error) {
	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, nil, errors.Wrap(err, "apply filter logger")
	}
	if fromBlock != 0 && lookBack != 0 {
		return nil, nil, errors.New("only one needs to be set either fromBlock or lookBack")
	}

	var _fromBlock *big.Int
	if fromBlock != 0 {
		_fromBlock = big.NewInt(int64(fromBlock))
	} else {
		headerNow, err := client.HeaderByNumber(ctx, nil)
		if err != nil {
			return nil, nil, errors.Wrap(err, "get latest eth block header")
		}
		_fromBlock = headerNow.Number
	}

	if lookBack != 0 {
		blockNums := constants.BlocksPerMinute * lookBack.Minutes()

		headerNow, err := client.HeaderByNumber(ctx, nil)
		if err != nil {
			return nil, nil, errors.Wrap(err, "get latest eth block header")
		}
		_fromBlock = big.NewInt(headerNow.Number.Int64() - int64(blockNums))
	}

	logger = log.With(logger, "component", ComponentName)
	ctx, stop := context.WithCancel(ctx)

	dstChan := make(chan types.Log)
	return &TrackerEvents{
		client:           client,
		ctx:              ctx,
		stop:             stop,
		logger:           logger,
		addrs:            addrs,
		fromBlock:        _fromBlock,
		eventQuery:       eventQuery,
		dstChan:          dstChan,
		reorgWaitPeriod:  reorgWaitPeriod,
		reorgWaitPending: make(map[string]context.CancelFunc),
		// To be on the safe side create the cache few times bigger then the expected block count during the reorg wait.
		cacheSentTXs: gcache.New(int(math.Max(100, 100*constants.BlocksPerSecond*reorgWaitPeriod.Seconds()))).LRU().Build(),
	}, dstChan, nil
}
func (self *TrackerEvents) Start() error {
	level.Info(self.logger).Log("msg", "starting",
		"monitorAddrs", fmt.Sprintf("%+v", self.addrs),
		"fromBlock", self.fromBlock,
		"eventQuery", fmt.Sprintf("%+v", self.eventQuery),
		"reorgWaitPeriod", self.reorgWaitPeriod,
	)

	ctxS, cnclS := context.WithTimeout(self.ctx, 5*time.Minute)
	defer cnclS()
	err := self.sendHistoricalLogs(ctxS)
	if err != nil {
		return errors.Wrap(err, "sending historical logs")
	}

	src, subs, err := self.waitSubscribe()
	if err != nil {
		return errors.Wrap(err, "creating subs, this should never happen")
	}
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

			ctxS, cnclS = context.WithTimeout(self.ctx, 5*time.Minute)
			defer cnclS()

			err = self.sendHistoricalLogs(ctxS)
			if err != nil {
				level.Error(self.logger).Log("msg", "sending historical logs")
			}

			src, subs, err = self.waitSubscribe()
			cncl()
			if err != nil {
				return errors.Wrap(err, "creating subs, this should never happen")
			}
			ctx, cncl = context.WithCancel(self.ctx)
			go self.listen(ctx, src)
		}
	}
}

func (self *TrackerEvents) HistoricalLogs(ctx context.Context, fromBlock int64) ([]types.Log, error) {
	q, err := events.CreateFilterQuery(self.addrs, self.eventQuery, big.NewInt(fromBlock), nil)
	if err != nil {
		return nil, err
	}

	return self.client.FilterLogs(ctx, *q)
}

func (self *TrackerEvents) sendHistoricalLogs(ctx context.Context) error {
	q, err := self.createFilterQuery()
	if err != nil {
		return errors.Wrap(err, "creating filter query")
	}

	logs, err := self.client.FilterLogs(ctx, *q)
	if err != nil {
		return errors.Wrap(err, "getting historical logs")
	}
	for _, log := range logs {
		if events.IsCachedForReorg(self.logger, self.cacheSentTXs, log) {
			level.Info(self.logger).Log("msg", "skipping event that has already been sent", "id", events.HashForReorg(log))
			continue
		}

		if err := events.CacheForReorg(self.logger, self.cacheSentTXs, log); err != nil {
			level.Error(self.logger).Log("msg", "adding tx event cache", "err", err)
		}

		select {
		case self.dstChan <- log:
		case <-self.ctx.Done():
			return nil
		}
	}
	return nil
}

func (self *TrackerEvents) listen(ctx context.Context, src chan types.Log) {
	level.Info(self.logger).Log("msg", "starting new subs listener")

	for {
		select {
		case <-ctx.Done():
			level.Info(self.logger).Log("msg", "subscription listener canceled")
			return
		case event := <-src:
			hash := events.HashForReorg(event)
			level.Debug(self.logger).Log("msg", "new event received", "hash", hash)

			if event.Removed {
				self.cancelPending(hash)
				level.Debug(self.logger).Log("msg", "canceling event due to reorg", "hash", hash)
				continue
			}
			ctx, cncl := context.WithCancel(self.ctx)
			self.addPending(hash, cncl)

			go func(ctxReorg context.Context, event types.Log, hash string) {
				if self.reorgWaitPeriod > 0 {
					waitReorg := time.NewTicker(self.reorgWaitPeriod)
					defer waitReorg.Stop()

					select {
					case <-waitReorg.C:
					case <-ctxReorg.Done():
						level.Debug(self.logger).Log("msg", "canceled due to reorg", "hash", hash)
						return
					}
				}
				// With short reorg wait it is possible to try and send the same TX twice so this check mitigates that.

				if events.IsCachedForReorg(self.logger, self.cacheSentTXs, event) {
					level.Info(self.logger).Log("msg", "skipping event that has already been sent", "id", hash)
					return
				}
				if err := events.CacheForReorg(self.logger, self.cacheSentTXs, event); err != nil {
					level.Error(self.logger).Log("msg", "adding tx event cache", "err", err)
				}

				self.cancelPending(hash) // Cleanup the ctx.
				level.Debug(self.logger).Log("msg", "sending event", "hash", hash)
				select {
				case self.dstChan <- event:
					// In case of a subs error this is used to pick up from the last block that was logged.
					self.mtx.Lock()
					self.fromBlock = big.NewInt(int64(event.BlockNumber))
					self.mtx.Unlock()
					return
				case <-self.ctx.Done():
					return
				}

			}(ctx, event, hash)
		}
	}
}

func (self *TrackerEvents) Stop() {
	self.stop()
}

func (self *TrackerEvents) waitSubscribe() (chan types.Log, event.Subscription, error) {
	ticker := time.NewTicker(1)
	defer ticker.Stop()
	var resetTicker sync.Once

	for {
		select {
		case <-self.ctx.Done():
			return nil, &NoopSubs{}, nil // To avoid panics in the caller.
		case <-ticker.C:
			resetTicker.Do(func() { ticker.Reset(time.Second) })
		}

		q, err := self.createFilterQuery()
		if err != nil {
			return nil, nil, errors.Wrap(err, "creating filter query")
		}

		src := make(chan types.Log)

		subs, err := self.client.SubscribeFilterLogs(self.ctx, *q, src)
		if err != nil {
			level.Error(self.logger).Log("msg", "subscription to events failed", "err", err)
			continue
		}

		level.Info(self.logger).Log("msg", "subscription created")
		return src, subs, nil
	}
}

func (self *TrackerEvents) createFilterQuery() (*ethereum.FilterQuery, error) {
	self.mtx.Lock()
	defer self.mtx.Unlock()
	q, err := events.CreateFilterQuery(self.addrs, self.eventQuery, self.fromBlock, nil)
	if err != nil {
		return nil, err
	}
	defer level.Debug(self.logger).Log("msg", "query created", "params", fmt.Sprintf("%+v", q))
	return q, nil
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
