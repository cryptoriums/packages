// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package localnode

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

	"github.com/cryptoriums/packages/compiler"
	"github.com/cryptoriums/packages/logging"
	tx_p "github.com/cryptoriums/packages/tx"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	ethlogger "github.com/ethereum/go-ethereum/eth/tracers/logger"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/pkg/errors"
)

type NodeType string

var (
	Hardhat NodeType = "hardhat"
	Anvil   NodeType = "anvil"
)

const DefaultUrl = "ws://127.0.0.1:8545"

func initAccounts() []tx_p.Account {
	var Accounts []tx_p.Account
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

		acc, err := tx_p.AccountFromPrvKey(addr)
		if err != nil {
			panic(err)
		}
		Accounts = append(Accounts, acc)
	}

	return Accounts
}

type LocalNode struct {
	nodeType    NodeType
	forkNodeURL string
	cmd         *exec.Cmd
	accounts    []tx_p.Account
	rpcClient   *rpc.Client
	logger      log.Logger
}

func New(ctx context.Context, logger log.Logger, nodeType NodeType, forkNodeURL string, blockNumber string) (*LocalNode, error) {
	if forkNodeURL == "" {
		return nil, errors.Errorf("invalid forkNodeURL")
	}

	if nodeType == "" {
		nodeType = Hardhat
	}

	ln := &LocalNode{
		nodeType:    nodeType,
		forkNodeURL: forkNodeURL,
		accounts:    initAccounts(),
		logger:      logger,
	}

	var args []string

	switch ln.nodeType {
	case Hardhat:
		args = append(args, "npx", "hardhat", "node", "--fork", ln.forkNodeURL)
	case Anvil:
		args = append(args, "anvil", "--fork-url", ln.forkNodeURL)
	}

	if blockNumber != "" {
		args = append(args, "--fork-block-number", blockNumber)
	}

	level.Info(logger).Log("nodeForkArgs", strings.Join(args, " "))

	ln.cmd = fork(ln.logger, args...)

	rpcClient, err := rpc.DialContext(ctx, DefaultUrl)
	if err != nil {
		return nil, errors.Wrap(err, "creating rpc client")
	}
	ln.rpcClient = rpcClient

	if err := ln.SetNextBlockBaseFeePerGas(ctx, "0x0"); err != nil {
		return nil, err
	}

	return ln, nil
}

func (self *LocalNode) Stop() error {
	if self.rpcClient != nil {
		self.rpcClient.Close()
	}
	if self.cmd == nil {
		return errors.Errorf("no cmd found")
	}

	pgid, err := syscall.Getpgid(self.cmd.Process.Pid)
	if err != nil {
		return errors.Wrap(err, "failed to get PID")
	}

	if err := syscall.Kill(-pgid, 9); err != nil {
		return errors.Wrap(err, "failed to kill")
	}

	return nil
}

func (self *LocalNode) GetAccounts() []tx_p.Account {
	return self.accounts
}

func (self *LocalNode) GetNodeURL() string {
	return DefaultUrl
}

func fork(logger log.Logger, args ...string) *exec.Cmd {
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
		client, err := ethclient.DialContext(ctx, DefaultUrl)
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

func (self *LocalNode) Reset(ctx context.Context, blockNum uint64) error {
	if blockNum != 0 {
		opts := struct {
			Forking struct {
				JsonRpcUrl  string `json:"jsonRpcUrl"`
				BlockNumber uint64 `json:"blockNumber"`
			} `json:"forking"`
		}{
			struct {
				JsonRpcUrl  string `json:"jsonRpcUrl"`
				BlockNumber uint64 `json:"blockNumber"`
			}{
				JsonRpcUrl:  self.forkNodeURL,
				BlockNumber: blockNum,
			},
		}
		if err := self.rpcClient.CallContext(ctx, nil, string(self.nodeType)+"_reset", opts); err != nil {
			return errors.Wrap(err, "_reset")
		}
		return nil
	}

	if err := self.rpcClient.CallContext(ctx, nil, string(self.nodeType)+"_reset"); err != nil {
		return errors.Wrap(err, "_reset")
	}

	return nil
}

func (self *LocalNode) SetNextBlockBaseFeePerGas(ctx context.Context, blockBaseFee string) error {
	callSetNextBlockBaseFeePerGas := string(self.nodeType) + "_setNextBlockBaseFeePerGas"
	if err := self.rpcClient.CallContext(ctx, nil, callSetNextBlockBaseFeePerGas, blockBaseFee); err != nil {
		return errors.Wrap(err, "setNextBlockBaseFeePerGas")
	}

	return nil
}

func (self *LocalNode) ReplaceContract(ctx context.Context, contractPath string, contractName string, contractAddrToReplace common.Address) error {
	objectsDst := "tmp"

	var filePaths map[string]string
	if _, err := os.Stat(contractPath); err != nil {
		return errors.Wrap(err, "getting contract file stats")

	}

	compilerVer, err := compiler.CompilerVersion(contractPath)
	if err != nil {
		return errors.Wrap(err, "get contracts compiler version")
	}

	if compilerVer[0:1] != "v" {
		compilerVer = "v" + compilerVer
	}
	filePaths = map[string]string{
		contractPath: compilerVer,
	}

	types, abis, bins, _, _, err := compiler.GetContractObjects(filePaths, nil)
	if err != nil {
		return errors.Wrap(err, "get contracts object")
	}
	err = compiler.ExportABI(objectsDst, abis)
	if err != nil {
		return errors.Wrap(err, "Export ABI")
	}

	err = compiler.ExportBin(objectsDst, types, bins)
	if err != nil {
		return errors.Wrap(err, "Export Bins")
	}

	_bin, err := os.ReadFile(path.Join(objectsDst, contractName+".bin"))
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

	callSetCode := string(self.nodeType) + "_setCode"
	err = self.rpcClient.CallContext(ctx, nil, callSetCode, contractAddrToReplace, "0x"+bin[startDeployBin:])
	if err != nil {
		return errors.Wrapf(err, "%s", callSetCode)
	}

	return nil
}

func (self *LocalNode) Mine(ctx context.Context) error {
	if err := self.rpcClient.CallContext(ctx, nil, "evm_mine"); err != nil {
		return errors.Wrap(err, "calling evm_mine")
	}

	return nil
}

func (self *LocalNode) DisableAutoMine(ctx context.Context) error {
	err := self.rpcClient.CallContext(ctx, nil, "evm_setAutomine", false)
	if err != nil {
		return errors.Wrap(err, "calling evm_setAutomine")
	}

	return nil
}

func (self *LocalNode) SetNextBlockTimestamp(ctx context.Context, ts int64) error {
	err := self.rpcClient.CallContext(ctx, nil, "evm_setNextBlockTimestamp", big.NewInt(ts))
	if err != nil {
		return errors.Wrap(err, "calling evm_setNextBlockTimestamp")
	}

	return nil
}

func (self *LocalNode) IncreaseTime(ctx context.Context, seconds int64) error {
	err := self.rpcClient.CallContext(ctx, nil, "evm_increaseTime", big.NewInt(seconds))
	if err != nil {
		return errors.Wrap(err, "calling evm_increaseTime")
	}

	return nil
}

func (self *LocalNode) DebugTraceTransaction(ctx context.Context, hash common.Hash) (*ethlogger.ExecutionResult, error) {
	opts := struct {
		DisableMemory  bool `json:"disableMemory"`
		DisableStack   bool `json:"disableStack"`
		DisableStorage bool `json:"disableStorage"`
	}{}

	result := &ethlogger.ExecutionResult{}

	err := self.rpcClient.CallContext(ctx, result, "debug_traceTransaction", hash.Hex(), opts)
	if err != nil {
		return nil, errors.Wrapf(err, "calling %s", "debug_traceTransaction")
	}

	return result, nil
}

func (self *LocalNode) TxWithImpersonateAccountWithGasPriceAndData(ctx context.Context, from common.Address, to common.Address, gasPrice *big.Int, data []byte) (*types.Receipt, error) {
	callImpersonateAccount := string(self.nodeType) + "_impersonateAccount"
	err := self.rpcClient.CallContext(ctx, nil, callImpersonateAccount, from)
	if err != nil {
		return nil, errors.Wrapf(err, "calling %s", callImpersonateAccount)
	}

	gasPriceStr := ""
	if gasPrice != nil {
		gasPriceStr = hexutil.Encode([]byte(gasPrice.String()))
	}

	optsT := tx_p.SendTransactionOpts{
		From:     from,
		To:       to,
		Data:     hexutil.Encode(data),
		GasPrice: gasPriceStr,
	}
	var txHash string
	err = self.rpcClient.CallContext(ctx, &txHash, "eth_sendTransaction", optsT)
	if err != nil {
		return nil, errors.Wrap(err, "calling eth_sendTransaction")
	}

	callStopImpersonatingAccount := string(self.nodeType) + "_stopImpersonatingAccount"
	err = self.rpcClient.CallContext(ctx, nil, callStopImpersonatingAccount, from)
	if err != nil {
		return nil, errors.Wrapf(err, "calling %s", callStopImpersonatingAccount)
	}

	ethClient := ethclient.NewClient(self.rpcClient)
	rcpt, err := ethClient.TransactionReceipt(ctx, common.HexToHash(txHash))
	if err != nil {
		return nil, errors.Wrapf(err, "getting TX receipt:%v", common.HexToHash(txHash))
	}
	if rcpt.Status != types.ReceiptStatusSuccessful {
		return nil, errors.Errorf("TX receipt not ok:%v", rcpt.Status)
	}
	return rcpt, nil
}

func (self *LocalNode) TxWithImpersonateAccountWithData(ctx context.Context, from common.Address, to common.Address, data []byte) (*types.Receipt, error) {
	return self.TxWithImpersonateAccountWithGasPriceAndData(ctx, from, to, nil, data)

}

func (self *LocalNode) TxWithImpersonateAccountWithGasPrice(ctx context.Context, from common.Address, to common.Address, gasPrice *big.Int, abiJ string, funcName string, args ...interface{}) (*types.Receipt, error) {
	abiParsed, err := abi.JSON(strings.NewReader(abiJ))
	if err != nil {
		return nil, errors.Wrap(err, "parsing the abi")
	}
	data, err := abiParsed.Pack(funcName, args...)
	if err != nil {
		return nil, errors.Wrap(err, "packing the args")
	}

	return self.TxWithImpersonateAccountWithGasPriceAndData(ctx, from, to, gasPrice, data)
}

func (self *LocalNode) TxWithImpersonateAccount(ctx context.Context, from common.Address, to common.Address, abiJ string, funcName string, args ...interface{}) (*types.Receipt, error) {
	return self.TxWithImpersonateAccountWithGasPrice(ctx, from, to, nil, abiJ, funcName, args...)
}

func (self *LocalNode) SetBalance(ctx context.Context, of common.Address, amnt *big.Int) error {
	callSetBalance := string(self.nodeType) + "_setBalance"

	err := self.rpcClient.CallContext(ctx, nil, callSetBalance, of, hexutil.EncodeBig(amnt))
	if err != nil {
		return errors.Wrapf(err, "calling %s", callSetBalance)
	}

	return nil
}

func (self *LocalNode) SetStorageAt(ctx context.Context, addr common.Address, idx string, val string) error {
	callSetStorageAt := string(self.nodeType) + "_setStorageAt"

	err := self.rpcClient.CallContext(ctx, nil, callSetStorageAt, addr.Hex(), idx, val)
	if err != nil {
		return errors.Wrapf(err, "calling %s", callSetStorageAt)
	}

	return nil
}
