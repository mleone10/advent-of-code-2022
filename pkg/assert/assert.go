package assert

import "testing"

func Equals[T comparable](t *testing.T, actual, expected T) {
	if actual != expected {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}
}
