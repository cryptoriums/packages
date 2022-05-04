// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package env

import (
	"strings"
	"testing"

	"github.com/cryptoriums/packages/testutil"
)

func FuzzEncryptDecrypt(f *testing.F) {
	testcases := []struct {
		Pass, Input string
	}{
		{Pass: "aaa", Input: "Hello, world"},
	}
	for _, tc := range testcases {
		f.Add(tc.Input, tc.Pass)
	}
	f.Fuzz(func(t *testing.T, input, pass string) {
		input = strings.ReplaceAll(input, EncryptIndicator, "")
		encr, err := Encrypt(input, pass)
		testutil.Ok(t, err, input)
		testutil.Assert(t, IsEncrypted(encr))
		decr, err := Decrypt(encr, pass)
		testutil.Ok(t, err, input)
		testutil.Equals(t, decr, input)
	})
}
