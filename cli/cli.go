// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package cli

import (
	"reflect"

	"github.com/alecthomas/kong"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

func AddressDecoder() kong.MapperFunc {
	return func(ctx *kong.DecodeContext, target reflect.Value) error {
		t, err := ctx.Scan.PopValue("address")
		if err != nil {
			return err
		}
		var addr common.Address
		switch v := t.Value.(type) {
		case string:
			if !common.IsHexAddress(v) {
				return errors.Errorf("address is not a hex string:%s", v)
			}
			addr = common.HexToAddress(v)
		default:
			return errors.Errorf("expected address string but got %q", v)
		}
		target.Set(reflect.ValueOf(addr))
		return nil
	}
}
