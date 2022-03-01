// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package format

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Unfortunate hack to enable json parsing of human readable time strings
// see https://github.com/golang/go/issues/10275
// code from https://stackoverflow.com/questions/48050945/how-to-unmarshal-json-into-durations.
type Duration struct {
	time.Duration
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return []byte("\"" + d.String() + "\""), nil
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		d.Duration = time.Duration(value * float64(time.Second))
		return nil
	case string:
		dur, err := time.ParseDuration(value)
		if err != nil {
			return err
		}
		d.Duration = dur
		return nil
	default:
		return errors.Errorf("invalid duration")
	}
}

var (
	now = time.Now // Set as a var to allow overriding in the tests.
)

func EOD() int64 {
	year, month, day := now().Date()
	eod := time.Date(year, month, day, 0, 0, 0, 0, now().Location())

	return eod.Unix()
}

func BOD() int64 {
	secsInAday := int64(86400)
	return EOD() - secsInAday
}
