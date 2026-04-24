package maps

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func HasKey[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}

func Get[K comparable, V any](m map[K]V, key K) V {
	return m[key]
}

func GetOk[K comparable, V any](m map[K]V, key K) (V, bool) {
	v, ok := m[key]
	return v, ok
}

func Set[K comparable, V any](m map[K]V, key K, value V) {
	m[key] = value
}

func Filter[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		if fn(k, v) {
			result[k] = v
		}
	}
	return result
}

func Map[K comparable, V any, R any](m map[K]V, fn func(K, V) R) map[K]R {
	result := make(map[K]R, len(m))
	for k, v := range m {
		result[k] = fn(k, v)
	}
	return result
}

func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	total := 0
	for _, m := range maps {
		total += len(m)
	}
	result := make(map[K]V, total)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

func Clone[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}