// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package events

import (
	"context"
	"math/big"
	"strconv"

	"github.com/bluele/gcache"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/pkg/errors"
)

// NewLogFiltererWithRedundancy creates a ContractFilterer that can use multiple backends and it ensures that the same logs is never sent twice.
func NewLogFiltererWithRedundancy(logger log.Logger, logFilterers []ethereum.LogFilterer) ethereum.LogFilterer {
	return &LogFiltererWithRedundancy{
		logger:       logger,
		logFilterers: logFilterers,
		err:          make(chan error),
	}
}

type LogFiltererWithRedundancy struct {
	logger       log.Logger
	err          chan error
	multiSubs    []*MultiSubscription
	logFilterers []ethereum.LogFilterer
}

func (self *LogFiltererWithRedundancy) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	cacheStore := gcache.New(100).LRU().Build()
	var logsAll [][]types.Log
	for _, logFilterer := range self.logFilterers {
		logs, err := logFilterer.FilterLogs(ctx, query)
		if err != nil {
			return nil, errors.Wrap(err, "getting past logs")
		}
		logsAll = append(logsAll, logs)
	}

	if len(logsAll) == 0 {
		return nil, nil
	}

	biggestArray := logsAll[0]
	biggestArrayIdx := 0
	for i, logs := range logsAll {
		if i > 0 && (len(logsAll[i]) > len(logsAll[i-1])) {
			biggestArray = logs
			biggestArrayIdx = i
		}
	}

	logsAll = append(logsAll[:biggestArrayIdx], logsAll[biggestArrayIdx+1:]...)

	var logsDeduped []types.Log
	for i, log := range biggestArray {
		if IsCached(self.logger, cacheStore, log) {
			continue
		}
		logsDeduped = append(logsDeduped, log)
		err := Cache(self.logger, cacheStore, log)
		if err != nil {
			level.Error(self.logger).Log("msg", "caching log entry", "err", err)
		}

		for _, logs := range logsAll {
			if len(logs) < i+1 {
				continue
			}
			if IsCached(self.logger, cacheStore, logs[i]) {
				continue
			}
			logsDeduped = append(logsDeduped, logs[i])
			err := Cache(self.logger, cacheStore, logs[i])
			if err != nil {
				level.Error(self.logger).Log("msg", "setting Cache", "err", err)
			}
		}
	}
	return logsDeduped, nil
}

func IsCached(logger log.Logger, cache gcache.Cache, log types.Log) bool {
	hash := HashFromFields(log)
	_, err := cache.Get(hash)

	if err != gcache.KeyNotFoundError {
		level.Debug(logger).Log("msg", "log is cached", "hash", hash, "err", err)
		return true
	}
	return false
}

func Cache(logger log.Logger, cache gcache.Cache, log types.Log) error {
	hash := HashFromFields(log)
	level.Debug(logger).Log("msg", "caching log", "hash", hash)
	return cache.Set(hash, true)
}

func (self *LogFiltererWithRedundancy) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, chDst chan<- types.Log) (ethereum.Subscription, error) {
	chSrc := make(chan types.Log)

	var subs []ethereum.Subscription
	var errSrc []<-chan error
	for _, logFilterer := range self.logFilterers {
		sub, err := logFilterer.SubscribeFilterLogs(ctx, query, chSrc)
		if err != nil {
			return nil, errors.Wrap(err, "creating SubscribeFilterLogs subscription")
		}
		subs = append(subs, sub)
		errSrc = append(errSrc, sub.Err())
	}

	multiSub := NewMultiSubscription(ctx, self.logger, subs, chSrc, chDst, errSrc)

	self.multiSubs = append(self.multiSubs, multiSub)

	return multiSub, nil
}

func NewMultiSubscription(
	ctx context.Context,
	logger log.Logger,
	subs []ethereum.Subscription,
	chSrc chan types.Log,
	chDst chan<- types.Log,
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

	go func(chSrc chan types.Log, chDst chan<- types.Log) {
		for {
			select {
			case log := <-chSrc:
				if IsCached(logger, sub.cacheStore, log) {
					continue
				}
				select {
				case <-sub.ctx.Done():
					return
				case chDst <- log:
					level.Debug(logger).Log("msg", "event sent", "log", log.TxHash)
					err := Cache(logger, sub.cacheStore, log)
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

func HashFromFields(log types.Log) string {
	topicStr := ""
	for _, topic := range log.Topics {
		topicStr += topic.Hex() + ","
	}
	return "BlockNum:" + strconv.Itoa(int(log.BlockNumber)) + "-TxHash:" + log.TxHash.Hex() + "-Topics:" + topicStr + "-BlockHash:" + log.BlockHash.Hex() + "-Index:" + strconv.Itoa(int(log.Index)) + "-Removed:" + strconv.FormatBool(log.Removed)
}

func CreateFilterQuery(addrs []common.Address, qI [][]interface{}, fromBlock *big.Int) (*ethereum.FilterQuery, error) {
	topics, err := abi.MakeTopics(qI...)
	if err != nil {
		return nil, err
	}

	q := &ethereum.FilterQuery{
		Addresses: addrs,
		Topics:    topics,
		FromBlock: fromBlock,
	}

	return q, nil
}
