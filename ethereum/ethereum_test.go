// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package ethereum

import (
	"crypto/ecdsa"
	"fmt"
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
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-kit/log/level"
	"golang.org/x/net/context"
)

func TestSubscriptionWithRedundancy(t *testing.T) {
	logger := logging.NewLogger()
	logging.ApplyFilter("debug", logger)
	ctx := context.Background()

	abi, err := abi.JSON(strings.NewReader(simple.SimpleStorageABI))
	testutil.Ok(t, err)

	sk, err := crypto.GenerateKey()
	testutil.Ok(t, err)

	transactOpts, err := bind.NewKeyedTransactorWithChainID(sk, big.NewInt(1337))

	var (
		logFilterers []ethereum.LogFilterer
		contactAddrs []common.Address
		contracts    []*simple.SimpleStorage
		backends     []*backends.SimulatedBackend
	)
	logsCount := 4
	for i := 0; i < logsCount; i++ {
		backend := getSimBackend(t, sk)
		testutil.Ok(t, err)
		addr, _, contract, err := simple.DeploySimpleStorage(transactOpts, backend)
		testutil.Ok(t, err)
		_, err = contract.Set(transactOpts, "1111")
		testutil.Ok(t, err)
		backend.Commit()
		logFilterers = append(logFilterers, backend)
		contactAddrs = append(contactAddrs, addr)
		contracts = append(contracts, contract)
		backends = append(backends, backend)
	}

	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   nil,
		Addresses: contactAddrs,
		Topics:    [][]common.Hash{{abi.Events["StorageSet"].ID}},
	}

	sink := make(chan types.Log)

	subs, err := NewSubscriptionWithRedundancy(ctx, logger, logFilterers, query, sink)
	defer subs.Unsubscribe()

	level.Info(logger).Log("msg", "new log", "data", fmt.Sprintf("%+v", (<-sink)))
	time.Sleep(time.Second)

	rcvd, sent := subs.(*SubscriptionWithRedundancy).EventsCount()
	testutil.Equals(t, int64(logsCount), rcvd)
	testutil.Equals(t, int64(1), sent)

	_, err = contracts[0].Set(transactOpts, "2222")
	testutil.Ok(t, err)
	backends[0].Commit()

	level.Info(logger).Log("msg", "new log", "data", fmt.Sprintf("%+v", (<-sink)))
	time.Sleep(time.Second)

	rcvd, sent = subs.(*SubscriptionWithRedundancy).EventsCount()
	testutil.Equals(t, int64(logsCount+1), rcvd)
	testutil.Equals(t, int64(2), sent)
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
