// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package format

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_ExpandTimeVars(t *testing.T) {

	type testcase struct {
		now         string
		expectedEOD string
		expectedBOD string
	}

	cases := []testcase{
		{
			"2000-01-02 00:00:01 +0000 UTC",
			"2000-01-02 00:00:00 +0000 UTC",
			"2000-01-01 00:00:00 +0000 UTC",
		},
		{
			"2000-01-02 01:00:00 +0000 UTC",
			"2000-01-02 00:00:00 +0000 UTC",
			"2000-01-01 00:00:00 +0000 UTC",
		},
		{
			"2000-01-02 23:59:59 +0000 UTC",
			"2000-01-02 00:00:00 +0000 UTC",
			"2000-01-01 00:00:00 +0000 UTC",
		},
	}

	for _, tc := range cases {
		var n time.Time
		var err error
		now = func() time.Time {
			n, err = time.Parse("2006-01-02 15:04:05 +0000 UTC", tc.now)
			require.NoError(t, err)
			return n
		}

		ts := EOD()
		require.Equal(t, tc.expectedEOD, time.Unix(int64(ts), 0).UTC().String())

		ts = BOD()
		require.Equal(t, tc.expectedBOD, time.Unix(int64(ts), 0).UTC().String())

	}
}
