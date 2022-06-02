// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package client_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/cryptoriums/packages/client"
	"github.com/cryptoriums/packages/env"
	"github.com/cryptoriums/packages/testing/contracts/bindings/gauge"
	"github.com/cryptoriums/packages/testutil"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-kit/log"
)

func TestEthCall(t *testing.T) {
	ctx := context.Background()

	e, err := env.LoadFromEnvVarOrFile("env", "../env.json", "mainnet")
	testutil.Ok(t, err)

	client, err := client.NewClientCachedNetID(ctx, log.NewNopLogger(), e.Nodes[0].URL)
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
