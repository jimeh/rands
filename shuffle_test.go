package rands

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func factorial(n int) int {
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}

	return factorial
}

func TestShuffle(t *testing.T) {
	t.Parallel()

	t.Run("n < 0", func(t *testing.T) {
		t.Parallel()

		err := Shuffle(-1, func(_, _ int) {})

		require.ErrorIs(t, err, ErrInvalidShuffleNegativeN)
		require.ErrorIs(t, err, ErrShuffle)
		require.ErrorIs(t, err, Err)
	})

	t.Run("n == 0", func(t *testing.T) {
		t.Parallel()

		swapCount := 0
		err := Shuffle(0, func(_, _ int) {
			swapCount++
		})
		require.NoError(t, err)

		assert.Equal(t, 0, swapCount)
	})

	t.Run("n == 1", func(t *testing.T) {
		t.Parallel()

		swapCount := 0
		err := Shuffle(1, func(_, _ int) {
			swapCount++
		})
		require.NoError(t, err)

		assert.Equal(t, 0, swapCount)
	})

	t.Run("n == 2", func(t *testing.T) {
		t.Parallel()

		swapCount := 0
		err := Shuffle(2, func(_, _ int) {
			swapCount++
		})
		require.NoError(t, err)

		assert.Equal(t, 1, swapCount)
	})

	t.Run("basic", func(t *testing.T) {
		t.Parallel()

		arr := make([]int, 100)
		for i := range arr {
			arr[i] = i
		}

		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)

		err := Shuffle(len(arr), func(i, j int) {
			arr[i], arr[j] = arr[j], arr[i]
		})
		require.NoError(t, err, "Shuffle returned an error")

		assert.NotEqual(t, arrCopy, arr, "Shuffle did not change the array")
		assert.ElementsMatch(t, arrCopy, arr, "Shuffle changed elements")
	})

	t.Run("swaps", func(t *testing.T) {
		t.Parallel()

		swapSame := 0
		swapDifferent := 0
		arr := make([]int, 100)
		for j := range arr {
			arr[j] = j
		}

		err := Shuffle(len(arr), func(i, j int) {
			if i == j {
				swapSame++
			} else {
				swapDifferent++
			}

			arr[i], arr[j] = arr[j], arr[i]
		})
		require.NoError(t, err, "Shuffle returned an error")

		// Fisher-Yates with n elements should make exactly n-1 swaps
		assert.Equal(t, len(arr)-1, swapSame+swapDifferent,
			"Unexpected swaps count",
		)

		// Ensure we have more different-element swaps than self-swaps. The
		// lower the input shuffle n value, the more likely this assertion will
		// fail. For a n=100 shuffle, this is exceptionally unlikely to fail.
		assert.Greater(t, swapDifferent, swapSame,
			"Expected more different-element swaps than self-swaps",
		)
	})

	t.Run("swap ranges", func(t *testing.T) {
		t.Parallel()

		n := 32
		runs := 1000

		for run := 0; run < runs; run++ {
			called := 0
			err := Shuffle(n, func(i, j int) {
				called++

				// Verify indices are in bounds.
				assert.True(t, i >= 0 && i < n, "Out of bounds index i = %d", i)
				assert.True(t, j >= 0 && j < n, "Out of bounds index j = %d", j)

				// For Fisher-Yates, i should be > 0 and j should be in range
				// [0,i].
				assert.Greater(t, i, 0, "Expected i > 0, got i=%d", i)
				assert.True(t,
					j >= 0 && j <= i,
					"Expected j in range [0,%d], got j=%d", i, j,
				)
			})
			require.NoError(t, err, "Shuffle returned an error")

			// Fisher-Yates with n elements should make exactly n-1 swaps
			expected := n - 1
			assert.Equal(t, expected, called,
				"Expected %d swap calls, got %d", expected, called,
			)
		}
	})

	t.Run("all permutations", func(t *testing.T) {
		t.Parallel()

		// Use a small array of 5 elements to make it feasible to track all
		// permutations.
		n := 5
		fact := factorial(n) // 120
		runs := fact * 3000  // 360000

		permCounts := make(map[string]int)
		for i := 0; i < runs; i++ {
			arr := make([]int, n)
			for i := range arr {
				arr[i] = i
			}

			err := Shuffle(len(arr), func(i, j int) {
				arr[i], arr[j] = arr[j], arr[i]
			})
			require.NoError(t, err, "Shuffle returned an error")

			// Convert the permutation to a string key and count it.
			key := fmt.Sprintf("%v", arr)
			permCounts[key]++
		}

		assert.Equal(t, fact, len(permCounts),
			"Expected %d different permutations", fact,
		)

		wantCount := float64(runs) / float64(fact)
		margin := 0.15
		minAcceptable := int(wantCount * (1 - margin))
		maxAcceptable := int(wantCount * (1 + margin))

		for perm, count := range permCounts {
			assert.True(t,
				count >= minAcceptable && count <= maxAcceptable,
				"Non-uniform distribution for %s: count=%d, expected=%v±%v",
				perm, count, wantCount, wantCount*margin,
			)
		}
	})

	t.Run("distribution", func(t *testing.T) {
		t.Parallel()
		// Track which positions received which random indices
		n := 100
		posCounts := make([]map[int]int, n)
		for i := range posCounts {
			posCounts[i] = make(map[int]int)
		}

		runs := 3000
		for run := 0; run < runs; run++ {
			err := Shuffle(n, func(i, j int) {
				posCounts[i][j]++
			})
			require.NoError(t, err, "Shuffle returned an error")
		}

		// For each position, check that it received a reasonable distribution.
		for i := n - 1; i >= n-len(posCounts); i-- {
			// Calculate how many unique positions we should expect.
			// Position i should receive random positions from 0 to i, and
			// allow for some statistical variation.
			want := int(float64(i+1) * 0.9)
			assert.GreaterOrEqual(t,
				len(posCounts[i]), want,
				"Position %d: expected ~%d unique indices, got %d",
				i, want, len(posCounts[i]),
			)
		}
	})
}

func BenchmarkShuffle(b *testing.B) {
	ranges := []int{32, 64, 128, 1024, 4096}
	for _, n := range ranges {
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Shuffle(n, func(_, _ int) {})
			}
		})
	}
}

func TestShuffleSlice(t *testing.T) {
	t.Parallel()

	t.Run("empty slice", func(t *testing.T) {
		t.Parallel()

		slice := []int{}
		err := ShuffleSlice(slice)
		require.NoError(t, err)
		assert.Empty(t, slice)
	})

	t.Run("single element", func(t *testing.T) {
		t.Parallel()

		slice := []int{42}
		origSlice := make([]int, len(slice))
		copy(origSlice, slice)

		err := ShuffleSlice(slice)
		require.NoError(t, err)

		assert.Equal(t,
			origSlice, slice, "Single element slice should remain unchanged",
		)
	})

	t.Run("two elements", func(t *testing.T) {
		t.Parallel()

		slice := []int{1, 2}
		origSlice := make([]int, len(slice))
		copy(origSlice, slice)

		err := ShuffleSlice(slice)
		require.NoError(t, err)

		// With two elements, the slice might remain the same or be swapped
		assert.Len(t, slice, len(origSlice))
		assert.ElementsMatch(t, origSlice, slice)
	})

	t.Run("basic", func(t *testing.T) {
		t.Parallel()

		slice := make([]int, 100)
		for i := range slice {
			slice[i] = i
		}

		sliceCopy := make([]int, len(slice))
		copy(sliceCopy, slice)

		err := ShuffleSlice(slice)
		require.NoError(t, err, "ShuffleSlice returned an error")

		assert.NotEqual(t,
			sliceCopy, slice, "ShuffleSlice did not change the slice",
		)
		assert.ElementsMatch(t,
			sliceCopy, slice, "ShuffleSlice changed elements",
		)
	})

	t.Run("string slice", func(t *testing.T) {
		t.Parallel()

		strSlice := []string{"a", "b", "c", "d", "e"}
		strCopy := make([]string, len(strSlice))
		copy(strCopy, strSlice)

		err := ShuffleSlice(strSlice)
		require.NoError(t, err)

		assert.ElementsMatch(t, strCopy, strSlice)
	})

	t.Run("struct slice", func(t *testing.T) {
		t.Parallel()

		type testStruct struct {
			id   int
			name string
		}
		structSlice := []testStruct{
			{1, "one"},
			{2, "two"},
			{3, "three"},
			{4, "four"},
		}
		structCopy := make([]testStruct, len(structSlice))
		copy(structCopy, structSlice)

		err := ShuffleSlice(structSlice)
		require.NoError(t, err)
		assert.ElementsMatch(t, structCopy, structSlice)
	})

	t.Run("all permutations", func(t *testing.T) {
		t.Parallel()

		// Use a small slice of 5 elements to make it feasible to track all
		// permutations.
		n := 5
		fact := factorial(n) // 120
		runs := fact * 3000  // 360000

		permCounts := make(map[string]int)
		for i := 0; i < runs; i++ {
			slice := make([]int, n)
			for j := range slice {
				slice[j] = j
			}

			err := ShuffleSlice(slice)
			require.NoError(t, err, "ShuffleSlice returned an error")

			// Convert the permutation to a string key and count it.
			key := fmt.Sprintf("%v", slice)
			permCounts[key]++
		}

		assert.Equal(t, fact, len(permCounts),
			"Expected %d different permutations", fact,
		)

		wantCount := float64(runs) / float64(fact)
		margin := 0.15
		minAcceptable := int(wantCount * (1 - margin))
		maxAcceptable := int(wantCount * (1 + margin))

		for perm, count := range permCounts {
			assert.True(t,
				count >= minAcceptable && count <= maxAcceptable,
				"Non-uniform distribution for %s: count=%d, expected=%v±%v",
				perm, count, wantCount, wantCount*margin,
			)
		}
	})

	t.Run("distribution", func(t *testing.T) {
		t.Parallel()

		// Track where each original index ends up after shuffling
		n := 100
		// posCounts[originalPos][newPos] tracks how many times
		// the element originally at position i ended up at position j
		posCounts := make([]map[int]int, n)
		for i := range posCounts {
			posCounts[i] = make(map[int]int)
		}

		runs := 3000
		for run := 0; run < runs; run++ {
			// Create a slice where the value is its original position
			slice := make([]int, n)
			for i := range slice {
				slice[i] = i
			}

			err := ShuffleSlice(slice)
			require.NoError(t, err, "ShuffleSlice returned an error")

			// Track where each original position ended up
			for newPos, origPos := range slice {
				posCounts[origPos][newPos]++
			}
		}

		// For each original position, check that it was distributed
		// reasonably across all possible new positions
		for i := n - 1; i >= n-len(posCounts); i-- {
			// Calculate how many unique positions we should expect.
			// Position i should receive random positions from 0 to i, and
			// allow for some statistical variation.
			want := int(float64(i+1) * 0.9)
			assert.GreaterOrEqual(t,
				len(posCounts[i]), want,
				"Original position %d: expected ~%d unique positions, got %d",
				i, want, len(posCounts[i]),
			)
		}
	})
}

func BenchmarkShuffleSlice(b *testing.B) {
	ranges := []int{32, 64, 128, 1024, 4096}
	for _, n := range ranges {
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			b.StopTimer()
			slice := make([]int, n)
			for i := range slice {
				slice[i] = i
			}
			b.StartTimer()

			for i := 0; i < b.N; i++ {
				_ = ShuffleSlice(slice)
			}
		})
	}
}
