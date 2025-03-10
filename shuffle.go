package rands

import (
	"fmt"
)

var (
	ErrShuffle                 = fmt.Errorf("%w: shuffle", Err)
	ErrInvalidShuffleNegativeN = fmt.Errorf(
		"%w: n must not be negative", ErrShuffle,
	)
)

// Shuffle randomizes the order of a collection of n elements using
// cryptographically secure random values from crypto/rand. It implements the
// Fisher-Yates shuffle algorithm.
//
// The swap function is called to exchange values at indices i and j. This
// signature is compatible with Shuffle from math/rand and math/rand/v2 for easy
// migration.
func Shuffle(n int, swap func(i, j int)) error {
	if n < 0 {
		return ErrInvalidShuffleNegativeN
	}

	for i := n - 1; i > 0; i-- {
		j, err := Int(i + 1)
		if err != nil {
			return err
		}

		swap(i, j)
	}

	return nil
}

// ShuffleSlice randomizes the order of elements in a slice in-place using
// cryptographically secure random values from crypto/rand.
//
// It implements the Fisher-Yates shuffle algorithm.
func ShuffleSlice[T any](slice []T) error {
	// If the slice has one or no elements, there's nothing to shuffle.
	if len(slice) < 2 {
		return nil
	}

	return Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}
