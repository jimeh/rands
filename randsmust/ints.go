package randsmust

import "github.com/jimeh/rands"

// Int generates a random int ranging between 0 and max.
func Int(max int) int {
	r, err := rands.Int(max)
	if err != nil {
		panic(err)
	}

	return r
}

// Int64 generates a random int64 ranging between 0 and max.
func Int64(max int64) int64 {
	r, err := rands.Int64(max)
	if err != nil {
		panic(err)
	}

	return r
}
