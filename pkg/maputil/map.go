package maputil

func Keys[K comparable, V any](m map[K]V) []K {
	var keys []K

	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	var vals []V

	for _, v := range m {
		vals = append(vals, v)
	}

	return vals
}
