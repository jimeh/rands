package randsmust

import "github.com/jimeh/rands"

// Shuffle randomizes the order of a collection of n elements using
// cryptographically secure random values from crypto/rand. It implements the
// Fisher-Yates shuffle algorithm.
//
// The swap function is called to exchange values at indices i and j. This
// signature is compatible with Shuffle from math/rand and math/rand/v2 for easy
// migration.
//
// If an error occurs during shuffling, this function will panic.
func Shuffle(n int, swap func(i, j int)) {
	err := rands.Shuffle(n, swap)
	if err != nil {
		panic(err)
	}
}

// ShuffleSlice randomizes the order of elements in a slice in-place using
// cryptographically secure random values from crypto/rand.
//
// It implements the Fisher-Yates shuffle algorithm.
//
// If an error occurs during shuffling, this function will panic.
func ShuffleSlice[T any](slice []T) {
	err := rands.ShuffleSlice(slice)
	if err != nil {
		panic(err)
	}
}
