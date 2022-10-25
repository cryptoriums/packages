// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package head

import (
	"context"
	"math"
	"sync"
	"time"

	"github.com/bluele/gcache"
	"github.com/cryptoriums/packages/constants"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

const (
	ComponentName = "trackerHead"
	defaultDelay  = 10 * time.Second
)

type Contract interface {
	NetID() int64
	Addr() common.Address
}

type TrackerHead struct {
	logger      log.Logger
	ctx         context.Context
	stop        context.CancelFunc
	client      ethereum.ChainReader
	cacheHeadTX gcache.Cache
	dstChan     chan *types.Block

	reorgWaitPeriod time.Duration
}

func New(
	ctx context.Context,
	logger log.Logger,
	client ethereum.ChainReader,
	reorgWaitPeriod time.Duration,
) (*TrackerHead, chan *types.Block, error) {
	logger = log.With(logger, "component", ComponentName)
	ctx, stop := context.WithCancel(ctx)

	dstChan := make(chan *types.Block)
	return &TrackerHead{
		client:          client,
		ctx:             ctx,
		stop:            stop,
		logger:          logger,
		dstChan:         dstChan,
		reorgWaitPeriod: reorgWaitPeriod,
		// To be on the safe side create the cache few times bigger then the expected block count during the reorg wait.
		cacheHeadTX: gcache.New(int(math.Max(100, 3*constants.BlocksPerSecond*reorgWaitPeriod.Seconds()))).LRU().Build(),
	}, dstChan, nil
}
func (self *TrackerHead) Start() error {
	level.Info(self.logger).Log("msg", "starting", "reorgWaitPeriod", self.reorgWaitPeriod)

	src, subs := self.waitSubscribe()
	defer func() {
		if subs != nil {
			subs.Unsubscribe()
		}
	}()

	go self.listen(src)

	for {
		select {
		case <-self.ctx.Done():
			return nil
		case err := <-subs.Err():
			level.Error(self.logger).Log("msg", "subscription failed will try to resubscribe", "err", err)
			src, subs = self.waitSubscribe()
			self.listen(src)
		}
	}
}

func (self *TrackerHead) listen(src chan *types.Header) {
	level.Info(self.logger).Log("msg", "starting new subs listener")

	for {
		select {
		case <-self.ctx.Done():
			level.Info(self.logger).Log("msg", "subscription listener canceled")
			return
		case event := <-src:
			logger := log.With(self.logger, "block", event.Number)

			level.Debug(logger).Log("msg", "new block")
			if self.reorgWaitPeriod == 0 {
				go func(event *types.Header) {
					ctx, cncl := context.WithTimeout(self.ctx, time.Minute)
					defer cncl()
					block, err := self.client.BlockByNumber(ctx, event.Number)
					if err != nil {
						level.Error(logger).Log("msg", "getting full block from head hash", "err", err, "num", event.Number, "hash", event.Hash())
						return
					}
					select {
					case self.dstChan <- block:
					case <-self.ctx.Done():
						return
					}
				}(event)
				continue
			}

			// Doesn't need a cancelation ctx as
			// the BlockByNumber query returns the same block for reorg events.
			go func(event *types.Header) {
				waitForReorg := time.NewTicker(self.reorgWaitPeriod)
				defer waitForReorg.Stop()

				select {
				case <-waitForReorg.C:
				case <-self.ctx.Done():
					return
				}

				ctx, cncl := context.WithTimeout(self.ctx, 2*time.Minute)
				defer cncl()

				// Duplicate event numbers will still return the same block when using this query.
				block, err := self.client.BlockByNumber(ctx, event.Number)
				if err != nil {
					level.Error(logger).Log("msg", "getting full block from head hash", "err", err, "num", event.Number, "hash", event.Hash())
					return
				}

				_, err = self.cacheHeadTX.Get(block.Hash().Hex())
				if err != gcache.KeyNotFoundError {
					level.Debug(logger).Log("msg", "skipping head block that has already been processed")
					return
				}

				if err := self.cacheHeadTX.Set(block.Hash().Hex(), true); err != nil {
					level.Error(logger).Log("msg", "adding head block tx cache", "err", err)
				}

				select {
				case self.dstChan <- block:
					return
				case <-self.ctx.Done():
					return
				}

			}(event)
		}
	}
}

func (self *TrackerHead) Stop() {
	self.stop()
}

func (self *TrackerHead) waitSubscribe() (chan *types.Header, event.Subscription) {
	output := make(chan *types.Header)

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

		subs, err := self.client.SubscribeNewHead(self.ctx, output)
		if err != nil {
			level.Error(self.logger).Log("msg", "subscription to head failed", "err", err)
			continue
		}
		level.Info(self.logger).Log("msg", "subscription created", "eventName", "head")
		return output, subs
	}
}

type NoopSubs struct{}

func (self *NoopSubs) Unsubscribe()      {}
func (self *NoopSubs) Err() <-chan error { return nil }
