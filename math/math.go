// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package math

import (
	"math"
)

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

// ConfidenceInDifference calculates the percentage difference between the max and min and subtract this from 100%.
// Example:
// min 1, max 2
// Difference is 1 which is 100% so the final confidence is 100-100 equals 0%.
func ConfidenceInDifference(min, max float64) float64 {
	return 100 - (math.Abs(min-max)/min)*100
}
