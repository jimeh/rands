package randsmust

import "github.com/jimeh/rands"

// Int generates a random int ranging between 0 and nMax.
func Int(nMax int) int {
	r, err := rands.Int(nMax)
	if err != nil {
		panic(err)
	}

	return r
}

// Int64 generates a random int64 ranging between 0 and nMax.
func Int64(nMax int64) int64 {
	r, err := rands.Int64(nMax)
	if err != nil {
		panic(err)
	}

	return r
}
