// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package testutil

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"syscall"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
	"github.com/pmezard/go-difflib/difflib"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/prometheus/model/labels"
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
	tb.Fatalf("\033[31m%s:%d: "+msg+"\033[39m\n\n", filepath.Base(file), line)
}

func IsRevertErrWithMessage(tb testing.TB, err error, msg string) {
	_, file, line, _ := runtime.Caller(1)

	errRpc, ok := errors.Cause(err).(rpc.Error)
	if !ok {
		tb.Fatalf("\033[31m%s:%d:not a rpc.Error\n\n unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
	}

	// Internal JSON-RPC error.
	// https://www.jsonrpc.org/specification
	if errRpc.ErrorCode() != -32603 {
		tb.Fatalf("\033[31m%s:%d:not a expected revert code\n\n unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
	}

	if msg != "" {
		if errRpc.Error() != fmt.Sprintf("Error: VM Exception while processing transaction: reverted with reason string '%v'", msg) {
			tb.Fatalf("\033[31m%s:%d:not a expected revert message"+msg+"\n\n unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		}
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
	tb.Fatalf("\033[31m%s:%d:"+msg+"\n\n unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
}

func OkIgnoreNotFount(tb testing.TB, err error, v ...interface{}) {
	if os.IsNotExist(err) {
		return
	}
	Ok(tb, err, v)
}

func HardhatFork(t testing.TB, args ...string) *exec.Cmd {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	cmdReaderStdOut, err := cmd.StdoutPipe()
	Ok(t, err)
	cmdReaderStdErr, err := cmd.StderrPipe()
	Ok(t, err)

	go func() {
		scanner := bufio.NewScanner(cmdReaderStdOut)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			t.Log(scanner.Text())
		}
	}()

	go func() {
		scanner := bufio.NewScanner(cmdReaderStdErr)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			panic(scanner.Text())
		}
	}()

	Ok(t, cmd.Start())

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
		t.Log("error connecting will retry")
		time.Sleep(time.Second)
	}
	return cmd
}

func KillCmd(t testing.TB, cmd *exec.Cmd) {
	pgid, err := syscall.Getpgid(cmd.Process.Pid)
	Ok(t, err)
	Ok(t, syscall.Kill(-pgid, 9))
}

// NotOk fails the test if an err is nil.
func NotOk(tb testing.TB, err error, v ...interface{}) {
	tb.Helper()
	if err != nil {
		return
	}
	_, file, line, _ := runtime.Caller(1)

	var msg string
	if len(v) > 0 {
		msg = fmt.Sprintf(v[0].(string), v[1:]...)
	}
	tb.Fatalf("\033[31m%s:%d:"+msg+"\n\n expected error, got nothing \033[39m\n\n", filepath.Base(file), line)
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
	tb.Fatal(sprintfWithLimit("\033[31m%s:%d:"+msg+"\n\n\texp: %#v\n\n\tgot: %#v%s\033[39m\n\n", filepath.Base(file), line, exp, act, diff(exp, act)))
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

func ToFloat64(c prometheus.Collector, labelsToMatch ...labels.Label) float64 {
	var (
		m     prometheus.Metric
		mChan = make(chan prometheus.Metric)
		done  = make(chan struct{})
	)

	pb := &dto.Metric{}
	go func() {
	Main:
		for m = range mChan {
			_pb := &dto.Metric{}
			if err := m.Write(_pb); err != nil {
				panic(errors.Wrap(err, "writing into the dto metric"))
			}

			if len(labelsToMatch) == 0 {
				pb = _pb
				continue Main
			}
			for i, l := range labelsToMatch {
				if *_pb.Label[i].Name != l.Name || *_pb.Label[i].Value != l.Value {
					continue Main
				}
			}
			pb = _pb
		}
		close(done)
	}()

	c.Collect(mChan)
	close(mChan)
	<-done

	if pb.Gauge != nil {
		return pb.Gauge.GetValue()
	}
	if pb.Counter != nil {
		return pb.Counter.GetValue()
	}
	if pb.Untyped != nil {
		return pb.Untyped.GetValue()
	}
	panic(fmt.Errorf("collected a non-gauge/counter/untyped metric: %s", pb))
}
