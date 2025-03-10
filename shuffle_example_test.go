package rands_test

import (
	"fmt"
	"log"

	"github.com/jimeh/rands"
)

func ExampleShuffle() {
	numbers := []int{1, 2, 3, 4, 5}

	err := rands.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(numbers) // => [2 4 5 1 3]
}

func ExampleShuffleSlice() {
	mixed := []any{1, "two", 3.14, true, nil}

	err := rands.ShuffleSlice(mixed)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(mixed) // => [3.14 true 1 two <nil>]
}
