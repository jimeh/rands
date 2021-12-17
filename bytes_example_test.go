package rands_test

import (
	"fmt"
	"log"

	"github.com/jimeh/rands"
)

func ExampleBytes() {
	b, err := rands.Bytes(8)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", b) // => [181 153 143 235 241 20 208 173]
}
