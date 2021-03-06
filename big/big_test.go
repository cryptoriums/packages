// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package big

import (
	"fmt"
	"math/big"
	"strconv"
	"testing"

	"github.com/cryptoriums/packages/testutil"
	"github.com/ethereum/go-ethereum/params"
)

func TestMulWad(t *testing.T) {
	type testcase struct {
		a   *big.Int
		b   *big.Int
		exp *big.Int
	}

	e16, ok := big.NewInt(0).SetString("10000000000000000", 10)
	testutil.Assert(t, ok)
	e17, ok := big.NewInt(0).SetString("100000000000000000", 10)
	testutil.Assert(t, ok)
	e18, ok := big.NewInt(0).SetString("1000000000000000000", 10)
	testutil.Assert(t, ok)
	e19, ok := big.NewInt(0).SetString("10000000000000000000", 10)
	testutil.Assert(t, ok)
	e20, ok := big.NewInt(0).SetString("100000000000000000000", 10)
	testutil.Assert(t, ok)
	e22, ok := big.NewInt(0).SetString("10000000000000000000000", 10)
	testutil.Assert(t, ok)

	cases := []testcase{
		{
			big.NewInt(0),
			big.NewInt(1e18),
			big.NewInt(0),
		}, {
			big.NewInt(1e18),
			big.NewInt(0),
			big.NewInt(0),
		},
		{
			big.NewInt(1e18),
			big.NewInt(1),
			big.NewInt(1),
		},
		{
			big.NewInt(1111111111111111111),
			big.NewInt(2000000000000000000),
			big.NewInt(2222222222222222222),
		},
		{
			big.NewInt(1e7),
			big.NewInt(1e6),
			big.NewInt(1),
		},
		{
			e17,
			e17,
			e16,
		},
		{
			e18,
			e17,
			e17,
		}, {
			e18,
			e19,
			e19,
		}, {
			e20,
			e20,
			e22,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			testutil.Equals(t, tc.exp, MulWad(tc.a, tc.b))

		})
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
		{
			5e9,
			params.Ether,
			0.000000005,
		},
	}

	for i, tc := range cases {
		input, ok := big.NewInt(0).SetString(fmt.Sprintf("%.0f", tc.input), 10)
		testutil.Assert(t, ok, "failed to convert string to big int")

		act := BigIntToFloatDiv(input, tc.devider)
		testutil.Equals(t, tc.expected, act, "Case:"+strconv.Itoa(i))
	}
}
