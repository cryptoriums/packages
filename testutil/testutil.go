// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package testutil

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

// TolerantVerifyLeakMain verifies go leaks but excludes the go routines that are
// launched as side effects of some of our dependencies.
func TolerantVerifyLeakMain(m *testing.M) {
	goleak.VerifyTestMain(m,
		goleak.IgnoreTopFunction("github.com/rjeczalik/notify.(*nonrecursiveTree).dispatch"),
		goleak.IgnoreTopFunction("github.com/rjeczalik/notify.(*nonrecursiveTree).internal"),
		goleak.IgnoreTopFunction("github.com/ethereum/go-ethereum/metrics.(*meterArbiter).tick"),
		goleak.IgnoreTopFunction("github.com/ethereum/go-ethereum/consensus/ethash.(*remoteSealer).loop"),
		goleak.IgnoreTopFunction("github.com/ethereum/go-ethereum/core.(*txSenderCacher).cache"),
		goleak.IgnoreTopFunction("github.com/ethereum/go-ethereum/core/state/snapshot.(*diskLayer).generate"),
		goleak.IgnoreTopFunction("github.com/ethereum/go-ethereum/accounts/abi/bind/backends.nullSubscription.func1"),
		goleak.IgnoreTopFunction("github.com/ethereum/go-ethereum/core.(*BlockChain).update"),
		goleak.IgnoreTopFunction("github.com/ethereum/go-ethereum/eth/filters.(*EventSystem).eventLoop"),
	)
}

func IsRevertErr(tb testing.TB, err error) {
	_, file, line, _ := runtime.Caller(1)

	if err == nil {
		tb.Fatalf("\033[31m%s:%d: expected an error, but got nil\033[39m\n", file, line)
	}

	errRpc, ok := errors.Cause(err).(rpc.Error)
	if !ok {
		tb.Fatalf("\033[31m%s:%d:not a rpc.Error\n\n unexpected error: %s\033[39m\n\n", file, line, err.Error())
	}

	// Internal JSON-RPC error.
	// https://www.jsonrpc.org/specification
	if errRpc.ErrorCode() != -32603 {
		tb.Fatalf("\033[31m%s:%d:not a expected revert code\n\n unexpected error: %s\033[39m\n\n", file, line, err.Error())
	}
}

func IsRevertErrWithMessage(tb testing.TB, err error, msg string) {
	_, file, line, _ := runtime.Caller(1)

	if err == nil {
		tb.Fatalf("\033[31m%s:%d\n expected error, got nothing \033[39m\n\n", file, line)
	}

	errRpc, ok := errors.Cause(err).(rpc.Error)
	if !ok {
		tb.Fatalf("\033[31m%s:%d:not a rpc.Error\n\n unexpected error: %s\033[39m\n\n", file, line, err.Error())
	}
	if errRpc.Error() != fmt.Sprintf("Error: VM Exception while processing transaction: reverted with reason string '%v'", msg) {
		tb.Fatalf("\033[31m%s:%d:not a expected revert message:"+msg+"\n\n unexpected error: %s\033[39m\n\n", file, line, err.Error())
	}
}

func OkIgnoreNotFount(tb testing.TB, err error, v ...interface{}) {
	if os.IsNotExist(err) {
		return
	}
	require.NoError(tb, err, v)
}

func KillCmd(t testing.TB, cmd *exec.Cmd) {
	pgid, err := syscall.Getpgid(cmd.Process.Pid)
	require.NoError(t, err)
	require.NoError(t, syscall.Kill(-pgid, 9))
}

func GetSimBackend(t *testing.T, sk *ecdsa.PrivateKey) *backends.SimulatedBackend {

	if sk == nil {
		var err error
		sk, err = crypto.GenerateKey()
		require.NoError(t, err)
	}

	faucetAddr := crypto.PubkeyToAddress(sk.PublicKey)
	addr := map[common.Address]core.GenesisAccount{
		common.BytesToAddress([]byte{1}): {Balance: big.NewInt(1)}, // ECRecover
		common.BytesToAddress([]byte{2}): {Balance: big.NewInt(1)}, // SHA256
		common.BytesToAddress([]byte{3}): {Balance: big.NewInt(1)}, // RIPEMD
		common.BytesToAddress([]byte{4}): {Balance: big.NewInt(1)}, // Identity
		common.BytesToAddress([]byte{5}): {Balance: big.NewInt(1)}, // ModExp
		common.BytesToAddress([]byte{6}): {Balance: big.NewInt(1)}, // ECAdd
		common.BytesToAddress([]byte{7}): {Balance: big.NewInt(1)}, // ECScalarMul
		common.BytesToAddress([]byte{8}): {Balance: big.NewInt(1)}, // ECPairing
		faucetAddr:                       {Balance: new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(9))},
	}
	alloc := core.GenesisAlloc(addr)
	return backends.NewSimulatedBackend(alloc, 80000000)
}

func SkipCI(t *testing.T) {
	if os.Getenv("GITHUB_ACTIONS") != "" {
		t.Skip("Skipping testing in GH Actions environment")
	}
}
