package randsmust_test

import (
	"fmt"

	"github.com/jimeh/rands/randsmust"
)

func ExampleBytes() {
	b := randsmust.Bytes(8)
	fmt.Printf("%+v\n", b) // => [6 99 106 54 163 188 28 152]
}
