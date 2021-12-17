package randsmust

import "github.com/jimeh/rands"

// Bytes generates a byte slice of n number of random bytes.
func Bytes(n int) []byte {
	r, err := rands.Bytes(n)
	if err != nil {
		panic(err)
	}

	return r
}
