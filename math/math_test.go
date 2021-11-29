// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package math

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"testing"

	"github.com/cryptoriums/packages/testutil"
	"github.com/ethereum/go-ethereum/params"
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

func TestFloatToBigIntMul(t *testing.T) {
	type testcase struct {
		input      float64
		multiplier float64
		expected   float64
	}

	cases := []testcase{
		{
			1,
			params.Ether,
			params.Ether,
		},
		{
			10,
			params.Ether,
			params.Ether * 10,
		},
		{
			10000,
			params.Ether,
			params.Ether * 10000,
		},
		{
			0.1,
			params.Ether,
			params.Ether / 10,
		},
		{
			0.01,
			params.Ether,
			params.Ether / 100,
		},
		{
			1,
			params.GWei,
			params.GWei,
		},
		{
			0.1,
			params.GWei,
			params.GWei / 10,
		},
	}

	for i, tc := range cases {
		act := FloatToBigIntMul(tc.input, tc.multiplier)
		exp, ok := big.NewInt(0).SetString(fmt.Sprintf("%.0f", tc.expected), 10)
		testutil.Assert(t, ok, "failed to convert string to big int")

		testutil.Equals(t, exp, act, "Case:"+strconv.Itoa(i))
	}
}

func TestBigIntToFloatDiv(t *testing.T) {
	type testcase struct {
		input    float64
		devider  float64
		expected float64
	}

	cases := []testcase{
		{
			params.Ether,
			params.Ether,
			1,
		},
		{
			10 * params.Ether,
			params.Ether,
			10,
		},
		{
			params.Ether / 10,
			params.Ether,
			0.1,
		},
		{
			params.Ether / 100,
			params.Ether,
			0.01,
		},
	}

	for i, tc := range cases {
		input, ok := big.NewInt(0).SetString(fmt.Sprintf("%.0f", tc.input), 10)
		testutil.Assert(t, ok, "failed to convert string to big int")

		act := BigIntToFloatDiv(input, tc.devider)
		testutil.Equals(t, tc.expected, act, "Case:"+strconv.Itoa(i))
	}
}
