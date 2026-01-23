package utils

import (
	"math/rand"
)

/*
RandBool

	This function returns a random boolean value based on the current time
*/
func RandBool() bool {
	return rand.Intn(2) == 1
}

// RandChance returns true with probability 1/denominator
func RandChance(denominator int) bool {
	if denominator <= 0 {
		return false // invalid input
	}
	return rand.Intn(denominator) == 0
}
