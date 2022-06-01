// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package math

import (
	"math"
	"strconv"
	"testing"

	"github.com/cryptoriums/packages/testutil"
)

func TestPercentageDiff(t *testing.T) {
	type testcase struct {
		old      float64
		new      float64
		expected float64
	}

	cases := []testcase{
		{
			1,
			10,
			900,
		},
		{
			10,
			1,
			-900,
		},
		{
			0.01,
			0.1,
			900,
		},
		{
			0.1,
			0.01,
			-900,
		},
		{
			0,
			1,
			math.MaxFloat64,
		},
		{
			0,
			-1,
			-math.MaxFloat64,
		},
		{
			1,
			0,
			math.MaxFloat64,
		},
		{
			-1,
			0,
			-math.MaxFloat64,
		},
		{
			1,
			-1,
			-200,
		},
	}

	for i, tc := range cases {

		testutil.Equals(t, tc.expected, PercentageDiff(tc.old, tc.new), "Case:"+strconv.Itoa(i))
	}
}
