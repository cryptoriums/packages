// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package events

import (
	"context"
	"os"
	"testing"

	"github.com/cryptoriums/packages/constants"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-kit/log"
	"github.com/pkg/errors"
)

func TestSubs(t *testing.T) {
	ctx := context.Background()
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))

	eventsTrackerTransfer, eventsTransferOutput, err := New(
		ctx,
		logger,
		Config{LogLevel: "info"},
		client,
		[]common.Address{master.Addr()},
		0,
		0,
		[][]interface{}{{master.Abi().Events[contracts.EventNameTransfer].ID}},
		constants.ReorgEventWaitFast,
	)
	if err != nil {
		return nil, errors.Wrap(err, "creating events tracker")
	}
}
