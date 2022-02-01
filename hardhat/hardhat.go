// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package hardhat

import (
	"context"
	"math/big"
	"os"
	"path"
	"strings"

	contraget "github.com/cryptoriums/contraget/pkg/cli"
	"github.com/cryptoriums/packages/ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
)

func ReplaceContract(ctx context.Context, nodeURL string, contractPath string, contractName string, contractAddrToReplace common.Address) error {
	rpcClient, err := rpc.DialContext(ctx, nodeURL)
	if err != nil {
		return errors.Wrap(err, "creating rpc client")
	}

	cfg := &contraget.Cli{
		Path:       contractPath,
		Name:       "contract",
		ObjectsDst: "tmp",
	}
	defer os.RemoveAll(cfg.DownloadDst)

	if err := contraget.Run(cfg); err != nil {
		return errors.Wrap(err, "generating the contract bin file")
	}

	_bin, err := os.ReadFile(path.Join(cfg.ObjectsDst, contractName+".bin"))
	if err != nil {
		return errors.Wrap(err, "reading the bin file")
	}
	bin := string(_bin)
	indexDeployBin := strings.LastIndex(bin, "60806040523480156")

	err = rpcClient.CallContext(ctx, nil, "hardhat_setCode", contractAddrToReplace, "0x"+bin[indexDeployBin:])
	if err != nil {
		return errors.Wrap(err, "hardhat_setCode call")
	}

	return nil
}

func TxWithImpersonateAccount(ctx context.Context, nodeURL string, from common.Address, to common.Address, abiJ string, funcName string, args []interface{}) (string, error) {
	rpcClient, err := rpc.DialContext(ctx, nodeURL)
	if err != nil {
		return "", errors.Wrap(err, "creating rpc client")
	}

	err = rpcClient.CallContext(ctx, nil, "hardhat_impersonateAccount", from)
	if err != nil {
		return "", errors.Wrap(err, "calling hardhat_impersonateAccount")
	}

	abiParsed, err := abi.JSON(strings.NewReader(abiJ))
	if err != nil {
		return "", errors.Wrap(err, "parsing the abi")
	}
	data, err := abiParsed.Pack(funcName, args...)
	if err != nil {
		return "", errors.Wrap(err, "packing the args")
	}
	optsT := ethereum.SendTransactionOpts{
		From: from,
		To:   to,
		Data: hexutil.Encode(data),
	}
	var txHash string
	err = rpcClient.CallContext(ctx, &txHash, "eth_sendTransaction", optsT)
	if err != nil {
		return "", errors.Wrap(err, "calling eth_sendTransaction")
	}

	return txHash, nil
}

func SetBalance(ctx context.Context, nodeURL string, of common.Address, amnt *big.Int) error {
	rpcClient, err := rpc.DialContext(ctx, nodeURL)
	if err != nil {
		return errors.Wrap(err, "creating rpc client")
	}

	err = rpcClient.CallContext(ctx, nil, "hardhat_setBalance", of, hexutil.EncodeBig(amnt))
	if err != nil {
		return errors.Wrap(err, "calling hardhat_impersonateAccount")
	}

	return nil
}
