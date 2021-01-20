package rands_test

import (
	"fmt"

	"github.com/jimeh/rands"
)

func ExampleBytes() {
	b, _ := rands.Bytes(8)
	fmt.Printf("%+v\n", b) // => [0 220 137 243 135 204 34 63]
}
