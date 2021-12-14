// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package events

import (
	"context"
	"strconv"

	"github.com/bluele/gcache"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/pkg/errors"
)

type EthClientWithFiltererRedundancy struct {
	*ethclient.Client
	bind.ContractFilterer
}

func NewEthClientWithFiltererRedundancy(logger log.Logger, clients []*ethclient.Client) *EthClientWithFiltererRedundancy {
	var filterers []ethereum.LogFilterer
	for _, client := range clients {
		filterers = append(filterers, client)
	}
	return &EthClientWithFiltererRedundancy{
		Client:           clients[0], // For the functions that don't offer redundancy just call the first client.
		ContractFilterer: NewLogFiltererWithRedundancy(logger, filterers),
	}
}

func (self *EthClientWithFiltererRedundancy) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	return self.ContractFilterer.FilterLogs(ctx, query)
}

func (self *EthClientWithFiltererRedundancy) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, chDst chan<- types.Log) (ethereum.Subscription, error) {
	return self.ContractFilterer.SubscribeFilterLogs(ctx, query, chDst)
}

// NewLogFiltererWithRedundancy creates a ContractFilterer that can use multiple backends and it ensures that the same logs is never sent twice.
func NewLogFiltererWithRedundancy(logger log.Logger, logFilterers []ethereum.LogFilterer) ethereum.LogFilterer {
	return &LogFiltererWithRedundancy{
		logger:       logger,
		logFilterers: logFilterers,
		err:          make(chan error),
		cacheStore:   gcache.New(1000).LRU().Build(),
	}
}

type LogFiltererWithRedundancy struct {
	logger       log.Logger
	err          chan error
	multiSubs    []*MultiSubscription
	cacheStore   gcache.Cache
	logFilterers []ethereum.LogFilterer
}

func (self *LogFiltererWithRedundancy) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	cacheStore := gcache.New(1000).LRU().Build()
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
		if isCached(self.logger, cacheStore, log) {
			continue
		}
		logsDeduped = append(logsDeduped, log)
		err := cache(self.logger, cacheStore, log)
		if err != nil {
			level.Error(self.logger).Log("msg", "caching log entry", "err", err)
		}

		for _, logs := range logsAll {
			if len(logs) < i+1 {
				continue
			}
			if isCached(self.logger, cacheStore, logs[i]) {
				continue
			}
			logsDeduped = append(logsDeduped, logs[i])
			err := cache(self.logger, cacheStore, logs[i])
			if err != nil {
				level.Error(self.logger).Log("msg", "setting cache", "err", err)
			}
		}
	}
	return logsDeduped, nil
}

func isCached(logger log.Logger, cache gcache.Cache, log types.Log) bool {
	hash := HashFromLogAllFields(log)
	_, err := cache.Get(hash)

	if err != gcache.KeyNotFoundError {
		level.Debug(logger).Log("msg", "log is cached", "hash", hash, "err", err)
		return true
	}
	return false
}

func cache(logger log.Logger, cache gcache.Cache, log types.Log) error {
	hash := HashFromLogAllFields(log)
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

	multiSub := NewMultiSubscription(ctx, self.logger, subs, chSrc, chDst, errSrc, self.cacheStore)

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
	cacheStore gcache.Cache,
) *MultiSubscription {
	ctx, cncl := context.WithCancel(ctx)

	sub := &MultiSubscription{
		ctx:    ctx,
		cncl:   cncl,
		errDst: make(chan error),
		subs:   subs,
		errSrc: errSrc,
	}

	go func(chSrc chan types.Log) {
		for {
			select {
			case log := <-chSrc:
				if isCached(logger, cacheStore, log) {
					continue
				}
				select {
				case <-sub.ctx.Done():
					return
				case chDst <- log:
					level.Debug(logger).Log("msg", "event sent", "log", log.TxHash)
					err := cache(logger, cacheStore, log)
					if err != nil {
						level.Error(logger).Log("msg", "setting cache", "err", err)
					}
				}
			case <-sub.ctx.Done():
				return
			}
		}
	}(chSrc)

	for _, err := range errSrc {
		go func(err <-chan error) {
			select {
			case <-ctx.Done():
				return
			case errI := <-err:
				sub.errDst <- errI
			}
		}(err)
	}

	return sub
}

type MultiSubscription struct {
	ctx    context.Context
	cncl   context.CancelFunc
	errDst chan error
	errSrc []<-chan error
	subs   []ethereum.Subscription
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

func HashFromLog(log types.Log) string {
	// Using the topics data will cause a race when more than one TX include a log with the same topics, but it is highly unlikely.
	topicStr := ""
	for _, topic := range log.Topics {
		topicStr += topic.Hex() + ","
	}
	return log.TxHash.Hex() + "-topics:" + topicStr
}

func HashFromLogAllFields(log types.Log) string {
	hash := HashFromLog(log)
	return hash + "-blockHash:" + log.BlockHash.Hex() + "-txIndex:" + log.TxHash.Hex() + "-removed:" + strconv.FormatBool(log.Removed)
}
