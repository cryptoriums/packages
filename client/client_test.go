// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package client_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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

type JsonrpcMessage struct {
	Version string          `json:"jsonrpc,omitempty"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
	Error   *JsonError      `json:"error,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
}

type JsonError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func TestNodeDenote(t *testing.T) {
	ctx, cncl := context.WithTimeout(context.Background(), 30*time.Second)
	defer cncl()

	svr1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := &JsonrpcMessage{}
		err := json.NewDecoder(r.Body).Decode(m)
		if err == io.EOF {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		fmt.Println("1", m.Method)

		switch m.Method {
		case "net_version":
			fmt.Fprintf(w, `{"jsonrpc":"2.0","id":1,"result":"1"}`)
		case "eth_gasPrice":
			w.WriteHeader(http.StatusBadRequest)
		case "eth_getBalance":
			fmt.Fprintf(w, `{"jsonrpc":"2.0","id":1,"result":"0x2"}`)
		case "eth_blockNumber":
			fmt.Fprintf(w, `{"jsonrpc":"2.0","id":1,"result":"0x9999"}`)
		}

	}))
	defer svr1.Close()

	svr2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := &JsonrpcMessage{}
		err := json.NewDecoder(r.Body).Decode(m)
		if err == io.EOF {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		fmt.Println("2", m.Method)

		switch m.Method {
		case "net_version":
			fmt.Fprintf(w, `{"jsonrpc":"2.0","id":1,"result":"1"}`)
		case "eth_gasPrice":
			fmt.Fprintf(w, `{"jsonrpc":"2.0","id":1,"result":"0x3"}`)
		case "eth_getBalance":
			fmt.Fprintf(w, `{"jsonrpc":"2.0","id":1,"result":"0x4"}`)
		case "eth_blockNumber":
			fmt.Fprintf(w, `{"jsonrpc":"2.0","id":1,"result":"0x6"}`)
		}
	}))
	defer svr2.Close()

	clt, err := client.NewClientWithRetry(ctx, log.NewNopLogger(), client.Config{}, []string{svr1.URL, svr2.URL})
	testutil.Ok(t, err)

	res, err := clt.BalanceAt(ctx, common.Address{}, nil)
	testutil.Ok(t, err)
	testutil.Equals(t, "2", res.String())

	res, err = clt.SuggestGasPrice(ctx)
	testutil.Ok(t, err)
	testutil.Equals(t, "3", res.String())

	res, err = clt.BalanceAt(ctx, common.Address{}, nil)
	testutil.Ok(t, err)
	testutil.Equals(t, "4", res.String())

	bNum, err := clt.BlockNumber(ctx)
	testutil.Ok(t, err)
	testutil.Equals(t, uint64(6), bNum)

}
