package mathutil

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](args ...T) T {
	max := args[0]

	for _, arg := range args[1:] {
		if arg > max {
			max = arg
		}
	}

	return max
}

func Min[T constraints.Ordered](args ...T) T {
	min := args[0]

	for _, arg := range args[1:] {
		if arg < min {
			min = arg
		}
	}

	return min
}

func Sum[T constraints.Ordered](args ...T) T {
	sum := args[0]

	for _, arg := range args[1:] {
		sum += arg
	}

	return sum
}

func FrequencyList[T comparable](args ...T) map[T]int {
	freqs := map[T]int{}

	for _, arg := range args {
		freqs[arg] += 1
	}

	return freqs
}

func SlidingSum[T constraints.Ordered](width int, args []T) []T {
	if len(args) < width {
		return []T{}
	}

	var slidingSum []T

	for i := range args[:len(args)-width+1] {
		slidingSum = append(slidingSum, Sum(args[i:i+width]...))
	}

	return slidingSum
}
