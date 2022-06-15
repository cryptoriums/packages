// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package localnode

import (
	"context"
	"math/big"
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

func Test_Hardhat(t *testing.T) {
	ctx := context.Background()

	e, err := env.LoadFromEnvVarOrFile("env", "../env.json", "mainnet")
	testutil.Ok(t, err)

	cmd := Fork(log.NewNopLogger(), "npx", "hardhat", "node", "--fork", e.Nodes[0].URL, "--fork-block-number", "13858002")
	defer testutil.KillCmd(t, cmd)

	client, err := ethclient.DialContext(ctx, DefaultUrl)
	testutil.Ok(t, err)

	t.Run("ReplaceContract", func(t *testing.T) {
		err = ReplaceContract(
			ctx,
			Hardhat,
			DefaultUrl,
			"../testing/contracts/source/Booster.sol",
			"Booster",
			common.HexToAddress(boosterContract),
		)
		testutil.Ok(t, err)

		boosterInstance, err := booster.NewBooster(common.HexToAddress(boosterContract), client)
		testutil.Ok(t, err)

		repl, err := boosterInstance.PoolLength(&bind.CallOpts{Context: ctx})
		testutil.Ok(t, err)

		testutil.Equals(t, repl, big.NewInt(67))
	})

	t.Run("ImpersonateAccountReplaceContract", func(t *testing.T) {
		from := common.HexToAddress("0xa3c5a1e09150b75ff251c1a7815a07182c3de2fb")
		to := common.HexToAddress(boosterContract)
		newFeeManager := common.HexToAddress("0x0")

		// Set some balance of the account which will run the impersonated TX.
		{
			newBalance := big_p.FloatToBigIntMul(1000, params.Ether)
			err = SetBalance(ctx, Hardhat, DefaultUrl, from, newBalance)
			testutil.Ok(t, err)

			newAct, err := client.BalanceAt(ctx, from, nil)
			testutil.Ok(t, err)
			testutil.Equals(t, newAct, newBalance)
		}

		_, err = TxWithImpersonateAccount(
			ctx,
			Hardhat,
			DefaultUrl,
			from,
			to,
			booster.BoosterABI,
			"setFeeManager",
			newFeeManager,
		)
		testutil.Ok(t, err)

		boosterInstance, err := booster.NewBooster(common.HexToAddress(boosterContract), client)
		testutil.Ok(t, err)

		addrAct, err := boosterInstance.FeeManager(&bind.CallOpts{Context: ctx})
		testutil.Ok(t, err)

		testutil.Equals(t, addrAct, newFeeManager)
	})

}

// func Test_Anvil(t *testing.T) {

// 	ctx := context.Background()

// 	e, err := env.LoadFromEnvVarOrFile("env", "../env.json", "mainnet")
// 	testutil.Ok(t, err)

// 	cmd := Fork(log.NewNopLogger(), "anvil", "--fork-url", e.Nodes[0].URL, "--fork-block-number", "13858002")
// 	defer testutil.KillCmd(t, cmd)

// 	client, err := ethclient.DialContext(ctx, DefaultUrl)
// 	testutil.Ok(t, err)

// 	t.Run("ReplaceContract", func(t *testing.T) {
// 		err = ReplaceContract(
// 			ctx,
// 			Anvil,
// 			DefaultUrl,
// 			"../testing/contracts/source/Booster.sol",
// 			"Booster",
// 			common.HexToAddress(boosterContract),
// 		)
// 		testutil.Ok(t, err)

// 		boosterInstance, err := booster.NewBooster(common.HexToAddress(boosterContract), client)
// 		testutil.Ok(t, err)

// 		repl, err := boosterInstance.PoolLength(&bind.CallOpts{Context: ctx})
// 		testutil.Ok(t, err)

// 		testutil.Equals(t, repl, big.NewInt(71))
// 	})

// 	t.Run("ImpersonateAccountReplaceContract", func(t *testing.T) {
// 		from := common.HexToAddress("0xa3c5a1e09150b75ff251c1a7815a07182c3de2fb")
// 		to := common.HexToAddress(boosterContract)
// 		newFeeManager := common.HexToAddress("0x0")

// 		// Set some balance of the account which will run the impersonated TX.
// 		{
// 			newBalance := big_p.FloatToBigIntMul(1000, params.Ether)
// 			err = SetBalance(ctx, Anvil, DefaultUrl, from, newBalance)
// 			testutil.Ok(t, err)

// 			newAct, err := client.BalanceAt(ctx, from, nil)
// 			testutil.Ok(t, err)
// 			testutil.Equals(t, newAct, newBalance)
// 		}

// 		_, err = TxWithImpersonateAccount(
// 			ctx,
// 			Anvil,
// 			DefaultUrl,
// 			from,
// 			to,
// 			booster.BoosterABI,
// 			"setFeeManager",
// 			newFeeManager,
// 		)
// 		testutil.Ok(t, err)

// 		boosterInstance, err := booster.NewBooster(common.HexToAddress(boosterContract), client)
// 		testutil.Ok(t, err)

// 		addrAct, err := boosterInstance.FeeManager(&bind.CallOpts{Context: ctx})
// 		testutil.Ok(t, err)

// 		testutil.Equals(t, addrAct, newFeeManager)
// 	})

// }
