package utils

import (
	"math/rand"

	rand2 "github.com/chutommy/rand"
)

// Rand returns a math/rand Rand instance with a cryptographically
// generated seed for true random generating.
func Rand() *rand.Rand {
	return rand2.NewRand()
}
