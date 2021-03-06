// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package localnode

import (
	"context"
	"math/big"
	"strconv"
	"testing"

	big_p "github.com/cryptoriums/packages/big"
	"github.com/cryptoriums/packages/env"
	"github.com/cryptoriums/packages/testing/contracts/bindings/booster"
	"github.com/cryptoriums/packages/testutil"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/log"
)

const boosterContract string = "0xf403c135812408bfbe8713b5a23a04b3d48aae31"
const boosterFeeManager string = "0xa3c5a1e09150b75ff251c1a7815a07182c3de2fb"

var blockNumber = uint64(13858047)

func Test_Hardhat(t *testing.T) {
	ctx := context.Background()
	e, err := env.LoadFromEnvVarOrFile("env", "../env.json", "http")
	testutil.Ok(t, err)

	ln, err := New(ctx, log.NewNopLogger(), Hardhat, e.Nodes[0].URL, strconv.Itoa(int(blockNumber)))
	testutil.Ok(t, err)

	defer func() {
		if err := ln.Stop(); err != nil {
			testutil.Ok(t, err)
		}
	}()

	client, err := ethclient.DialContext(ctx, ln.GetNodeURL())
	testutil.Ok(t, err)

	t.Run("MineAndReset", func(t *testing.T) {
		initialBlock, err := client.BlockNumber(ctx)
		testutil.Ok(t, err)

		testutil.Equals(t, initialBlock, uint64(blockNumber))

		err = ln.Mine(ctx)
		testutil.Ok(t, err)
		err = ln.Mine(ctx)
		testutil.Ok(t, err)

		blockAfterMine, err := client.BlockNumber(ctx)
		testutil.Ok(t, err)
		testutil.Equals(t, initialBlock+2, blockAfterMine)

		blockToResetExp := blockAfterMine - 1
		err = ln.Reset(ctx, blockToResetExp)
		testutil.Ok(t, err)
		blockAfterResetAct, err := client.BlockNumber(ctx)
		testutil.Ok(t, err)
		testutil.Equals(t, blockToResetExp, blockAfterResetAct)
	})

	t.Run("ReplaceContract", func(t *testing.T) {
		err = ln.ReplaceContract(
			ctx,
			"../testing/contracts/source/Booster.sol",
			"Booster",
			common.HexToAddress(boosterContract),
		)
		testutil.Ok(t, err)

		boosterInstance, err := booster.NewBooster(common.HexToAddress(boosterContract), client)
		testutil.Ok(t, err)

		repl, err := boosterInstance.PoolLength(&bind.CallOpts{Context: ctx})
		testutil.Ok(t, err)

		testutil.Equals(t, big.NewInt(67), repl)
	})

	t.Run("SetBalance", func(t *testing.T) {
		from := ln.GetAccounts()[0].PublicKey
		newBalance := big_p.FloatToBigIntMul(1000, params.Ether)
		err = ln.SetBalance(ctx, from, newBalance)
		testutil.Ok(t, err)

		newAct, err := client.BalanceAt(ctx, from, nil)
		testutil.Ok(t, err)
		testutil.Equals(t, newBalance, newAct)
	})

	t.Run("ImpersonateAccountReplaceContract", func(t *testing.T) {
		initialManager := common.HexToAddress(boosterFeeManager)
		to := common.HexToAddress(boosterContract)
		newFeeManager := ln.GetAccounts()[0].PublicKey

		boosterInstance, err := booster.NewBooster(common.HexToAddress(boosterContract), client)
		testutil.Ok(t, err)

		addrFeeMngr, err := boosterInstance.FeeManager(&bind.CallOpts{Context: ctx})
		testutil.Ok(t, err)

		testutil.Equals(t, initialManager, addrFeeMngr)

		// Set some balance of the account which will run the impersonated TX.
		{
			newBalance := big_p.FloatToBigIntMul(1000, params.Ether)
			err = ln.SetBalance(ctx, initialManager, newBalance)
			testutil.Ok(t, err)

			newAct, err := client.BalanceAt(ctx, initialManager, nil)
			testutil.Ok(t, err)
			testutil.Equals(t, newBalance, newAct)
		}

		_, err = ln.TxWithImpersonateAccount(
			ctx,
			initialManager,
			to,
			booster.BoosterABI,
			"setFeeManager",
			newFeeManager,
		)
		testutil.Ok(t, err)

		addrAct, err := boosterInstance.FeeManager(&bind.CallOpts{Context: ctx})
		testutil.Ok(t, err)
		testutil.Equals(t, newFeeManager, addrAct)

		// Revert the fee manager with ImpersonateAccountWithData without ABI parsing.
		_, err = ln.TxWithImpersonateAccountWithData(
			ctx,
			newFeeManager,
			to,
			common.Hex2Bytes("472d35b9000000000000000000000000a3c5a1e09150b75ff251c1a7815a07182c3de2fb"),
		)
		testutil.Ok(t, err)

		addrAct, err = boosterInstance.FeeManager(&bind.CallOpts{Context: ctx})
		testutil.Ok(t, err)

		testutil.Equals(t, initialManager, addrAct)
	})
}

func Test_Foundry_Anvil(t *testing.T) {
	ctx := context.Background()

	e, err := env.LoadFromEnvVarOrFile("env", "../env.json", "http")
	testutil.Ok(t, err)

	ln, err := New(ctx, log.NewNopLogger(), Anvil, e.Nodes[0].URL, strconv.Itoa(int(blockNumber)))
	testutil.Ok(t, err)

	defer func() {
		if err := ln.Stop(); err != nil {
			testutil.Ok(t, err)
		}
	}()

	client, err := ethclient.DialContext(ctx, ln.GetNodeURL())
	testutil.Ok(t, err)

	t.Run("MineAndReset", func(t *testing.T) {
		initialBlock, err := client.BlockNumber(ctx)
		testutil.Ok(t, err)

		testutil.Equals(t, initialBlock, uint64(blockNumber))

		err = ln.Mine(ctx)
		testutil.Ok(t, err)
		err = ln.Mine(ctx)
		testutil.Ok(t, err)

		blockAfterMine, err := client.BlockNumber(ctx)
		testutil.Ok(t, err)
		testutil.Equals(t, initialBlock+2, blockAfterMine)

		blockToResetExp := blockAfterMine - 1
		err = ln.Reset(ctx, blockToResetExp)
		testutil.Ok(t, err)
		blockAfterResetAct, err := client.BlockNumber(ctx)
		testutil.Ok(t, err)
		testutil.Equals(t, blockToResetExp, blockAfterResetAct)
	})

	t.Run("ReplaceContract", func(t *testing.T) {

		err = ln.ReplaceContract(
			ctx,
			"../testing/contracts/source/Booster.sol",
			"Booster",
			common.HexToAddress(boosterContract),
		)
		testutil.Ok(t, err)

		boosterInstance, err := booster.NewBooster(common.HexToAddress(boosterContract), client)
		testutil.Ok(t, err)

		repl, err := boosterInstance.PoolLength(&bind.CallOpts{Context: ctx})
		testutil.Ok(t, err)

		testutil.Equals(t, big.NewInt(67), repl)
	})

	t.Run("SetBalance", func(t *testing.T) {
		from := ln.GetAccounts()[0].PublicKey
		newBalance := big_p.FloatToBigIntMul(1000, params.Ether)
		err = ln.SetBalance(ctx, from, newBalance)
		testutil.Ok(t, err)

		newAct, err := client.BalanceAt(ctx, from, nil)
		testutil.Ok(t, err)
		testutil.Equals(t, newBalance, newAct)
	})

	t.Run("ImpersonateAccountReplaceContract", func(t *testing.T) {
		initialManager := common.HexToAddress(boosterFeeManager)
		to := common.HexToAddress(boosterContract)
		newFeeManager := ln.GetAccounts()[0].PublicKey

		boosterInstance, err := booster.NewBooster(common.HexToAddress(boosterContract), client)
		testutil.Ok(t, err)

		addrFeeMngr, err := boosterInstance.FeeManager(&bind.CallOpts{Context: ctx})
		testutil.Ok(t, err)

		testutil.Equals(t, initialManager, addrFeeMngr)

		// Set some balance of the account which will run the impersonated TX.
		{
			newBalance := big_p.FloatToBigIntMul(1000, params.Ether)
			err = ln.SetBalance(ctx, initialManager, newBalance)
			testutil.Ok(t, err)

			newAct, err := client.BalanceAt(ctx, initialManager, nil)
			testutil.Ok(t, err)
			testutil.Equals(t, newBalance, newAct)
		}

		_, err = ln.TxWithImpersonateAccount(
			ctx,
			initialManager,
			to,
			booster.BoosterABI,
			"setFeeManager",
			newFeeManager,
		)
		testutil.Ok(t, err)

		addrAct, err := boosterInstance.FeeManager(&bind.CallOpts{Context: ctx})
		testutil.Ok(t, err)

		testutil.Equals(t, newFeeManager, addrAct)

		// Revert the fee manager with ImpersonateAccountWithData without ABI parsing.
		_, err = ln.TxWithImpersonateAccountWithData(
			ctx,
			newFeeManager,
			to,
			common.Hex2Bytes("472d35b9000000000000000000000000a3c5a1e09150b75ff251c1a7815a07182c3de2fb"),
		)
		testutil.Ok(t, err)

		addrAct, err = boosterInstance.FeeManager(&bind.CallOpts{Context: ctx})
		testutil.Ok(t, err)

		testutil.Equals(t, initialManager, addrAct)
	})

}
