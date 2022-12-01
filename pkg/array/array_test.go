package array_test

import (
	"testing"

	"github.com/mleone10/advent-of-code-2022/pkg/array"
	"github.com/mleone10/advent-of-code-2022/pkg/assert"
)

func TestMax(t *testing.T) {
	assert.Equal(t, array.Max([]int{5, 10, 15}), 15)
	assert.Equal(t, array.Max([]float32{1.1, 2.2, 3.3}), 3.3)
}

func TestMin(t *testing.T) {
	assert.Equal(t, array.Min([]int{5, 10, 15}), 5)
	assert.Equal(t, array.Min([]float32{1.1, 2.2, 3.3}), 1.1)
}

func TestSum(t *testing.T) {
	assert.Equal(t, array.Sum([]int{5, 10, 15}), 30)
	assert.Equal(t, array.Sum([]float64{1.1, 2.2, 3.3}), 6.6)
}

func TestTake(t *testing.T) {
	assert.ArraysEqual(t, array.Take([]int{1, 2, 3, 4, 5}, 3), []int{1, 2, 3})
	assert.ArraysEqual(t, array.Take([]int{1, 2, 3}, 3), []int{1, 2, 3})
	assert.ArraysEqual(t, array.Take([]int{1, 2}, 3), []int{1, 2})
	assert.ArraysEqual(t, array.Take([]string{"foo", "bar", "fizz", "buzz"}, 3), []string{"foo", "bar", "fizz"})
}

func TestFrequencyList(t *testing.T) {
	testIntFreqs := array.FrequencyList([]int{5, 5, 10, 10, 10, 15})
	assert.Equal(t, testIntFreqs[5], 2)
	assert.Equal(t, testIntFreqs[10], 3)
	assert.Equal(t, testIntFreqs[15], 1)

	testFloatFreqs := array.FrequencyList([]float32{1.1, 1.1, 2.2, 2.2, 2.2, 3.3})
	assert.Equal(t, testFloatFreqs[1.1], 2)
	assert.Equal(t, testFloatFreqs[2.2], 3)
	assert.Equal(t, testFloatFreqs[3.3], 1)
}

func TestSlidingSum(t *testing.T) {
	assert.ArraysEqual(t, array.SlidingSum(2, []int{1, 2, 3, 4, 5, 6}), []int{3, 5, 7, 9, 11})
	assert.ArraysEqual(t, array.SlidingSum(4, []int{1, 2, 3, 4, 5, 6}), []int{10, 14, 18})
	assert.ArraysEqual(t, array.SlidingSum(3, []int{1, 2}), []int{})
}