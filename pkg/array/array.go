// Array contains utility functions to simplify common operations on slices and arrays.
package array

import "golang.org/x/exp/constraints"

// Max returns the largest value among a slice of Ordered values via the `>` operator.
func Max[T constraints.Ordered](args []T) T {
	max := args[0]

	for _, arg := range args[1:] {
		if arg > max {
			max = arg
		}
	}

	return max
}

// Min returns the smallest value among a slice of Ordered values via the `<` operator.
func Min[T constraints.Ordered](args []T) T {
	min := args[0]

	for _, arg := range args[1:] {
		if arg < min {
			min = arg
		}
	}

	return min
}

// Sum returns the aggregate total of a slice of Ordered values via the `+=` operator.
func Sum[T constraints.Ordered](args []T) T {
	sum := args[0]

	for _, arg := range args[1:] {
		sum += arg
	}

	return sum
}

// Take returns the first n elements of the given slice.
func Take[T any](arr []T, n int) []T {
	if len(arr) <= n {
		return arr
	}

	return arr[:n]
}

// FrequencyList accepts a slice of comparable-type values and returns a map[T]int representing how many times each key appears in the input slice.
func FrequencyList[T comparable](args []T) map[T]int {
	freqs := map[T]int{}

	for _, arg := range args {
		freqs[arg] += 1
	}

	return freqs
}

// SlidingSum iterates over a given slice and sums sub-slices of w-width elements and returns an ordered slice of those sums.
func SlidingSum[T constraints.Ordered](w int, args []T) []T {
	if len(args) < w {
		return []T{}
	}

	var slidingSum []T

	for i := range args[:len(args)-w+1] {
		slidingSum = append(slidingSum, Sum(args[i:i+w]))
	}

	return slidingSum
}
