package slices

func Dedup[T comparable](arr []T) []T {
	if len(arr) == 0 {
		return arr
	}
	seen := make(map[T]struct{}, len(arr))
	result := make([]T, 0, len(arr))
	for _, v := range arr {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

func Filter[T any](arr []T, fn func(T) bool) []T {
	result := make([]T, 0, len(arr))
	for _, v := range arr {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func Map[T any, R any](arr []T, fn func(T) R) []R {
	result := make([]R, len(arr))
	for i, v := range arr {
		result[i] = fn(v)
	}
	return result
}

func FlatMap[T any, R any](arr []T, fn func(T) []R) []R {
	result := make([]R, 0)
	for _, v := range arr {
		result = append(result, fn(v)...)
	}
	return result
}

func Merge[T any](slices ...[]T) []T {
	total := 0
	for _, s := range slices {
		total += len(s)
	}
	result := make([]T, 0, total)
	for _, s := range slices {
		result = append(result, s...)
	}
	return result
}

func Contains[T comparable](arr []T, target T) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

func Index[T comparable](arr []T, target T) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

func Split[T any](s []T, chunk int) [][]T {
	length := len(s)
	if length == 0 || length <= chunk {
		return [][]T{s[:]}
	}

	var data [][]T
	for i := 0; i < length; i += chunk {
		next := i + chunk
		if next > length {
			data = append(data, s[i:])
			break
		}
		data = append(data, s[i:i+chunk])
	}

	return data
}

func Reverse[T any](arr []T) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func GroupBy[T any, K comparable](arr []T, fn func(T) K) map[K][]T {
	result := make(map[K][]T)
	for _, v := range arr {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}