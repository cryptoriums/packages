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
	"github.com/cryptoriums/packages/testing/contracts/bindings/booster"
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

	boosterInstance, err := booster.NewBooster(common.HexToAddress("0xf403c135812408bfbe8713b5a23a04b3d48aae31"), client)
	testutil.Ok(t, err)

	callOpts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: big.NewInt(14178089),
	}

	abi, err := gauge.GaugeMetaData.GetAbi()
	testutil.Ok(t, err)

	stakerAddr := common.HexToAddress("0x989aeb4d175e16225e39e87d0d97a3360524ad80")

	pool, err := boosterInstance.PoolInfo(nil, big.NewInt(0))
	testutil.Ok(t, err)

	results := []interface{}{
		new(*big.Int),
	}
	err = bind.NewBoundContract(pool.Gauge, *abi, client, client, client).Call(callOpts, &results, "claimable_tokens", stakerAddr)
	testutil.Ok(t, err)

	r := results[0].(**big.Int)

	testutil.Equals(t, (*r).String(), "448222059400430463396")

}
