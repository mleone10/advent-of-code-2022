package assert_test

import (
	"testing"

	"github.com/mleone10/advent-of-code-2022/pkg/assert"
)

func TestEquals_MatchingValues_Passes(t *testing.T) {
	testT := &testing.T{}

	assert.Equals(testT, 10, 10)
	assert.Equals(testT, "hello", "hello")
	assert.Equals(testT, true, true)

	if testT.Failed() {
		t.Errorf("unexpected failure when comparing known equal values")
	}
}

func TestEquals_DifferentValues_Fails(t *testing.T) {
	testT := &testing.T{}

	assert.Equals(testT, 10, 12)
	assert.Equals(testT, "hello", "world")
	assert.Equals(testT, true, false)

	if !testT.Failed() {
		t.Errorf("unexpected success when comparing known different values")
	}
}
