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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/log"
	"github.com/stretchr/testify/require"
)

const boosterContract string = "0xf403c135812408bfbe8713b5a23a04b3d48aae31"
const boosterFeeManager string = "0xa3c5a1e09150b75ff251c1a7815a07182c3de2fb"

var blockNumber = uint64(13858047)

func Test_Hardhat(t *testing.T) {
	ctx := context.Background()
	e, err := env.LoadFromEnvVarOrFile("env", "../env.json", "http")
	require.NoError(t, err)

	ln, err := New(ctx, log.NewNopLogger(), Hardhat, e.Nodes[0].URL, strconv.Itoa(int(blockNumber)))
	require.NoError(t, err)

	defer func() {
		if err := ln.Stop(); err != nil {
			require.NoError(t, err)
		}
	}()

	client, err := ethclient.DialContext(ctx, ln.GetNodeURL())
	require.NoError(t, err)

	t.Run("MineAndReset", func(t *testing.T) {
		initialBlock, err := client.BlockNumber(ctx)
		require.NoError(t, err)

		require.Equal(t, initialBlock, uint64(blockNumber))

		err = ln.Mine(ctx)
		require.NoError(t, err)
		err = ln.Mine(ctx)
		require.NoError(t, err)

		blockAfterMine, err := client.BlockNumber(ctx)
		require.NoError(t, err)
		require.Equal(t, initialBlock+2, blockAfterMine)

		blockToResetExp := blockAfterMine - 1
		err = ln.Reset(ctx, blockToResetExp)
		require.NoError(t, err)
		blockAfterResetAct, err := client.BlockNumber(ctx)
		require.NoError(t, err)
		require.Equal(t, blockToResetExp, blockAfterResetAct)
	})

	t.Run("ReplaceContract", func(t *testing.T) {
		err = ln.ReplaceContract(
			ctx,
			"../testing/contracts/source/Booster.sol",
			"Booster",
			common.HexToAddress(boosterContract),
		)
		require.NoError(t, err)

		boosterInstance, err := booster.NewBooster(common.HexToAddress(boosterContract), client)
		require.NoError(t, err)

		repl, err := boosterInstance.PoolLength(&bind.CallOpts{Context: ctx})
		require.NoError(t, err)

		require.Equal(t, big.NewInt(67), repl)
	})

	t.Run("SetBalance", func(t *testing.T) {
		from := ln.GetAccounts()[0].PublicKey
		newBalance := big_p.FromFloatMul(1000, params.Ether)
		err = ln.SetBalance(ctx, from, newBalance)
		require.NoError(t, err)

		newAct, err := client.BalanceAt(ctx, from, nil)
		require.NoError(t, err)
		require.Equal(t, newBalance, newAct)
	})

	t.Run("ImpersonateAccountReplaceContract", func(t *testing.T) {
		initialManager := common.HexToAddress(boosterFeeManager)
		to := common.HexToAddress(boosterContract)
		newFeeManager := ln.GetAccounts()[0].PublicKey

		boosterInstance, err := booster.NewBooster(common.HexToAddress(boosterContract), client)
		require.NoError(t, err)

		addrFeeMngr, err := boosterInstance.FeeManager(&bind.CallOpts{Context: ctx})
		require.NoError(t, err)

		require.Equal(t, initialManager, addrFeeMngr)

		// Set some balance of the account which will run the impersonated TX.
		{
			newBalance := big_p.FromFloatMul(1000, params.Ether)
			err = ln.SetBalance(ctx, initialManager, newBalance)
			require.NoError(t, err)

			newAct, err := client.BalanceAt(ctx, initialManager, nil)
			require.NoError(t, err)
			require.Equal(t, newBalance, newAct)
		}

		_, err = ln.TxWithImpersonateAccount(
			ctx,
			initialManager,
			to,
			booster.BoosterABI,
			"setFeeManager",
			newFeeManager,
		)
		require.NoError(t, err)

		addrAct, err := boosterInstance.FeeManager(&bind.CallOpts{Context: ctx})
		require.NoError(t, err)
		require.Equal(t, newFeeManager, addrAct)

		// Revert the fee manager with ImpersonateAccountWithData without ABI parsing.
		_, err = ln.TxWithImpersonateAccountWithData(
			ctx,
			newFeeManager,
			to,
			common.Hex2Bytes("472d35b9000000000000000000000000a3c5a1e09150b75ff251c1a7815a07182c3de2fb"),
		)
		require.NoError(t, err)

		addrAct, err = boosterInstance.FeeManager(&bind.CallOpts{Context: ctx})
		require.NoError(t, err)

		require.Equal(t, initialManager, addrAct)
	})
}

func Test_Foundry_Anvil(t *testing.T) {
	ctx := context.Background()

	e, err := env.LoadFromEnvVarOrFile("env", "../env.json", "http")
	require.NoError(t, err)

	ln, err := New(ctx, log.NewNopLogger(), Anvil, e.Nodes[0].URL, strconv.Itoa(int(blockNumber)))
	require.NoError(t, err)

	defer func() {
		if err := ln.Stop(); err != nil {
			require.NoError(t, err)
		}
	}()

	client, err := ethclient.DialContext(ctx, ln.GetNodeURL())
	require.NoError(t, err)

	t.Run("MineAndReset", func(t *testing.T) {
		initialBlock, err := client.BlockNumber(ctx)
		require.NoError(t, err)

		require.Equal(t, initialBlock, uint64(blockNumber))

		err = ln.Mine(ctx)
		require.NoError(t, err)
		err = ln.Mine(ctx)
		require.NoError(t, err)

		blockAfterMine, err := client.BlockNumber(ctx)
		require.NoError(t, err)
		require.Equal(t, initialBlock+2, blockAfterMine)

		blockToResetExp := blockAfterMine - 1
		err = ln.Reset(ctx, blockToResetExp)
		require.NoError(t, err)
		blockAfterResetAct, err := client.BlockNumber(ctx)
		require.NoError(t, err)
		require.Equal(t, blockToResetExp, blockAfterResetAct)
	})

	t.Run("ReplaceContract", func(t *testing.T) {

		err = ln.ReplaceContract(
			ctx,
			"../testing/contracts/source/Booster.sol",
			"Booster",
			common.HexToAddress(boosterContract),
		)
		require.NoError(t, err)

		boosterInstance, err := booster.NewBooster(common.HexToAddress(boosterContract), client)
		require.NoError(t, err)

		repl, err := boosterInstance.PoolLength(&bind.CallOpts{Context: ctx})
		require.NoError(t, err)

		require.Equal(t, big.NewInt(67), repl)
	})

	t.Run("SetBalance", func(t *testing.T) {
		from := ln.GetAccounts()[0].PublicKey
		newBalance := big_p.FromFloatMul(1000, params.Ether)
		err = ln.SetBalance(ctx, from, newBalance)
		require.NoError(t, err)

		newAct, err := client.BalanceAt(ctx, from, nil)
		require.NoError(t, err)
		require.Equal(t, newBalance, newAct)
	})

	t.Run("ImpersonateAccountReplaceContract", func(t *testing.T) {
		initialManager := common.HexToAddress(boosterFeeManager)
		to := common.HexToAddress(boosterContract)
		newFeeManager := ln.GetAccounts()[0].PublicKey

		boosterInstance, err := booster.NewBooster(common.HexToAddress(boosterContract), client)
		require.NoError(t, err)

		addrFeeMngr, err := boosterInstance.FeeManager(&bind.CallOpts{Context: ctx})
		require.NoError(t, err)

		require.Equal(t, initialManager, addrFeeMngr)

		// Set some balance of the account which will run the impersonated TX.
		{
			newBalance := big_p.FromFloatMul(1000, params.Ether)
			err = ln.SetBalance(ctx, initialManager, newBalance)
			require.NoError(t, err)

			newAct, err := client.BalanceAt(ctx, initialManager, nil)
			require.NoError(t, err)
			require.Equal(t, newBalance, newAct)
		}

		_, err = ln.TxWithImpersonateAccount(
			ctx,
			initialManager,
			to,
			booster.BoosterABI,
			"setFeeManager",
			newFeeManager,
		)
		require.NoError(t, err)

		addrAct, err := boosterInstance.FeeManager(&bind.CallOpts{Context: ctx})
		require.NoError(t, err)

		require.Equal(t, newFeeManager, addrAct)

		// Revert the fee manager with ImpersonateAccountWithData without ABI parsing.
		_, err = ln.TxWithImpersonateAccountWithData(
			ctx,
			newFeeManager,
			to,
			common.Hex2Bytes("472d35b9000000000000000000000000a3c5a1e09150b75ff251c1a7815a07182c3de2fb"),
		)
		require.NoError(t, err)

		addrAct, err = boosterInstance.FeeManager(&bind.CallOpts{Context: ctx})
		require.NoError(t, err)

		require.Equal(t, initialManager, addrAct)
	})

}
