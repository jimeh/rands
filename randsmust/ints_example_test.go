package randsmust_test

import (
	"fmt"

	"github.com/jimeh/rands/randsmust"
)

func ExampleInt() {
	n := randsmust.Int(2147483647)
	fmt.Printf("%d\n", n) // => 1616989970
}

func ExampleInt64() {
	n := randsmust.Int64(int64(9223372036854775807))
	fmt.Printf("%d\n", n) // => 1599573251306894157
}
