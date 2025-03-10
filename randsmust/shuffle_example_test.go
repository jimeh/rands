package randsmust_test

import (
	"fmt"

	"github.com/jimeh/rands/randsmust"
)

func ExampleShuffle() {
	numbers := []int{1, 2, 3, 4, 5}

	randsmust.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})

	fmt.Println(numbers) // => [4 2 5 3 1]
}

func ExampleShuffleSlice() {
	mixed := []any{1, "two", 3.14, true, nil}

	randsmust.ShuffleSlice(mixed)

	fmt.Println(mixed) // => [two <nil> true 3.14 1]
}
