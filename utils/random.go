package utils

import (
	"math/rand"
	"time"
)

// RandN returns a random value, 0 <= x < n.
func RandN(n int) int {
	// Update a seed
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(n)
}

// RandMN returns a random value, m <= x < n.
func RandMN(m int, n int) int {
	return RandN(n-m) + m
}

// RandChoiceString returns a value chosen its slice randomly.
func RandChoiceString(s []string) string {
	return s[RandN(len(s))]
}

// RandChoiceString2 returns a slice chosen its 2d-slice randomly.
func RandChoiceString2(s [][]string) []string {
	return s[RandN(len(s))]
}
