package assert

import "testing"

// TODO: implement ContainsKey, ContainsValue assertions

func Equal[T comparable](t *testing.T, actual, expected T) {
	if actual != expected {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}
}

func NotEqual[T comparable](t *testing.T, actual, expected T) {
	if actual == expected {
		t.Errorf("values unexpectedly equal %+v", actual)
	}
}

func ArraysEqual[T comparable](t *testing.T, actual, expected []T) {
	if len(actual) != len(expected) {
		t.Errorf("arrays are not equal length: wanted len=%v %v, got len=%v %v", len(expected), expected, len(actual), actual)
		return
	}

	for i, v := range actual {
		if expected[i] != v {
			t.Errorf("elements at index %d do not match: wanted %v, got %v", i, expected[i], v)
		}
	}
}

func Contains[T comparable](t *testing.T, arr []T, val T) {
	for _, v := range arr {
		if v == val {
			return
		}
	}

	t.Errorf("value %+v not found in slice", val)
}

func DoesNotContain[T comparable](t *testing.T, arr []T, val T) {
	for _, v := range arr {
		if v == val {
			t.Errorf("value %+v found in slice", val)
		}
	}
}

func IsNil(t *testing.T, val interface{}) {
	if val != nil {
		t.Errorf("value %+v was not nil", val)
	}
}
