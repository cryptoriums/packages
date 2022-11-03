// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package constants

import "time"

const (
	TxGasOverHead      = 21_000
	BlockTime          = float64(12)
	BlocksPerSecond    = float64(1 / BlockTime)
	BlocksPerMinute    = float64(60 / BlockTime)
	ReorgEventWaitSafe = time.Minute
	ReorgEventWaitSlow = 3 * time.Minute
	ReorgEventWaitFast = 10 * time.Second

	MainnetName = "mainnet"
	RopstenName = "ropsten"
	GoerliName  = "goerli"
	RinkebyName = "rinkeby"
	HardhatName = "hardhat"

	MainnetID = 1
	RopstenID = 3
	GoerliID  = 4
	RinkebyID = 5
	HardhatID = 31337

	MaxBlockGasLimit = 30000000
	MaxGasPriceGwei  = 10_000 // To have some failsafe when creating TX and passing WEI instead of GWEI.
)

var NetworksByID = map[int64]string{
	MainnetID: MainnetName,
	RopstenID: RopstenName,
	RinkebyID: RinkebyName,
	GoerliID:  GoerliName,
	HardhatID: HardhatName,
}

var NetworksByName = map[string]int64{
	MainnetName: MainnetID,
	RopstenName: RopstenID,
	RinkebyName: RinkebyID,
	GoerliName:  GoerliID,
	HardhatName: HardhatID,
}
