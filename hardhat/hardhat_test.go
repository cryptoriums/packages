// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package hardhat

import (
	"context"
	"math/big"
	"testing"

	"github.com/cryptoriums/packages/env"
	math_p "github.com/cryptoriums/packages/math"
	"github.com/cryptoriums/packages/testing/contracts/bindings/booster"
	"github.com/cryptoriums/packages/testutil"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/log"
)

func TestReplaceContract(t *testing.T) {
	ctx := context.Background()

	e, err := env.LoadFromEnvVarOrFile("env", "../env.json", "http")
	testutil.Ok(t, err)

	cmd := Fork(log.NewNopLogger(), "npx", "hardhat", "node", "--fork", e.Nodes[0].URL, "--fork-block-number", "13858002")
	defer testutil.KillCmd(t, cmd)

	err = ReplaceContract(
		ctx,
		DefaultUrl,
		"../testing/contracts/source/Booster.sol",
		"Booster",
		common.HexToAddress("0xf403c135812408bfbe8713b5a23a04b3d48aae31"),
	)
	testutil.Ok(t, err)

	client, err := ethclient.DialContext(ctx, DefaultUrl)
	testutil.Ok(t, err)

	boosterInstance, err := booster.NewBooster(common.HexToAddress("0xf403c135812408bfbe8713b5a23a04b3d48aae31"), client)
	testutil.Ok(t, err)

	repl, err := boosterInstance.PoolLength(&bind.CallOpts{Context: ctx})
	testutil.Ok(t, err)

	testutil.Equals(t, repl, big.NewInt(67))
}

func TestImpersonateAccount(t *testing.T) {
	ctx := context.Background()

	e, err := env.LoadFromEnvVarOrFile("env", "../env.json", "http")
	testutil.Ok(t, err)

	cmd := Fork(log.NewNopLogger(), "npx", "hardhat", "node", "--fork", e.Nodes[0].URL, "--fork-block-number", "13858002")
	defer testutil.KillCmd(t, cmd)

	client, err := ethclient.DialContext(ctx, DefaultUrl)
	testutil.Ok(t, err)

	from := common.HexToAddress("0xa3c5a1e09150b75ff251c1a7815a07182c3de2fb")
	to := common.HexToAddress("0xf403c135812408bfbe8713b5a23a04b3d48aae31")
	newAddr := common.HexToAddress("0x0")

	// Set some balance of the account which will run the impersonated TX.
	{
		newBalance := math_p.FloatToBigIntMul(1000, params.Ether)
		err = SetBalance(ctx, DefaultUrl, from, newBalance)
		testutil.Ok(t, err)

		newAct, err := client.BalanceAt(ctx, from, nil)
		testutil.Ok(t, err)
		testutil.Equals(t, newAct, newBalance)
	}

	_, err = TxWithImpersonateAccount(
		ctx,
		DefaultUrl,
		from,
		to,
		booster.BoosterABI,
		"setFeeManager",
		newAddr,
	)
	testutil.Ok(t, err)

	boosterInstance, err := booster.NewBooster(common.HexToAddress("0xf403c135812408bfbe8713b5a23a04b3d48aae31"), client)
	testutil.Ok(t, err)

	addrAct, err := boosterInstance.FeeManager(&bind.CallOpts{Context: ctx})
	testutil.Ok(t, err)

	testutil.Equals(t, addrAct, newAddr)
}
