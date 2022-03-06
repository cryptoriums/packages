package big

import (
	"math/big"
	"strconv"
	"testing"

	"github.com/cryptoriums/packages/testutil"
	"github.com/ethereum/go-ethereum/params"
)

func TestBigMulWad(t *testing.T) {
	type testcase struct {
		a   *big.Int
		b   *big.Int
		exp *big.Int
	}

	cases := []testcase{
		{
			big.NewInt(params.Ether),
			big.NewInt(params.GWei),
			big.NewInt(params.GWei),
		},
		{
			big.NewInt(params.Ether),
			big.NewInt(params.Ether),
			big.NewInt(params.Ether),
		},
		{
			big.NewInt(params.Ether),
			big.NewInt(params.Wei),
			big.NewInt(params.Wei),
		},
		{
			big.NewInt(params.Ether),
			big.NewInt(params.Ether / 2),
			big.NewInt(params.Ether / 2),
		},
	}

	for i, tc := range cases {
		testutil.Equals(t, tc.exp, MulWad(tc.a, tc.b), "Case:"+strconv.Itoa(i))
	}
}
