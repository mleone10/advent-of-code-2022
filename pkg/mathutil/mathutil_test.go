package mathutil_test

import (
	"testing"

	"github.com/mleone10/advent-of-code-2022/pkg/assert"
	"github.com/mleone10/advent-of-code-2022/pkg/mathutil"
)

func TestMax(t *testing.T) {
	assert.Equals(t, mathutil.Max(5, 10, 15), 15)
	assert.Equals(t, mathutil.Max(1.1, 2.2, 3.3), 3.3)
}

func TestMin(t *testing.T) {
	assert.Equals(t, mathutil.Min(5, 10, 15), 5)
	assert.Equals(t, mathutil.Min(1.1, 2.2, 3.3), 1.1)
}

func TestSum(t *testing.T) {
	assert.Equals(t, mathutil.Sum(5, 10, 15), 30)
	assert.Equals(t, mathutil.Sum(1.1, 2.2, 3.3), 6.6)
}

func TestFrequencyList(t *testing.T) {
	testIntFreqs := mathutil.FrequencyList(5, 5, 10, 10, 10, 15)
	assert.Equals(t, testIntFreqs[5], 2)
	assert.Equals(t, testIntFreqs[10], 3)
	assert.Equals(t, testIntFreqs[15], 1)

	testFloatFreqs := mathutil.FrequencyList(1.1, 1.1, 2.2, 2.2, 2.2, 3.3)
	assert.Equals(t, testFloatFreqs[1.1], 2)
	assert.Equals(t, testFloatFreqs[2.2], 3)
	assert.Equals(t, testFloatFreqs[3.3], 1)
}
