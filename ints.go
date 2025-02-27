package rands

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var ErrInvalidMaxInt = fmt.Errorf("%w: max cannot be less than 1", Err)

// Int generates a random int ranging between 0 and nMax.
func Int(nMax int) (int, error) {
	if nMax < 1 {
		return 0, ErrInvalidMaxInt
	}

	r, err := rand.Int(rand.Reader, big.NewInt(int64(nMax)))
	if err != nil {
		return 0, err
	}

	return int(r.Int64()), nil
}

// Int64 generates a random int64 ranging between 0 and nMax.
func Int64(nMax int64) (int64, error) {
	if nMax < 1 {
		return 0, ErrInvalidMaxInt
	}

	r, err := rand.Int(rand.Reader, big.NewInt(nMax))
	if err != nil {
		return 0, err
	}

	return r.Int64(), nil
}
