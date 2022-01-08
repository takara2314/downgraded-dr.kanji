package utils

import (
	"math/rand"
	"time"
)

// RandN returns a random value, 0 <= x < n.
func RandN(n int) int {
	// Update a seed
	rand.Seed(time.Now().Unix())

	return rand.Intn(n)
}

// RandChoiceString returns a value chosen its slice randomly.
func RandChoiceString(s []string) string {
	return s[RandN(len(s))]
}

// RandChoiceString2 returns a slice chosen its 2d-slice randomly.
func RandChoiceString2(s [][]string) []string {
	return s[RandN(len(s))]
}
