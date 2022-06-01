// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package big

import (
	"math/big"
)

var (
	E18 = big.NewInt(1e18)
	One = big.NewInt(1)
)

func Add(a, b *big.Int) *big.Int {
	return big.NewInt(0).Add(a, b)
}

func Sub(a, b *big.Int) *big.Int {
	return big.NewInt(0).Sub(a, b)
}

func Div(a, b *big.Int) *big.Int {
	return big.NewInt(0).Div(a, b)
}

func Mul(a, b *big.Int) *big.Int {
	return big.NewInt(0).Mul(a, b)
}

func MulWad(a, b *big.Int) *big.Int {
	if a.BitLen() == 0 {
		return a
	}
	if b.BitLen() == 0 {
		return b
	}
	c := new(big.Int).Mul(a, b)
	result := new(big.Int).Div(c, E18)

	// For really small devisions passing the e18 boundary
	// still need to return the smallest number.
	// For example the result ofdividing 1e17 by 1e18 is 0.1
	// but since big int can't represent floats this will return 0.
	// That is why in such deivisions we round up to 1
	// which represents 0.00..(18 zeros)..01
	if result.BitLen() == 0 {
		return One
	}
	return result
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
