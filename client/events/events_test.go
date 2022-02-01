// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package events

import (
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/cryptoriums/packages/logging"
	"github.com/cryptoriums/packages/testing/contracts/simple"
	"github.com/cryptoriums/packages/testutil"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

// TestLogsWithRedundancy ensures that events are deduplicated properly when using multiple backends.
//
//
// Here is an example how to use with node urls.
//
// infura, err := ethclient.Dial("wss://infura.io/......")
// alchemy, err := ethclient.Dial("wss://alchemyapi.io/......")
//
// client := NewEthClientWithFiltererRedundancy(ctx, logger, []*ethclient.Client{infura, alchemy})
// client.FilterLogs(ctx, query)
// client.SubscribeFilterLogs(ctx, query, ch)
// .
func TestLogsWithRedundancy(t *testing.T) {
	logger := logging.NewLogger()
	logger, err := logging.ApplyFilter("debug", logger)
	testutil.Ok(t, err)

	ctx := context.Background()

	abi, err := abi.JSON(strings.NewReader(simple.SimpleStorageABI))
	testutil.Ok(t, err)

	sk, err := crypto.GenerateKey()
	testutil.Ok(t, err)

	transactOpts, err := bind.NewKeyedTransactorWithChainID(sk, big.NewInt(1337))
	testutil.Ok(t, err)

	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   nil,
		Topics:    [][]common.Hash{{abi.Events["StorageSetA"].ID}},
	}

	// All backends have the same logs.
	{
		var (
			logFilterers []ethereum.LogFilterer
			backends     []*backends.SimulatedBackend
		)
		for i := 0; i < 4; i++ {
			backend := testutil.GetSimBackend(t, sk)
			_, _, contract, err := simple.DeploySimpleStorage(transactOpts, backend)
			testutil.Ok(t, err)
			_, err = contract.SetA(transactOpts, "1111")
			testutil.Ok(t, err)
			backend.Commit()
			logFilterers = append(logFilterers, backend)
			backends = append(backends, backend)
		}

		filterer := NewLogFiltererWithRedundancy(logger, logFilterers)

		logsExp, err := backends[0].FilterLogs(ctx, query)
		testutil.Ok(t, err)
		logsAct, err := filterer.FilterLogs(ctx, query)
		testutil.Ok(t, err)
		testutil.Equals(t, logsExp, logsAct)
	}

	// 2 backends have different logs.
	{
		backend1 := testutil.GetSimBackend(t, sk)
		_, _, contract1, err := simple.DeploySimpleStorage(transactOpts, backend1)
		testutil.Ok(t, err)

		backend2 := testutil.GetSimBackend(t, sk)
		_, _, contract2, err := simple.DeploySimpleStorage(transactOpts, backend2)
		testutil.Ok(t, err)

		filterer := NewLogFiltererWithRedundancy(logger, []ethereum.LogFilterer{backend1, backend2})

		_, err = contract1.SetA(transactOpts, "aaaa")
		testutil.Ok(t, err)
		backend1.Commit()

		logsExp, err := backend1.FilterLogs(ctx, query)
		testutil.Ok(t, err)

		_, err = contract2.SetA(transactOpts, "bbbb")
		testutil.Ok(t, err)
		backend2.Commit()
		logsExp1, err := backend2.FilterLogs(ctx, query)
		testutil.Ok(t, err)

		logsExp = append(logsExp, logsExp1...)

		logsAct, err := filterer.FilterLogs(ctx, query)
		testutil.Ok(t, err)

		testutil.Equals(t, logsExp, logsAct)
	}

	// 2 backends one has one extra log.
	{
		backend1 := testutil.GetSimBackend(t, sk)
		_, _, contract1, err := simple.DeploySimpleStorage(transactOpts, backend1)
		testutil.Ok(t, err)

		backend2 := testutil.GetSimBackend(t, sk)
		_, _, contract2, err := simple.DeploySimpleStorage(transactOpts, backend2)
		testutil.Ok(t, err)

		filterer := NewLogFiltererWithRedundancy(logger, []ethereum.LogFilterer{backend1, backend2})

		_, err = contract1.SetA(transactOpts, "aaaa")
		testutil.Ok(t, err)
		backend1.Commit()

		_, err = contract2.SetA(transactOpts, "aaaa")
		testutil.Ok(t, err)
		backend2.Commit()

		_, err = contract2.SetA(transactOpts, "bbbb")
		testutil.Ok(t, err)
		backend2.Commit()
		logsExp, err := backend2.FilterLogs(ctx, query)
		testutil.Ok(t, err)

		logsAct, err := filterer.FilterLogs(ctx, query)
		testutil.Ok(t, err)

		testutil.Equals(t, logsExp, logsAct)
	}

	// Subscription with 2 backends both send the same logs.
	{
		backend1 := testutil.GetSimBackend(t, sk)
		_, _, contract1, err := simple.DeploySimpleStorage(transactOpts, backend1)
		testutil.Ok(t, err)

		backend2 := testutil.GetSimBackend(t, sk)
		_, _, contract2, err := simple.DeploySimpleStorage(transactOpts, backend2)
		testutil.Ok(t, err)

		filterer := NewLogFiltererWithRedundancy(logger, []ethereum.LogFilterer{backend1, backend2})

		ch := make(chan types.Log)
		subs, err := filterer.SubscribeFilterLogs(ctx, query, ch)
		testutil.Ok(t, err)
		defer subs.Unsubscribe()

		_, err = contract1.SetA(transactOpts, "2222")
		testutil.Ok(t, err)
		backend1.Commit()

		_, err = contract2.SetA(transactOpts, "2222")
		testutil.Ok(t, err)
		backend2.Commit()

		logsExp, err := backend1.FilterLogs(ctx, query)
		testutil.Ok(t, err)

		var logsAct []types.Log
		logsAct = append(logsAct, <-ch)

		testutil.Equals(t, logsExp, logsAct)

		select {
		case log := <-ch:
			t.Fatalf("there is an extra log:%+v", log)
		default:
		}
	}

	// Subscription with 2 backends one is missing a log.
	{
		backend1 := testutil.GetSimBackend(t, sk)
		_, _, contract1, err := simple.DeploySimpleStorage(transactOpts, backend1)
		testutil.Ok(t, err)

		backend2 := testutil.GetSimBackend(t, sk)
		_, _, contract2, err := simple.DeploySimpleStorage(transactOpts, backend2)
		testutil.Ok(t, err)

		filterer := NewLogFiltererWithRedundancy(logger, []ethereum.LogFilterer{backend1, backend2})

		ch := make(chan types.Log)
		subs, err := filterer.SubscribeFilterLogs(ctx, query, ch)
		testutil.Ok(t, err)
		defer subs.Unsubscribe()

		_, err = contract1.SetA(transactOpts, "2222")
		testutil.Ok(t, err)
		backend1.Commit()
		_, err = contract2.SetA(transactOpts, "2222")
		testutil.Ok(t, err)
		backend2.Commit()
		_, err = contract2.SetA(transactOpts, "3333")
		testutil.Ok(t, err)
		backend2.Commit()

		logsExp, err := backend2.FilterLogs(ctx, query)
		testutil.Ok(t, err)

		var logsAct []types.Log

		logsAct = append(logsAct, <-ch)
		logsAct = append(logsAct, <-ch)

		testutil.Equals(t, logsExp, logsAct)

		select {
		case log := <-ch:
			t.Fatalf("there is an extra log:%+v", log)
		default:
		}
	}

}

func TestMultipleSubsDeduplication(t *testing.T) {
	logger := logging.NewLogger()
	logger, err := logging.ApplyFilter("debug", logger)
	logging.ExitOnError(logger, err)
	ctx := context.Background()

	abi, err := abi.JSON(strings.NewReader(simple.SimpleStorageABI))
	testutil.Ok(t, err)

	sk, err := crypto.GenerateKey()
	testutil.Ok(t, err)

	transactOpts, err := bind.NewKeyedTransactorWithChainID(sk, big.NewInt(1337))
	testutil.Ok(t, err)

	queryA := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   nil,
		Topics:    [][]common.Hash{{abi.Events["StorageSetA"].ID}},
	}

	queryB := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   nil,
		Topics:    [][]common.Hash{{abi.Events["StorageSetB"].ID}},
	}

	// Different subscripitons shouldn't send to the same channel.
	{

		backend := testutil.GetSimBackend(t, sk)
		_, _, contract, err := simple.DeploySimpleStorage(transactOpts, backend)
		testutil.Ok(t, err)

		filterer := NewLogFiltererWithRedundancy(logger, []ethereum.LogFilterer{backend})

		chA := make(chan types.Log)
		subs, err := filterer.SubscribeFilterLogs(ctx, queryA, chA)
		testutil.Ok(t, err)
		defer subs.Unsubscribe()

		chB := make(chan types.Log)
		subs, err = filterer.SubscribeFilterLogs(ctx, queryB, chB)
		testutil.Ok(t, err)
		defer subs.Unsubscribe()

		_, err = contract.SetA(transactOpts, "aaaa1")
		testutil.Ok(t, err)
		backend.Commit()
		_, err = contract.SetA(transactOpts, "aaaa2")
		testutil.Ok(t, err)
		backend.Commit()

		_, err = contract.SetB(transactOpts, "bbbb1")
		testutil.Ok(t, err)
		backend.Commit()
		_, err = contract.SetB(transactOpts, "bbbb2")
		testutil.Ok(t, err)
		backend.Commit()
		_, err = contract.SetB(transactOpts, "bbbb3")
		testutil.Ok(t, err)
		backend.Commit()

		logsExpA, err := filterer.FilterLogs(ctx, queryA)
		testutil.Ok(t, err)

		logsActA := []types.Log{<-chA, <-chA}

		testutil.Equals(t, logsExpA, logsActA)
		select {
		case log := <-chA:
			t.Fatalf("there is an extra log:%+v", log)
		default:
		}
	}

	// Ensure that  the subs.Err() func receives an error even when one LogFilterer returns an error.
	{

		filterer := NewLogFiltererWithRedundancy(logger, []ethereum.LogFilterer{NewBackendSimulateErr(), testutil.GetSimBackend(t, nil)})

		subs, err := filterer.SubscribeFilterLogs(ctx, ethereum.FilterQuery{}, nil)
		testutil.Ok(t, err)

		time.Sleep(time.Second)
		select {
		case <-subs.Err():
		default:
			t.Fatalf("no error was received")
		}

	}
}

func NewBackendSimulateErr() *BackendSimulateErr {
	e := make(chan error)
	go func() {
		e <- errors.New("xxx")
	}()
	return &BackendSimulateErr{
		err: e,
	}
}

type BackendSimulateErr struct {
	err <-chan error
}

func (self *BackendSimulateErr) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	return nil, nil
}

func (self *BackendSimulateErr) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, chDst chan<- types.Log) (ethereum.Subscription, error) {
	return self, nil
}

func (self *BackendSimulateErr) Err() <-chan error {
	return self.err
}

func (self *BackendSimulateErr) Unsubscribe() {

}
