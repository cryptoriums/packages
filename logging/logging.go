// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package logging

import (
	logstd "log"
	"os"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/pkg/errors"
)

const DefaultTimeFormat = "02 Jan 15:04:05.00"

// NewLogger create a new logger.
func NewLogger() log.Logger {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	return log.With(logger, "ts", log.TimestampFormat(func() time.Time { return time.Now().UTC() }, DefaultTimeFormat), "caller", log.Caller(5))
}

// ApplyFilter applies a filter to logger based on component name.
func ApplyFilter(configLevel string, logger log.Logger) (log.Logger, error) {
	var lvl level.Option
	switch configLevel {
	case "":
		fallthrough
	case "info":
		lvl = level.AllowInfo()
	case "error":
		lvl = level.AllowError()
	case "warn":
		lvl = level.AllowWarn()
	case "debug":
		lvl = level.AllowDebug()
	default:
		return nil, errors.Errorf("unexpected log level:%v", configLevel)
	}

	return level.NewFilter(logger, lvl), nil
}

func ExitOnError(logger log.Logger, err error) {
	if err != nil {
		if logger == nil {
			logstd.Fatal(err)
		}
		level.Error(logger).Log("err", err)
		os.Exit(1)
	}
}
