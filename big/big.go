// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package big

import (
	"math/big"

	"github.com/ethereum/go-ethereum/params"
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
	return big.NewInt(0).Div(big.NewInt(0).Mul(a, b), big.NewInt(params.Ether))
}
