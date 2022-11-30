package mathutil_test

import (
	"testing"

	"github.com/mleone10/advent-of-code-2022/pkg/assert"
	"github.com/mleone10/advent-of-code-2022/pkg/mathutil"
)

func TestMax(t *testing.T) {
	assert.Equal(t, mathutil.Max(5, 10, 15), 15)
	assert.Equal(t, mathutil.Max(1.1, 2.2, 3.3), 3.3)
}

func TestMin(t *testing.T) {
	assert.Equal(t, mathutil.Min(5, 10, 15), 5)
	assert.Equal(t, mathutil.Min(1.1, 2.2, 3.3), 1.1)
}

func TestSum(t *testing.T) {
	assert.Equal(t, mathutil.Sum(5, 10, 15), 30)
	assert.Equal(t, mathutil.Sum(1.1, 2.2, 3.3), 6.6)
	assert.Equal(t, mathutil.Sum([]int{5, 10, 15}...), 30)
}

func TestFrequencyList(t *testing.T) {
	testIntFreqs := mathutil.FrequencyList(5, 5, 10, 10, 10, 15)
	assert.Equal(t, testIntFreqs[5], 2)
	assert.Equal(t, testIntFreqs[10], 3)
	assert.Equal(t, testIntFreqs[15], 1)

	testFloatFreqs := mathutil.FrequencyList(1.1, 1.1, 2.2, 2.2, 2.2, 3.3)
	assert.Equal(t, testFloatFreqs[1.1], 2)
	assert.Equal(t, testFloatFreqs[2.2], 3)
	assert.Equal(t, testFloatFreqs[3.3], 1)
}

func TestSlidingSum(t *testing.T) {
	assert.ArraysEqual(t, mathutil.SlidingSum(2, []int{1, 2, 3, 4, 5, 6}), []int{3, 5, 7, 9, 11})
	assert.ArraysEqual(t, mathutil.SlidingSum(4, []int{1, 2, 3, 4, 5, 6}), []int{10, 14, 18})
	assert.ArraysEqual(t, mathutil.SlidingSum(3, []int{1, 2}), []int{})
}
