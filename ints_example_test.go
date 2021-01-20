package rands_test

import (
	"fmt"

	"github.com/jimeh/rands"
)

func ExampleInt() {
	n, _ := rands.Int(2147483647)
	fmt.Printf("%d\n", n) // => 1334400235
}

func ExampleInt64() {
	n, _ := rands.Int64(int64(9223372036854775807))
	fmt.Printf("%d\n", n) // => 8256935979116161233
}
