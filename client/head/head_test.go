// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package head

import (
	"testing"

	"github.com/cryptoriums/packages/logging"
	"github.com/cryptoriums/packages/testutil"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
)

// TestHeadSubscriberWithRedundancy_SameHeaders ensures that events are deduplicated properly when using multiple backends.
//
// Here is an example how to use with node urls.
//
// infura, err := ethclient.Dial("wss://infura.io/......")
// alchemy, err := ethclient.Dial("wss://alchemyapi.io/......")
//
// client := NewHeadSubscriberWithRedundancy(ctx, logger, []HeadSubscriber{infura, alchemy})
// client.SubscribeNewHead(ctx, ch)
// .
func TestHeadSubscriberWithRedundancy_SameHeaders(t *testing.T) {
	logger := logging.NewLogger()
	logger, err := logging.ApplyFilter("debug", logger)
	require.NoError(t, err)

	ctx := context.Background()

	sk, err := crypto.GenerateKey()
	require.NoError(t, err)

	// All backends have the same headers.
	{
		var (
			headSubscribers []HeadSubscriber
			backends        []*backends.SimulatedBackend
		)
		for i := 0; i < 4; i++ {
			backend := testutil.GetSimBackend(t, sk)
			headSubscribers = append(headSubscribers, backend)
			backends = append(backends, backend)
		}

		chExp := make(chan *types.Header)
		chAct := make(chan *types.Header)

		headSubscriber := NewHeadSubscriberWithRedundancy(logger, headSubscribers)

		subsExp, err := backends[0].SubscribeNewHead(ctx, chExp)
		require.NoError(t, err)
		subsAct, err := headSubscriber.SubscribeNewHead(ctx, chAct)
		require.NoError(t, err)

		for _, b := range backends {
			b.Commit()
		}

		var headersExp []*types.Header
		var headersAct []*types.Header
		headersExp = append(headersExp, <-chExp)
		headersAct = append(headersAct, <-chAct)

		require.Equal(t, headersExp, headersAct)

		select {
		case header := <-chAct:
			t.Fatalf("there is an extra header:%+v", header)
		default:
		}

		subsExp.Unsubscribe()
		subsAct.Unsubscribe()
	}
}

// TestHeadSubscriberWithRedundancy_MultiCallsToSubscribeNewHead ensures that
// multi calls to SubscribeNewHead gets a new cache and send subs to all subscribers.
func TestHeadSubscriberWithRedundancy_MultiCallsToSubscribeNewHead(t *testing.T) {
	logger := logging.NewLogger()
	logger, err := logging.ApplyFilter("debug", logger)
	require.NoError(t, err)

	ctx := context.Background()

	sk, err := crypto.GenerateKey()
	require.NoError(t, err)

	backend := testutil.GetSimBackend(t, sk)

	chExp := make(chan *types.Header)
	chAct1 := make(chan *types.Header)
	chAct2 := make(chan *types.Header)

	headSubscriber := NewHeadSubscriberWithRedundancy(logger, []HeadSubscriber{backend})

	subsExp, err := backend.SubscribeNewHead(ctx, chExp)
	require.NoError(t, err)
	subsAct1, err := headSubscriber.SubscribeNewHead(ctx, chAct1)
	require.NoError(t, err)
	subsAct2, err := headSubscriber.SubscribeNewHead(ctx, chAct2)
	require.NoError(t, err)

	for i := 0; i < 100; i++ {
		backend.Commit()

		var headersExp []*types.Header
		var headersAct1 []*types.Header
		var headersAct2 []*types.Header
		headersExp = append(headersExp, <-chExp)
		headersAct1 = append(headersAct1, <-chAct1)
		headersAct2 = append(headersAct2, <-chAct2)

		require.Equal(t, headersExp, headersAct1)
		require.Equal(t, headersExp, headersAct2)

		select {
		case header := <-chAct1:
			t.Fatalf("there is an extra header:%+v", header)
		case header := <-chAct2:
			t.Fatalf("there is an extra header:%+v", header)
		default:
		}
	}

	subsExp.Unsubscribe()
	subsAct1.Unsubscribe()
	subsAct2.Unsubscribe()
}

func TestHeadSubscriberWithRedundancy_DifferentHeaders(t *testing.T) {
	logger := logging.NewLogger()
	logger, err := logging.ApplyFilter("debug", logger)
	require.NoError(t, err)

	ctx := context.Background()

	sk1, err := crypto.GenerateKey()
	require.NoError(t, err)
	backend1 := testutil.GetSimBackend(t, sk1)

	sk2, err := crypto.GenerateKey()
	require.NoError(t, err)
	backend2 := testutil.GetSimBackend(t, sk2)

	headSubscriber := NewHeadSubscriberWithRedundancy(logger, []HeadSubscriber{backend1, backend2})

	chExp1 := make(chan *types.Header)
	chExp2 := make(chan *types.Header)
	chAct := make(chan *types.Header)

	subsExp1, err := backend1.SubscribeNewHead(ctx, chExp1)
	require.NoError(t, err)
	subsExp2, err := backend2.SubscribeNewHead(ctx, chExp2)
	require.NoError(t, err)
	subsAct, err := headSubscriber.SubscribeNewHead(ctx, chAct)
	require.NoError(t, err)

	var headersExp []*types.Header
	var headersAct []*types.Header

	// The order of reading here matters to ensure that the
	// logs arrive in the same order and
	// also that the multi subscriber doesn't block the simulated backend sending.
	backend1.Commit()
	headersExp = append(headersExp, <-chExp1)
	headersAct = append(headersAct, <-chAct)

	backend2.Commit()
	headersExp = append(headersExp, <-chExp2)
	headersAct = append(headersAct, <-chAct)

	require.Equal(t, headersExp, headersAct)

	select {
	case header := <-chAct:
		t.Fatalf("there is an extra header:%+v", header)
	default:
	}

	subsExp1.Unsubscribe()
	subsExp2.Unsubscribe()
	subsAct.Unsubscribe()
}

func TestHeadSubscriberWithRedundancy_OneHasExtraHeader(t *testing.T) {
	logger := logging.NewLogger()
	logger, err := logging.ApplyFilter("debug", logger)
	require.NoError(t, err)

	ctx := context.Background()

	sk1, err := crypto.GenerateKey()
	require.NoError(t, err)
	backend1 := testutil.GetSimBackend(t, sk1)

	sk2, err := crypto.GenerateKey()
	require.NoError(t, err)
	backend2 := testutil.GetSimBackend(t, sk2)

	headSubscriber := NewHeadSubscriberWithRedundancy(logger, []HeadSubscriber{backend1, backend2})

	chExp1 := make(chan *types.Header)
	chExp2 := make(chan *types.Header)
	chAct := make(chan *types.Header)

	subsExp1, err := backend1.SubscribeNewHead(ctx, chExp1)
	require.NoError(t, err)

	subsExp2, err := backend2.SubscribeNewHead(ctx, chExp2)
	require.NoError(t, err)

	subsAct, err := headSubscriber.SubscribeNewHead(ctx, chAct)
	require.NoError(t, err)

	var headersExp []*types.Header
	var headersAct []*types.Header

	// The order here matters to ensure that the
	// logs arrive in the same order and
	// also that the multi subscriber doesn't block the simulated backend sending.
	backend1.Commit()
	headersExp = append(headersExp, <-chExp1)
	headersAct = append(headersAct, <-chAct)

	backend2.Commit()
	headersExp = append(headersExp, <-chExp2)
	headersAct = append(headersAct, <-chAct)

	backend2.Commit()
	headersExp = append(headersExp, <-chExp2)
	headersAct = append(headersAct, <-chAct)

	require.Equal(t, headersExp, headersAct)

	select {
	case header := <-chAct:
		t.Fatalf("there is an extra header:%+v", header)
	default:
	}

	subsExp1.Unsubscribe()
	subsExp2.Unsubscribe()
	subsAct.Unsubscribe()
}

// TestHeadSubscriberWithRedundancy_ErrCh ensures that
// the subs.Err() func receives an error even when one of the Headsubscribes returns an error.
func TestHeadSubscriberWithRedundancy_ErrCh(t *testing.T) {
	logger := logging.NewLogger()
	logger, err := logging.ApplyFilter("debug", logger)
	require.NoError(t, err)

	ctx := context.Background()

	errSent := make(chan struct{})
	headSubscriber := NewHeadSubscriberWithRedundancy(logger, []HeadSubscriber{NewBackendSimulateErr(errSent), testutil.GetSimBackend(t, nil)})

	subs, err := headSubscriber.SubscribeNewHead(ctx, nil)
	require.NoError(t, err)

	<-errSent
	select {
	case <-subs.Err():
	default:
		t.Fatalf("no error was received")
	}

}

func NewBackendSimulateErr(sent chan struct{}) *BackendSimulateErr {
	e := make(chan error)
	go func() {
		e <- errors.New("xxx")
		sent <- struct{}{}

	}()
	return &BackendSimulateErr{
		err: e,
	}
}

type BackendSimulateErr struct {
	err <-chan error
}

func (self *BackendSimulateErr) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	return nil, nil
}

func (self *BackendSimulateErr) SubscribeNewHead(ctx context.Context, chDst chan<- *types.Header) (ethereum.Subscription, error) {
	return self, nil
}

func (self *BackendSimulateErr) Err() <-chan error {
	return self.err
}

func (self *BackendSimulateErr) Unsubscribe() {

}
