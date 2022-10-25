// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package testutil

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"syscall"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
	"github.com/pmezard/go-difflib/difflib"
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

// Assert fails the test if the condition is false.
func Assert(tb testing.TB, condition bool, v ...interface{}) {
	tb.Helper()
	if condition {
		return
	}
	_, file, line, _ := runtime.Caller(1)

	var msg string
	if len(v) > 0 {
		msg = fmt.Sprintf(v[0].(string), v[1:]...)
	}
	tb.Fatalf("\033[31m%s:%d: "+msg+"\033[39m\n\n", file, line)
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

// Ok fails the test if an err is not nil.
func Ok(tb testing.TB, err error, v ...interface{}) {
	tb.Helper()
	if err == nil {
		return
	}
	_, file, line, _ := runtime.Caller(1)

	var msg string
	if len(v) > 0 {
		msg = fmt.Sprintf(v[0].(string), v[1:]...)
	}
	tb.Fatalf("\033[31m%s:%d:"+msg+"\n\n unexpected error: %s\033[39m\n\n", file, line, err.Error())
}

func OkIgnoreNotFount(tb testing.TB, err error, v ...interface{}) {
	if os.IsNotExist(err) {
		return
	}
	Ok(tb, err, v)
}

func KillCmd(t testing.TB, cmd *exec.Cmd) {
	pgid, err := syscall.Getpgid(cmd.Process.Pid)
	Ok(t, err)
	Ok(t, syscall.Kill(-pgid, 9))
}

// NotOk fails the test if an err is nil.
func NotOk(tb testing.TB, err error, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)

	var msg string
	if len(v) > 0 {
		msg = fmt.Sprintf(v[0].(string), v[1:]...)
	}

	tb.Fatalf("\033[31m%s:%d:"+msg+"\n\n expected error, got nothing \033[39m\n\n", file, line)
}

// Equals fails the test if exp is not equal to act.
func Equals(tb testing.TB, exp, act interface{}, v ...interface{}) {
	tb.Helper()
	if reflect.DeepEqual(exp, act) {
		return
	}
	_, file, line, _ := runtime.Caller(1)

	var msg string
	if len(v) > 0 {
		msg = fmt.Sprintf(v[0].(string), v[1:]...)
	}
	tb.Fatal(sprintfWithLimit("\033[31m%s:%d:"+msg+"\n\n\texp: %#v\n\n\tgot: %#v%s\033[39m\n\n", file, line, exp, act, diff(exp, act)))
}

func sprintfWithLimit(act string, v ...interface{}) string {
	s := fmt.Sprintf(act, v...)
	if len(s) > 10000 {
		return s[:10000] + "...(output trimmed)"
	}
	return s
}

func typeAndKind(v interface{}) (reflect.Type, reflect.Kind) {
	t := reflect.TypeOf(v)
	k := t.Kind()

	if k == reflect.Ptr {
		t = t.Elem()
		k = t.Kind()
	}
	return t, k
}

// diff returns a diff of both values as long as both are of the same type and
// are a struct, map, slice, array or string. Otherwise it returns an empty string.
func diff(expected interface{}, actual interface{}) string {
	if expected == nil || actual == nil {
		return ""
	}

	et, ek := typeAndKind(expected)
	at, _ := typeAndKind(actual)
	if et != at {
		return ""
	}

	if ek != reflect.Struct && ek != reflect.Map && ek != reflect.Slice && ek != reflect.Array && ek != reflect.String {
		return ""
	}

	var e, a string
	c := spew.ConfigState{
		Indent:                  " ",
		DisablePointerAddresses: true,
		DisableCapacities:       true,
		SortKeys:                true,
	}
	if et != reflect.TypeOf("") {
		e = c.Sdump(expected)
		a = c.Sdump(actual)
	} else {
		e = reflect.ValueOf(expected).String()
		a = reflect.ValueOf(actual).String()
	}

	diff, _ := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
		A:        difflib.SplitLines(e),
		B:        difflib.SplitLines(a),
		FromFile: "Expected",
		FromDate: "",
		ToFile:   "Actual",
		ToDate:   "",
		Context:  1,
	})
	return "\n\nDiff:\n" + diff
}

func GetSimBackend(t *testing.T, sk *ecdsa.PrivateKey) *backends.SimulatedBackend {

	if sk == nil {
		var err error
		sk, err = crypto.GenerateKey()
		Ok(t, err)
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
