// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package http

import (
	"strconv"
	"testing"
	"time"

	"github.com/cryptoriums/packages/testutil"
)

func Test_ExpandTimeVars(t *testing.T) {

	type testcase struct {
		input          string
		expectedOutput string
	}

	cases := []testcase{
		{
			"https://apps.bea.gov/api/data/?&UserID=92D67E4B-61D9-4614-859F-95741E8CB2D3&method=GetData&DataSetName=NIPA&TableName=T20804&Frequency=M&Year=$YEAR&ResultFormat=json",
			"https://apps.bea.gov/api/data/?&UserID=92D67E4B-61D9-4614-859F-95741E8CB2D3&method=GetData&DataSetName=NIPA&TableName=T20804&Frequency=M&Year=" + strconv.Itoa(time.Now().Year()) + "&ResultFormat=json",
		},
	}

	for _, tc := range cases {
		testutil.Equals(t, tc.expectedOutput, ExpandTimeVars(tc.input))

	}
}
