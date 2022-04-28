// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package ethereum_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/cryptoriums/packages/env"
	ethereum_p "github.com/cryptoriums/packages/ethereum"
	"github.com/cryptoriums/packages/hardhat"
	"github.com/cryptoriums/packages/testing/contracts/bindings/gauge"
	"github.com/cryptoriums/packages/testutil"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-kit/log"
)

func TestHasPendingTx(t *testing.T) {
	ctx := context.Background()
	e, err := env.LoadFromEnvVarOrFile("env", "../env.json", "http")
	testutil.Ok(t, err)

	acc, err := ethereum_p.AccountFromPrvKey("0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	testutil.Ok(t, err)

	cmd := hardhat.Fork(log.NewNopLogger(), "npx", "hardhat", "node", "--fork", e.Nodes[0].URL, "--fork-block-number", "13858002")
	defer testutil.KillCmd(t, cmd)

	testutil.Ok(t, hardhat.DisableAutoMine(ctx, hardhat.DefaultUrl))
	client, err := ethereum_p.NewClientCachedNetID(ctx, log.NewNopLogger(), hardhat.DefaultUrl)
	testutil.Ok(t, err)

	hasPending, nonceExp, err := ethereum_p.HasPendingTx(ctx, client, acc.PublicKey)
	testutil.Ok(t, err)
	testutil.Assert(t, hasPending == false)

	tx, _, err := ethereum_p.NewSignedTX(
		ctx,
		acc.PrivateKey,
		common.Address{},
		"",
		nonceExp,
		client.NetworkID(),
		"",
		nil,
		100_000,
		100,
		100,
		1,
	)

	testutil.Ok(t, client.SendTransaction(ctx, tx))

	hasPending, nonceAct, err := ethereum_p.HasPendingTx(ctx, client, acc.PublicKey)
	testutil.Ok(t, err)
	testutil.Assert(t, hasPending == true)

	testutil.Equals(t, nonceExp, nonceAct)

	testutil.Ok(t, hardhat.Mine(ctx, hardhat.DefaultUrl))

	hasPending, nonceNew, err := ethereum_p.HasPendingTx(ctx, client, acc.PublicKey)
	testutil.Ok(t, err)
	testutil.Assert(t, hasPending == false)

	testutil.Equals(t, nonceExp+1, nonceNew)

}

func TestEthCall(t *testing.T) {
	ctx := context.Background()

	e, err := env.LoadFromEnvVarOrFile("env", "../env.json", "mainnet")
	testutil.Ok(t, err)

	client, err := ethereum_p.NewClientCachedNetID(ctx, log.NewNopLogger(), e.Nodes[0].URL)
	testutil.Ok(t, err)

	callOpts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: big.NewInt(14178089),
	}

	abi, err := gauge.GaugeMetaData.GetAbi()
	testutil.Ok(t, err)

	stakerAddr := common.HexToAddress("0x989aeb4d175e16225e39e87d0d97a3360524ad80")
	gaugeAddr := common.HexToAddress("0x7ca5b0a2910B33e9759DC7dDB0413949071D7575")

	results := []interface{}{
		new(*big.Int),
	}
	err = bind.NewBoundContract(gaugeAddr, *abi, client, client, client).Call(callOpts, &results, "claimable_tokens", stakerAddr)
	testutil.Ok(t, err)

	r := results[0].(**big.Int)

	testutil.Equals(t, (*r).String(), "448222059400430463396")

}
