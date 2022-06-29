// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package head

import (
	"context"
	"fmt"

	"github.com/bluele/gcache"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/pkg/errors"
)

type HeadSubscriber interface {
	SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error)
}

// NewHeadSubscriberWithRedundancy creates a HeadSubscriber that can use multiple backends and it ensures that the same logs is never sent twice.
func NewHeadSubscriber(logger log.Logger, headSubscribers []HeadSubscriber) HeadSubscriber {
	return &HeadSubscriberWithRedundancy{
		logger:          logger,
		headSubscribers: headSubscribers,
		err:             make(chan error),
	}
}

// NewHeadSubscriberWithRedundancy creates a HeadSubscriber that can use multiple backends and it ensures that the same logs is never sent twice.
func NewHeadSubscriberWithRedundancy(logger log.Logger, headSubscribers []HeadSubscriber) HeadSubscriber {
	return &HeadSubscriberWithRedundancy{
		logger:          logger,
		headSubscribers: headSubscribers,
		err:             make(chan error),
	}
}

type HeadSubscriberWithRedundancy struct {
	logger          log.Logger
	err             chan error
	headSubscribers []HeadSubscriber
}

func IsCached(logger log.Logger, cache gcache.Cache, header *types.Header) bool {
	hash := HashFromAllFields(header)
	_, err := cache.Get(hash)

	if err != gcache.KeyNotFoundError {
		level.Debug(logger).Log("msg", "header is cached", "hash", hash, "err", err)
		return true
	}
	return false
}

func Cache(logger log.Logger, cache gcache.Cache, header *types.Header) error {
	hash := HashFromAllFields(header)
	level.Debug(logger).Log("msg", "caching header", "hash", hash)
	return cache.Set(hash, true)
}

func (self *HeadSubscriberWithRedundancy) SubscribeNewHead(ctx context.Context, chDst chan<- *types.Header) (ethereum.Subscription, error) {
	chSrc := make(chan *types.Header)

	var subs []ethereum.Subscription
	var errSrc []<-chan error
	for _, headSubscriber := range self.headSubscribers {
		sub, err := headSubscriber.SubscribeNewHead(ctx, chSrc)
		if err != nil {
			return nil, errors.Wrap(err, "creating SubscribeFilterLogs subscription")
		}
		subs = append(subs, sub)
		errSrc = append(errSrc, sub.Err())
	}

	logger := log.With(self.logger, "subscriptionID", fmt.Sprintf("%v", chDst))
	multiSub := NewMultiSubscription(ctx, logger, subs, chSrc, chDst, errSrc)
	return multiSub, nil
}

func NewMultiSubscription(
	ctx context.Context,
	logger log.Logger,
	subs []ethereum.Subscription,
	chSrc chan *types.Header,
	chDst chan<- *types.Header,
	errSrc []<-chan error,
) *MultiSubscription {
	ctx, cncl := context.WithCancel(ctx)

	sub := &MultiSubscription{
		ctx:        ctx,
		cncl:       cncl,
		errDst:     make(chan error),
		subs:       subs,
		errSrc:     errSrc,
		cacheStore: gcache.New(100).LRU().Build(),
	}

	go func(chSrc chan *types.Header, chDst chan<- *types.Header) {
		for {
			select {
			case header := <-chSrc:
				level.Debug(logger).Log("msg", "new header", "block", header.Number)
				if IsCached(logger, sub.cacheStore, header) {
					continue
				}
				select {
				case <-sub.ctx.Done():
					return
				case chDst <- header:
					level.Debug(logger).Log("msg", "header sent", "block", header.Number)
					err := Cache(logger, sub.cacheStore, header)
					if err != nil {
						level.Error(logger).Log("msg", "setting cache", "err", err)
					}
				}
			case <-sub.ctx.Done():
				return
			}
		}
	}(chSrc, chDst)

	for _, err := range errSrc {
		go func(err <-chan error) {
			select {
			case <-ctx.Done():
				sub.Unsubscribe()
				return
			case errI := <-err:
				sub.errDst <- errI
				sub.Unsubscribe()
			}
		}(err)
	}

	return sub
}

type MultiSubscription struct {
	ctx        context.Context
	cncl       context.CancelFunc
	errDst     chan error
	errSrc     []<-chan error
	subs       []ethereum.Subscription
	cacheStore gcache.Cache
}

func (self *MultiSubscription) Unsubscribe() {
	for _, sub := range self.subs {
		sub.Unsubscribe()
	}
	self.cncl()
}

func (self *MultiSubscription) Err() <-chan error {
	return self.errDst
}

func HashFromAllFields(header *types.Header) string {
	return "Number:" + header.Number.String() + "-MixDigest:" + header.MixDigest.Hex() + "-ParentHash:" + header.ParentHash.Hex() + "-TxHash:" + header.TxHash.Hex() + "-UncleHash:" + header.UncleHash.Hex() + "-ReceiptHash:" + header.ReceiptHash.Hex() + "-Root:" + header.Root.Hex()
}
