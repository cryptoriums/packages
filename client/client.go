// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package client

import (
	"context"
	"math/big"
	"strings"
	"sync"
	"time"

	ethereum_t "github.com/cryptoriums/packages/ethereum"
	"github.com/cryptoriums/packages/events"
	"github.com/cryptoriums/packages/logging"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

const (
	defaultRetry  = 2 * time.Second
	ComponentName = "ethClient"
)

type Config struct {
	LogLevel string
}

type ClientWithRetry struct {
	logger log.Logger
	nodes  []string
	*ethclient.Client
	ethClients []*ethclient.Client
	rpcClients []*rpc.Client
	netID      int64
	ethereum.LogFilterer
}

func NewClientWithRetry(ctx context.Context, logger log.Logger, cfg Config, envVars map[string]string) (ethereum_t.EthClientRpc, error) {
	nodes, err := parseNodes(envVars)
	if err != nil {
		return nil, err
	}
	ethClients, rpcClients, netID, err := ethereum_t.NewClients(ctx, logger, nodes)
	if err != nil {
		return nil, err
	}

	logger, err = logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}

	var filterers []ethereum.LogFilterer
	for _, filterer := range ethClients {
		filterers = append(filterers, filterer)
	}

	return &ClientWithRetry{
		logger:      log.With(logger, "component", ComponentName),
		nodes:       nodes,
		netID:       netID,
		Client:      ethClients[0], // For the functions that don't offer redundancy just call the first client.
		rpcClients:  rpcClients,
		ethClients:  ethClients,
		LogFilterer: events.NewLogFiltererWithRedundancy(logger, filterers),
	}, nil
}

func (self *ClientWithRetry) NetworkID() int64 {
	return self.netID
}

func (self *ClientWithRetry) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	return self.LogFilterer.FilterLogs(ctx, query)
}

func (self *ClientWithRetry) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return self.LogFilterer.SubscribeFilterLogs(ctx, query, ch)
}

func (self *ClientWithRetry) Close() {
	for _, client := range self.ethClients {
		client.Close()
	}
	for _, client := range self.rpcClients {
		client.Close()
	}
}

func (self *ClientWithRetry) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	var merr error
	for i, rpcClient := range self.rpcClients {
		if err := rpcClient.CallContext(ctx, result, method, args...); err == nil {
			return nil
		} else {
			level.Error(self.logger).Log("msg", "rpc call", "node", self.nodes[i], "err", err)
			merr = multierror.Append(merr, err)
		}
	}
	return merr
}

func (self *ClientWithRetry) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	ticker := time.NewTicker(time.Millisecond)
	var resetTickerOnce sync.Once
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			resetTickerOnce.Do(func() { ticker.Reset(time.Second) })
		}
		for i, ethClient := range self.ethClients {
			ctxRetry, cncl := context.WithTimeout(ctx, defaultRetry)
			defer cncl()
			result, err := ethClient.CallContract(ctxRetry, call, blockNumber)
			if err != nil {
				level.Error(self.logger).Log("msg", "CallContract", "node", self.nodes[i], "err", err)
				continue
			}
			return result, nil
		}
	}
}

func (self *ClientWithRetry) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	ticker := time.NewTicker(time.Millisecond)
	var resetTickerOnce sync.Once
	defer ticker.Stop()
	for {
		ctxRetry, cncl := context.WithTimeout(ctx, defaultRetry)
		defer cncl()
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		case <-ticker.C:
			resetTickerOnce.Do(func() { ticker.Reset(time.Second) })
		}
		for i, ethClient := range self.ethClients {
			result, err := ethClient.PendingNonceAt(ctxRetry, account)
			if err != nil {
				level.Error(self.logger).Log("msg", "PendingNonceAt", "node", self.nodes[i], "err", err)
				continue
			}
			return result, nil
		}
	}
}

func (self *ClientWithRetry) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	ticker := time.NewTicker(time.Millisecond)
	var resetTickerOnce sync.Once
	defer ticker.Stop()
	for {
		ctxRetry, cncl := context.WithTimeout(ctx, 30*time.Second)
		defer cncl()
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			resetTickerOnce.Do(func() { ticker.Reset(time.Second) })
		}
		for i, ethClient := range self.ethClients {
			result, err := ethClient.BlockByNumber(ctxRetry, number)
			if err != nil {
				level.Error(self.logger).Log("msg", "BlockByNumber", "node", self.nodes[i], "err", err)
				continue
			}
			return result, nil
		}
	}
}

func (self *ClientWithRetry) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	ticker := time.NewTicker(time.Millisecond)
	var resetTickerOnce sync.Once
	defer ticker.Stop()
	for {
		ctxRetry, cncl := context.WithTimeout(ctx, defaultRetry)
		defer cncl()
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			resetTickerOnce.Do(func() { ticker.Reset(time.Second) })
		}

		for i, ethClient := range self.ethClients {
			result, err := ethClient.SuggestGasTipCap(ctxRetry)
			if err != nil {
				level.Error(self.logger).Log("msg", "SuggestGasTipCap", "node", self.nodes[i], "err", err)
				continue
			}
			return result, nil
		}
	}
}

func (self *ClientWithRetry) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	ticker := time.NewTicker(time.Millisecond)
	var resetTickerOnce sync.Once
	defer ticker.Stop()
	for {
		ctxRetry, cncl := context.WithTimeout(ctx, defaultRetry)
		defer cncl()
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			resetTickerOnce.Do(func() { ticker.Reset(time.Second) })
		}

		for i, ethClient := range self.ethClients {
			result, err := ethClient.SuggestGasPrice(ctxRetry)
			if err != nil {
				level.Error(self.logger).Log("msg", "SuggestGasPrice", "node", self.nodes[i], "err", err)
				continue
			}
			return result, nil
		}
	}
}

func (self *ClientWithRetry) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	ticker := time.NewTicker(time.Millisecond)
	var resetTickerOnce sync.Once
	defer ticker.Stop()
	for {
		ctxRetry, cncl := context.WithTimeout(ctx, defaultRetry)
		defer cncl()
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			resetTickerOnce.Do(func() { ticker.Reset(time.Second) })
		}

		for i, ethClient := range self.ethClients {
			result, err := ethClient.PendingCodeAt(ctxRetry, account)
			if err != nil {
				level.Error(self.logger).Log("msg", "PendingCodeAt", "node", self.nodes[i], "err", err)
				continue
			}
			return result, nil
		}
	}
}

func (self *ClientWithRetry) EstimateGas(ctx context.Context, call ethereum.CallMsg) (gas uint64, err error) {
	ticker := time.NewTicker(time.Millisecond)
	var resetTickerOnce sync.Once
	defer ticker.Stop()
	for {
		ctxRetry, cncl := context.WithTimeout(ctx, defaultRetry)
		defer cncl()
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		case <-ticker.C:
			resetTickerOnce.Do(func() { ticker.Reset(time.Second) })
		}

		for i, ethClient := range self.ethClients {
			result, err := ethClient.EstimateGas(ctxRetry, call)
			if err != nil {
				level.Error(self.logger).Log("msg", "EstimateGas", "node", self.nodes[i], "err", err)
				continue
			}
			return result, nil
		}
	}
}

func (self *ClientWithRetry) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	ticker := time.NewTicker(time.Millisecond)
	var resetTickerOnce sync.Once
	defer ticker.Stop()
	for {
		ctxRetry, cncl := context.WithTimeout(ctx, defaultRetry)
		defer cncl()
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			resetTickerOnce.Do(func() { ticker.Reset(time.Second) })
		}

		for i, ethClient := range self.ethClients {
			err := ethClient.SendTransaction(ctxRetry, tx)
			if err != nil {
				level.Error(self.logger).Log("msg", "SendTransaction", "node", self.nodes[i], "err", err)
				continue
			}
			return nil
		}
	}
}

func (self *ClientWithRetry) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	ticker := time.NewTicker(time.Millisecond)
	var resetTickerOnce sync.Once
	defer ticker.Stop()
	for {
		ctxRetry, cncl := context.WithTimeout(ctx, defaultRetry)
		defer cncl()
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			resetTickerOnce.Do(func() { ticker.Reset(time.Second) })
		}
		for i, ethClient := range self.ethClients {
			result, err := ethClient.BalanceAt(ctxRetry, account, blockNumber)
			if err != nil {
				level.Error(self.logger).Log("msg", "BalanceAt", "node", self.nodes[i], "err", err)
				continue
			}
			return result, nil
		}
	}
}

func (self *ClientWithRetry) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	ticker := time.NewTicker(time.Millisecond)
	var resetTickerOnce sync.Once
	defer ticker.Stop()
	for {
		ctxRetry, cncl := context.WithTimeout(ctx, defaultRetry)
		defer cncl()
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			resetTickerOnce.Do(func() { ticker.Reset(time.Second) })
		}

		for i, ethClient := range self.ethClients {
			result, err := ethClient.HeaderByNumber(ctxRetry, number)
			if err != nil {
				level.Error(self.logger).Log("msg", "HeaderByNumber", "node", self.nodes[i], "err", err)
				continue
			}
			return result, nil
		}
	}
}

func parseNodes(envVars map[string]string) ([]string, error) {
	nodeURLs, ok := envVars[ethereum_t.NodeURLEnvName]
	if !ok {
		return nil, errors.Errorf("the env var doesn't exist NodeURLEnvName:%v", ethereum_t.NodeURLEnvName)
	}
	nodes := strings.Split(nodeURLs, ",")
	if len(nodes) == 0 {
		return nil, errors.New("the env var doesn't contain any node urls")
	}
	return nodes, nil
}
