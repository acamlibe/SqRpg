package main

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
