// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package math

import (
	"math"
	"math/big"
)

var (
	E18 = big.NewInt(1e18)
	One = big.NewInt(1)
)

func MulWad(a, b *big.Int) *big.Int {
	if a.BitLen() == 0 {
		return a
	}
	if b.BitLen() == 0 {
		return b
	}
	c := new(big.Int).Mul(a, b)
	result := new(big.Int).Div(c, E18)

	if result.BitLen() == 0 {
		return One
	}
	return result
}

func PercentageDiff(old, new float64) float64 {
	diff := new - old

	if old == 0 {
		if new > 0 {
			return math.MaxFloat64
		}
		return -math.MaxFloat64
	}
	if new == 0 {
		if old > 0 {
			return math.MaxFloat64
		}
		return -math.MaxFloat64
	}

	if old > new {
		return -math.Abs((diff / new) * 100)
	}
	return math.Abs((diff / old) * 100)

}

func BigIntToFloat(input *big.Int) float64 {
	fl, _ := big.NewFloat(0).SetInt(input).Float64()
	return fl
}

func BigIntToFloatDiv(input *big.Int, devider float64) float64 {
	fl := BigIntToFloat(input)
	if devider == 1 {
		return fl
	}
	f := 0.0
	if input != nil {
		return fl / devider
	}
	return f
}

func FloatToBigIntMul(input float64, multiplier float64) *big.Int {
	if input == 0 {
		return big.NewInt(0)
	}
	bigE18 := big.NewFloat(0).Mul(big.NewFloat(input), big.NewFloat(multiplier))
	result, _ := bigE18.Int(nil)
	return result
}

func FloatToBigInt(input float64) *big.Int {
	result, _ := big.NewFloat(input).Int(nil)
	return result
}

// ConfidenceInDifference calculates the percentage difference between the max and min and subtract this from 100%.
// Example:
// min 1, max 2
// Difference is 1 which is 100% so the final confidence is 100-100 equals 0%.
func ConfidenceInDifference(min, max float64) float64 {
	return 100 - (math.Abs(min-max)/min)*100
}
