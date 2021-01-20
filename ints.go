package rands

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var errInvalidMaxInt = fmt.Errorf("%w: max cannot be less than 1", errBase)

// Int generates a random int ranging between 0 and max.
func Int(max int) (int, error) {
	if max < 1 {
		return 0, errInvalidMaxInt
	}

	r, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}

	return int(r.Int64()), nil
}

// Int64 generates a random int64 ranging between 0 and max.
func Int64(max int64) (int64, error) {
	if max < 1 {
		return 0, errInvalidMaxInt
	}

	r, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return 0, err
	}

	return r.Int64(), nil
}
