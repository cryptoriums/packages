// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package big

import (
	"fmt"
	"math/big"
	"strconv"
	"testing"

	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	type testcase struct {
		a   *big.Int
		b   []*big.Int
		exp *big.Int
	}

	cases := []testcase{
		{
			big.NewInt(0),
			[]*big.Int{
				big.NewInt(1e18),
			},
			big.NewInt(1e18),
		},
		{

			big.NewInt(1e18),
			[]*big.Int{
				big.NewInt(1e18),
			},
			big.NewInt(2e18),
		},
		{

			big.NewInt(1e18),
			[]*big.Int{
				big.NewInt(1e18),
				big.NewInt(1e18),
			},
			big.NewInt(3e18),
		},
		{

			big.NewInt(0).Mul(big.NewInt(2), big.NewInt(3)), // Weird bug that caused the original number to be modified.
			[]*big.Int{
				big.NewInt(21000),
			},
			big.NewInt(21006),
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			originalA := tc.a.String()

			var originalB []*big.Int
			for _, b := range tc.b {
				originalB = append(originalB, big.NewInt(0).SetBits(b.Bits()))
			}
			require.Equal(t, tc.exp, Add(tc.a, tc.b...))

			// Verify that the original numbers are not modified.
			require.Equal(t, originalA, tc.a.String())
			for i, b := range originalB {
				require.Equal(t, b, tc.b[i])
			}
		})
	}
}

func TestMulWad(t *testing.T) {
	type testcase struct {
		a   *big.Int
		b   *big.Int
		exp *big.Int
	}

	e16, ok := big.NewInt(0).SetString("10000000000000000", 10)
	require.True(t, ok)
	e17, ok := big.NewInt(0).SetString("100000000000000000", 10)
	require.True(t, ok)
	e18, ok := big.NewInt(0).SetString("1000000000000000000", 10)
	require.True(t, ok)
	e19, ok := big.NewInt(0).SetString("10000000000000000000", 10)
	require.True(t, ok)
	e20, ok := big.NewInt(0).SetString("100000000000000000000", 10)
	require.True(t, ok)
	e22, ok := big.NewInt(0).SetString("10000000000000000000000", 10)
	require.True(t, ok)

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
			require.Equal(t, tc.exp, MulWad(tc.a, tc.b))

		})
	}
}

func TestFromFloatMul(t *testing.T) {
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
		{ // Some extremely large number shouldn't loose precision.
			8272739999999999549999999,
			1,
			8272739999999999549999999,
		},
		{ // Some extremely small number shouldn't loose precision.
			0.0000000000000000000000001234,
			1e28,
			1234,
		},
		{ // Some realistic eth price.
			1296.890000,
			1e18,
			1296890000000000196608, // For some reason there is some precision loss?
		},
	}

	for i, tc := range cases {
		exp, ok := big.NewInt(0).SetString(fmt.Sprintf("%.0f", tc.expected), 10)
		require.True(t, ok, "failed to convert string to big int")

		act := FromFloatMul(tc.input, tc.multiplier)
		require.Equal(t, exp, act, "Case:"+strconv.Itoa(i))
	}
}

func TestToFloatDiv(t *testing.T) {
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
		require.True(t, ok, "failed to convert string to big int")

		act := ToFloatDiv(input, tc.devider)
		require.Equal(t, tc.expected, act, "Case:"+strconv.Itoa(i))
	}
}

func TestMulFloat(t *testing.T) {
	testCases := []struct {
		name       string
		input      *big.Int
		multiplier float64
		expected   *big.Int
	}{
		{
			name:       "Case 1: Multiply 200 by 0.20",
			input:      big.NewInt(200),
			multiplier: 0.20,
			expected:   big.NewInt(40),
		},
		{
			name:       "Case 2: Multiply 100 by 0.75",
			input:      big.NewInt(100),
			multiplier: 0.75,
			expected:   big.NewInt(75),
		},
		{
			name:       "Case 3: Multiply 50 by 1.5",
			input:      big.NewInt(50),
			multiplier: 1.5,
			expected:   big.NewInt(75),
		},
		{
			name:       "Case 4: Multiply 200 by 0.50",
			input:      big.NewInt(200),
			multiplier: 0.50,
			expected:   big.NewInt(100),
		}, {
			name:       "Case 5: Some rounding error",
			input:      big.NewInt(100),
			multiplier: 0.333,
			expected:   big.NewInt(33),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := MulFloat(tc.input, tc.multiplier)
			require.Equal(t, tc.expected, result, "The result of the multiplication should be equal to the expected result.")
		})
	}
}

func TestPercentage(t *testing.T) {
	type testcase struct {
		value      *big.Int
		percentage int32
		exp        *big.Int
	}

	cases := []testcase{
		{
			value:      big.NewInt(100),
			percentage: 10,
			exp:        big.NewInt(10),
		},
		{
			value:      big.NewInt(200),
			percentage: 25,
			exp:        big.NewInt(50),
		},
		{
			value:      big.NewInt(333),
			percentage: 33,
			exp:        big.NewInt(109),
		},
		{
			value:      big.NewInt(500),
			percentage: 99,
			exp:        big.NewInt(495),
		},
		{
			value:      big.NewInt(100),
			percentage: 99,
			exp:        big.NewInt(99),
		},
	}

	for i, tc := range cases {
		t.Run("case_"+strconv.Itoa(i), func(t *testing.T) {
			result := Percentage(tc.value, tc.percentage)
			require.Equal(t, tc.exp, result)
		})
	}
}
