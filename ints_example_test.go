package rands_test

import (
	"fmt"
	"log"

	"github.com/jimeh/rands"
)

func ExampleInt() {
	n, err := rands.Int(2147483647)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", n) // => 1908357440
}

func ExampleInt64() {
	n, err := rands.Int64(int64(9223372036854775807))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", n) // => 6530460062499341591
}
