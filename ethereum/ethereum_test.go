// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package ethereum

import (
	"crypto/ecdsa"
	"math/big"
	"strings"
	"testing"

	"github.com/cryptoriums/packages/logging"
	"github.com/cryptoriums/packages/testing/contracts/simple"
	"github.com/cryptoriums/packages/testutil"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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
func TestLogsWithRedundancy(t *testing.T) {
	logger := logging.NewLogger()
	logging.ApplyFilter("debug", logger)
	ctx := context.Background()

	abi, err := abi.JSON(strings.NewReader(simple.SimpleStorageABI))
	testutil.Ok(t, err)

	sk, err := crypto.GenerateKey()
	testutil.Ok(t, err)

	transactOpts, err := bind.NewKeyedTransactorWithChainID(sk, big.NewInt(1337))

	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   nil,
		Topics:    [][]common.Hash{{abi.Events["StorageSet"].ID}},
	}

	// All backends have the same logs.
	{
		var (
			logFilterers []ethereum.LogFilterer
			contracts    []*simple.SimpleStorage
			backends     []*backends.SimulatedBackend
		)
		for i := 0; i < 4; i++ {
			backend := getSimBackend(t, sk)
			_, _, contract, err := simple.DeploySimpleStorage(transactOpts, backend)
			testutil.Ok(t, err)
			_, err = contract.Set(transactOpts, "1111")
			testutil.Ok(t, err)
			backend.Commit()
			logFilterers = append(logFilterers, backend)
			contracts = append(contracts, contract)
			backends = append(backends, backend)
		}

		filterer := NewLogFiltererWithRedundancy(ctx, logger, logFilterers)

		logsExp, err := backends[0].FilterLogs(ctx, query)
		logsAct, err := filterer.FilterLogs(ctx, query)
		testutil.Ok(t, err)
		testutil.Equals(t, logsExp, logsAct)
	}

	// 2 backends have different logs.
	{
		backend1 := getSimBackend(t, sk)
		_, _, contract1, err := simple.DeploySimpleStorage(transactOpts, backend1)
		testutil.Ok(t, err)

		backend2 := getSimBackend(t, sk)
		_, _, contract2, err := simple.DeploySimpleStorage(transactOpts, backend2)
		testutil.Ok(t, err)

		filterer := NewLogFiltererWithRedundancy(ctx, logger, []ethereum.LogFilterer{backend1, backend2})

		_, err = contract1.Set(transactOpts, "aaaa")
		testutil.Ok(t, err)
		backend1.Commit()

		logsExp, err := backend1.FilterLogs(ctx, query)
		testutil.Ok(t, err)

		_, err = contract2.Set(transactOpts, "bbbb")
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
		backend1 := getSimBackend(t, sk)
		_, _, contract1, err := simple.DeploySimpleStorage(transactOpts, backend1)
		testutil.Ok(t, err)

		backend2 := getSimBackend(t, sk)
		_, _, contract2, err := simple.DeploySimpleStorage(transactOpts, backend2)
		testutil.Ok(t, err)

		filterer := NewLogFiltererWithRedundancy(ctx, logger, []ethereum.LogFilterer{backend1, backend2})

		_, err = contract1.Set(transactOpts, "aaaa")
		testutil.Ok(t, err)
		backend1.Commit()

		_, err = contract2.Set(transactOpts, "aaaa")
		testutil.Ok(t, err)
		backend2.Commit()

		_, err = contract2.Set(transactOpts, "bbbb")
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
		backend1 := getSimBackend(t, sk)
		_, _, contract1, err := simple.DeploySimpleStorage(transactOpts, backend1)
		testutil.Ok(t, err)

		backend2 := getSimBackend(t, sk)
		_, _, contract2, err := simple.DeploySimpleStorage(transactOpts, backend2)
		testutil.Ok(t, err)

		filterer := NewLogFiltererWithRedundancy(ctx, logger, []ethereum.LogFilterer{backend1, backend2})

		ch := make(chan types.Log)
		subs, err := filterer.SubscribeFilterLogs(ctx, query, ch)
		testutil.Ok(t, err)
		defer subs.Unsubscribe()

		_, err = contract1.Set(transactOpts, "2222")
		testutil.Ok(t, err)
		backend1.Commit()

		_, err = contract2.Set(transactOpts, "2222")
		testutil.Ok(t, err)
		backend2.Commit()

		logsExp, err := filterer.FilterLogs(ctx, query)
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
		backend1 := getSimBackend(t, sk)
		_, _, contract1, err := simple.DeploySimpleStorage(transactOpts, backend1)
		testutil.Ok(t, err)

		backend2 := getSimBackend(t, sk)
		_, _, contract2, err := simple.DeploySimpleStorage(transactOpts, backend2)
		testutil.Ok(t, err)

		filterer := NewLogFiltererWithRedundancy(ctx, logger, []ethereum.LogFilterer{backend1, backend2})

		ch := make(chan types.Log)
		subs, err := filterer.SubscribeFilterLogs(ctx, query, ch)
		testutil.Ok(t, err)
		defer subs.Unsubscribe()

		_, err = contract1.Set(transactOpts, "2222")
		testutil.Ok(t, err)
		backend1.Commit()
		_, err = contract2.Set(transactOpts, "2222")
		testutil.Ok(t, err)
		backend2.Commit()
		_, err = contract2.Set(transactOpts, "3333")
		testutil.Ok(t, err)
		backend2.Commit()

		logsExp, err := filterer.FilterLogs(ctx, query)
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

func getSimBackend(t *testing.T, sk *ecdsa.PrivateKey) *backends.SimulatedBackend {

	faucetAddr := crypto.PubkeyToAddress(sk.PublicKey)
	addr := map[common.Address]core.GenesisAccount{
		common.BytesToAddress([]byte{1}): {Balance: big.NewInt(1)}, // ECRecover
		common.BytesToAddress([]byte{2}): {Balance: big.NewInt(1)}, // SHA256
		common.BytesToAddress([]byte{3}): {Balance: big.NewInt(1)}, // RIPEMD
		common.BytesToAddress([]byte{4}): {Balance: big.NewInt(1)}, // Identity
		common.BytesToAddress([]byte{5}): {Balance: big.NewInt(1)}, // ModExp
		common.BytesToAddress([]byte{6}): {Balance: big.NewInt(1)}, // ECAdd
		common.BytesToAddress([]byte{7}): {Balance: big.NewInt(1)}, // ECScalarMul
		common.BytesToAddress([]byte{8}): {Balance: big.NewInt(1)}, // ECPairing
		faucetAddr:                       {Balance: new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(9))},
	}
	alloc := core.GenesisAlloc(addr)
	return backends.NewSimulatedBackend(alloc, 80000000)
}
