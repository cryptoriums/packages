package head

import (
	"context"
	"errors"
	"math/big"
	"testing"
	"time"

	"github.com/cryptoriums/packages/logging"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

type testChainReader struct {
	sub *testSubscription
}

var _ ethereum.ChainReader = &testChainReader{}

func (tcr *testChainReader) BlockByHash(ctx context.Context, hash common.Hash) (block *types.Block, err error) {
	block = types.NewBlock(&types.Header{}, nil, nil, nil, nil)
	return
}

func (tcr *testChainReader) BlockByNumber(ctx context.Context, number *big.Int) (block *types.Block, err error) {
	block = types.NewBlock(&types.Header{}, nil, nil, nil, nil)
	return
}

func (tcr *testChainReader) HeaderByHash(ctx context.Context, hash common.Hash) (header *types.Header, err error) {
	header = &types.Header{}
	return
}

func (tcr *testChainReader) HeaderByNumber(ctx context.Context, number *big.Int) (header *types.Header, err error) {
	header = &types.Header{}
	return
}

func (tcr *testChainReader) TransactionCount(ctx context.Context, blockHash common.Hash) (count uint, err error) {
	count = 0
	return
}

func (tcr *testChainReader) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (tx *types.Transaction, err error) {
	return
}

func (tcr *testChainReader) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (sub ethereum.Subscription, err error) {
	tcr.sub = &testSubscription{
		ctx:       ctx,
		chHeaders: ch,
		chErrors:  make(chan error),
	}
	sub = tcr.sub
	return
}

type testSubscription struct {
	ctx       context.Context
	chHeaders chan<- *types.Header
	chErrors  chan error
}

var _ ethereum.Subscription = &testSubscription{}

func (ts *testSubscription) Unsubscribe() {
}

func (ts *testSubscription) Err() <-chan error {
	return ts.chErrors
}

func TestTrackerHead(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := logging.NewLogger()
	client := &testChainReader{}

	tracker, blocks, err := New(
		ctx,
		logger,
		client,
		0,
	)
	require.NoError(t, err)

	_ = blocks

	go func() {
		err := tracker.Start()
		require.NoError(t, err)
	}()

	go func() {
		time.Sleep(time.Second)
		sub := client.sub

		time.Sleep(time.Second)
		sub.chHeaders <- &types.Header{}

		sub.chErrors <- errors.New("random error")

		time.Sleep(time.Second)
		sub = client.sub

		time.Sleep(time.Second)
		sub.chHeaders <- &types.Header{}
		time.Sleep(time.Second)

		sub.chErrors <- errors.New("random error")

		time.Sleep(time.Second)
		sub = client.sub

		time.Sleep(time.Second)
		sub.chHeaders <- &types.Header{}
		time.Sleep(time.Second)

		cancel()
	}()

	var output []*types.Block

	func() {
		timeout := time.After(time.Second * 10)

		for {
			select {
			case <-timeout:
				t.Fatal("deadlock")
			case <-ctx.Done():
				return
			case blk := <-blocks:
				output = append(output, blk)
			}
		}
	}()

	require.Len(t, output, 3)
}
