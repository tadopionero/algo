package main

import (
	"math"
)

// IsPowOfTwoUseLog This function checks if a number is a power of two using the logarithm.
// The limiting degree can be from 0 to 63.
// See alternatives in the binary package.
func IsPowOfTwoUseLog(number float64) bool {
	if number == 0 || math.Round(number) == math.MaxInt64 {
		return false
	}
	log := math.Log2(number)
	return log == math.Round(log)
}

func test_is_power_of_use_log() {
	println(IsPowOfTwoUseLog(5))
}
