// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package hardhat

import (
	"context"
	"math/big"
	"os"
	"testing"

	"github.com/cryptoriums/packages/client"
	math_p "github.com/cryptoriums/packages/math"
	"github.com/cryptoriums/packages/pkg/contracts/booster"
	"github.com/cryptoriums/packages/private_file"
	"github.com/cryptoriums/packages/testutil"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

func TestReplaceContract(t *testing.T) {
	ctx := context.Background()

	envFileData, err := os.ReadFile("../.env")
	testutil.Ok(t, err)
	envVars, err := private_file.SetEnvVars(envFileData)
	testutil.Ok(t, err)
	nodes, err := client.ParseNodes(envVars)
	testutil.Ok(t, err)

	cmd := testutil.HardhatFork(t, "npx", "hardhat", "node", "--fork", nodes[0], "--fork-block-number", "13858002")
	defer testutil.KillCmd(t, cmd)

	err = ReplaceContract(
		ctx,
		"http://127.0.0.1:8545",
		"../contracts/Booster.sol",
		"Booster",
		common.HexToAddress("0xf403c135812408bfbe8713b5a23a04b3d48aae31"),
	)
	testutil.Ok(t, err)

	client, err := ethclient.DialContext(ctx, "http://127.0.0.1:8545/")
	testutil.Ok(t, err)

	boosterInstance, err := booster.NewBooster(common.HexToAddress("0xf403c135812408bfbe8713b5a23a04b3d48aae31"), client)
	testutil.Ok(t, err)

	repl, err := boosterInstance.PoolLength(&bind.CallOpts{Context: ctx})
	testutil.Ok(t, err)

	testutil.Equals(t, repl, big.NewInt(67))
}

func TestImpersonateAccount(t *testing.T) {
	ctx := context.Background()

	envFileData, err := os.ReadFile("../.env")
	testutil.Ok(t, err)
	envVars, err := private_file.SetEnvVars(envFileData)
	testutil.Ok(t, err)
	nodes, err := client.ParseNodes(envVars)
	testutil.Ok(t, err)

	cmd := testutil.HardhatFork(t, "npx", "hardhat", "node", "--fork", nodes[0], "--fork-block-number", "13858002")
	defer testutil.KillCmd(t, cmd)

	client, err := ethclient.DialContext(ctx, "http://127.0.0.1:8545/")
	testutil.Ok(t, err)

	nodeURL := "http://127.0.0.1:8545"
	from := common.HexToAddress("0xa3c5a1e09150b75ff251c1a7815a07182c3de2fb")
	to := common.HexToAddress("0xf403c135812408bfbe8713b5a23a04b3d48aae31")
	newAddr := common.HexToAddress("0x0")

	// Set some balance of the account which will run the impersonated TX.
	{
		newBalance := math_p.FloatToBigIntMul(1000, params.Ether)
		err = SetBalance(ctx, nodeURL, from, newBalance)
		testutil.Ok(t, err)

		newAct, err := client.BalanceAt(ctx, from, nil)
		testutil.Ok(t, err)
		testutil.Equals(t, newAct, newBalance)
	}

	_, err = TxWithImpersonateAccount(
		ctx,
		nodeURL,
		from,
		to,
		booster.BoosterABI,
		"setFeeManager",
		[]interface{}{
			newAddr,
		},
	)
	testutil.Ok(t, err)

	boosterInstance, err := booster.NewBooster(common.HexToAddress("0xf403c135812408bfbe8713b5a23a04b3d48aae31"), client)
	testutil.Ok(t, err)

	addrAct, err := boosterInstance.FeeManager(&bind.CallOpts{Context: ctx})
	testutil.Ok(t, err)

	testutil.Equals(t, addrAct, newAddr)
}
