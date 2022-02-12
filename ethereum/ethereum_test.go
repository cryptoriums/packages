// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package ethereum_test

import (
	"context"
	"math/big"
	"os"
	"testing"

	"github.com/cryptoriums/packages/client"
	ethereum_p "github.com/cryptoriums/packages/ethereum"
	"github.com/cryptoriums/packages/private_file"
	"github.com/cryptoriums/packages/testing/contracts/bindings/gauge"
	"github.com/cryptoriums/packages/testutil"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-kit/log"
)

func TestEthCall(t *testing.T) {
	ctx := context.Background()

	envFileData, err := os.ReadFile("../.env")
	testutil.OkIgnoreNotFount(t, err)
	envVars, err := private_file.SetEnvVars(envFileData)
	testutil.Ok(t, err)
	nodes, err := client.ParseNodes(envVars)
	testutil.Ok(t, err)

	client, err := ethereum_p.NewClientCachedNetID(ctx, log.NewNopLogger(), nodes[0])
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
