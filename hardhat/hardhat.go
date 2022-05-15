// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package hardhat

import (
	"bufio"
	"context"
	"math/big"
	"os"
	"os/exec"
	"path"
	"strings"
	"syscall"
	"time"

	contraget "github.com/cryptoriums/contraget/pkg/cli"
	ethereum_p "github.com/cryptoriums/packages/ethereum"
	"github.com/cryptoriums/packages/logging"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/pkg/errors"
)

const DefaultUrl = "ws://127.0.0.1:8545"

var Accounts []ethereum_p.Account

func init() {
	for _, addr := range []string{
		"0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
		"0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d",
		"0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a",
		"0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6",
		"0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a",
		"0x92db14e403b83dfe3df233f83dfa3a0d7096f21ca9b0d6d6b8d88b2b4ec1564e",
		"0x4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356",
		"0xdbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97",
		"0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6",
		"0xf214f2b2cd398c806f84e317254e0f0b801d0643303237d97a22a48e01628897"} {

		acc, err := ethereum_p.AccountFromPrvKey(addr)
		if err != nil {
			panic(err)
		}
		Accounts = append(Accounts, acc)
	}
}

func Fork(logger log.Logger, args ...string) *exec.Cmd {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	cmdReaderStdOut, err := cmd.StdoutPipe()
	logging.ExitOnError(logger, err)
	cmdReaderStdErr, err := cmd.StderrPipe()
	logging.ExitOnError(logger, err)

	go func() {
		scanner := bufio.NewScanner(cmdReaderStdOut)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			level.Info(logger).Log("msg", scanner.Text())
		}
	}()

	go func() {
		scanner := bufio.NewScanner(cmdReaderStdErr)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			panic(scanner.Text())
		}
	}()
	logging.ExitOnError(logger, cmd.Start())

	for {
		ctx, cncl := context.WithTimeout(context.Background(), 2*time.Second)
		defer cncl()
		client, err := ethclient.DialContext(ctx, "http://localhost:8545")
		if err == nil {
			_, err := client.BlockNumber(ctx)
			if err == nil {
				break
			}
		}
		level.Error(logger).Log("msg", "error connecting will retry")
		time.Sleep(time.Second)
	}
	return cmd
}

func ReplaceContract(ctx context.Context, nodeURL string, contractPath string, contractName string, contractAddrToReplace common.Address) error {
	rpcClient, err := rpc.DialContext(ctx, nodeURL)
	if err != nil {
		return errors.Wrap(err, "creating rpc client")
	}
	defer rpcClient.Close()

	cfg := &contraget.Cli{
		Path:        contractPath,
		ObjectsDst:  "tmp",
		DownloadDst: "tmp",
	}

	if err := contraget.Run(cfg); err != nil {
		return errors.Wrap(err, "generating the contract bin file")
	}

	_bin, err := os.ReadFile(path.Join(cfg.ObjectsDst, contractName+".bin"))
	if err != nil {
		return errors.Wrap(err, "reading the bin file")
	}
	bin := string(_bin)
	// For solidity contracts
	startDeployBin := strings.LastIndex(bin, "60806040523480156")
	// For vyper contracts.
	if startDeployBin == -1 {
		startDeployBin = strings.LastIndex(bin, "600436")
	}

	if startDeployBin == -1 {
		return errors.New("start index of runtime Bytecode not found in the generated binary file")
	}

	err = rpcClient.CallContext(ctx, nil, "hardhat_setCode", contractAddrToReplace, "0x"+bin[startDeployBin:])
	if err != nil {
		return errors.Wrap(err, "hardhat_setCode call")
	}

	return nil
}

func DisableAutoMine(ctx context.Context, nodeURL string) error {
	rpcClient, err := rpc.DialContext(ctx, nodeURL)
	if err != nil {
		return errors.Wrap(err, "creating rpc client")
	}
	defer rpcClient.Close()

	err = rpcClient.CallContext(ctx, nil, "evm_setAutomine", false)
	if err != nil {
		return errors.Wrap(err, "calling evm_setAutomine")
	}

	return nil
}

func Mine(ctx context.Context, nodeURL string) error {
	rpcClient, err := rpc.DialContext(ctx, nodeURL)
	if err != nil {
		return errors.Wrap(err, "creating rpc client")
	}
	defer rpcClient.Close()

	err = rpcClient.CallContext(ctx, nil, "evm_mine")
	if err != nil {
		return errors.Wrap(err, "calling evm_mine")
	}

	return nil
}

func SetNextBlockTimestamp(ctx context.Context, nodeURL string, ts int64) error {
	rpcClient, err := rpc.DialContext(ctx, nodeURL)
	if err != nil {
		return errors.Wrap(err, "creating rpc client")
	}
	defer rpcClient.Close()

	err = rpcClient.CallContext(ctx, nil, "evm_setNextBlockTimestamp", big.NewInt(ts))
	if err != nil {
		return errors.Wrap(err, "calling evm_setNextBlockTimestamp")
	}

	return nil
}

func TxWithImpersonateAccount(ctx context.Context, nodeURL string, from common.Address, to common.Address, abiJ string, funcName string, args []interface{}) (string, error) {
	rpcClient, err := rpc.DialContext(ctx, nodeURL)
	if err != nil {
		return "", errors.Wrap(err, "creating rpc client")
	}
	defer rpcClient.Close()

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
	optsT := ethereum_p.SendTransactionOpts{
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
	defer rpcClient.Close()

	err = rpcClient.CallContext(ctx, nil, "hardhat_setBalance", of, hexutil.EncodeBig(amnt))
	if err != nil {
		return errors.Wrap(err, "calling hardhat_setBalance")
	}

	return nil
}

func SetStorageAt(ctx context.Context, nodeURL string, addr common.Address, idx string, val string) error {
	rpcClient, err := rpc.DialContext(ctx, nodeURL)
	if err != nil {
		return errors.Wrap(err, "creating rpc client")
	}
	defer rpcClient.Close()

	err = rpcClient.CallContext(ctx, nil, "hardhat_setStorageAt", addr.Hex(), idx, val)
	if err != nil {
		return errors.Wrap(err, "calling hardhat_setStorageAt")
	}

	return nil
}
