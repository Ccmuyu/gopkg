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
	if chunk <= 0 {
		return nil
	}
	if length == 0 {
		return nil
	}
	if length <= chunk {
		return [][]T{s[:]}
	}

	n := (length + chunk - 1) / chunk
	data := make([][]T, 0, n)
	for i := 0; i < length; i += chunk {
		next := i + chunk
		if next > length {
			data = append(data, s[i:])
			break
		}
		data = append(data, s[i:next])
	}

	return data
}

func Reverse[T any](arr []T) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func GroupBy[T any, K comparable](arr []T, fn func(T) K) map[K][]T {
	countMap := make(map[K]int)
	for _, v := range arr {
		countMap[fn(v)]++
	}
	result := make(map[K][]T, len(countMap))
	for k, count := range countMap {
		result[k] = make([]T, 0, count)
	}
	for _, v := range arr {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}

func Reduce[T any, R any](arr []T, initial R, fn func(R, T) R) R {
	result := initial
	for _, v := range arr {
		result = fn(result, v)
	}
	return result
}

func Some[T any](arr []T, fn func(T) bool) bool {
	for _, v := range arr {
		if fn(v) {
			return true
		}
	}
	return false
}

func Every[T any](arr []T, fn func(T) bool) bool {
	for _, v := range arr {
		if !fn(v) {
			return false
		}
	}
	return true
}

func None[T any](arr []T, fn func(T) bool) bool {
	return !Some(arr, fn)
}

func Without[T comparable](arr []T, items ...T) []T {
	if len(arr) == 0 {
		return arr
	}
	skip := make(map[T]struct{}, len(items))
	for _, item := range items {
		skip[item] = struct{}{}
	}
	result := make([]T, 0, len(arr))
	for _, v := range arr {
		if _, ok := skip[v]; !ok {
			result = append(result, v)
		}
	}
	return result
}

func Intersection[T comparable](a, b []T) []T {
	seen := make(map[T]struct{}, len(a))
	for _, v := range a {
		seen[v] = struct{}{}
	}
	result := make([]T, 0)
	added := make(map[T]struct{})
	for _, v := range b {
		if _, ok := seen[v]; ok {
			if _, addedOk := added[v]; !addedOk {
				result = append(result, v)
				added[v] = struct{}{}
			}
		}
	}
	return result
}

func Union[T comparable](slices ...[]T) []T {
	seen := make(map[T]struct{})
	total := 0
	for _, s := range slices {
		total += len(s)
	}
	result := make([]T, 0, total)
	for _, s := range slices {
		for _, v := range s {
			if _, ok := seen[v]; !ok {
				seen[v] = struct{}{}
				result = append(result, v)
			}
		}
	}
	return result
}

func Difference[T comparable](a, b []T) (inANotB, inBNotA []T) {
	setB := make(map[T]struct{}, len(b))
	for _, v := range b {
		setB[v] = struct{}{}
	}
	inANotB = make([]T, 0)
	for _, v := range a {
		if _, ok := setB[v]; !ok {
			inANotB = append(inANotB, v)
		}
	}
	setA := make(map[T]struct{}, len(a))
	for _, v := range a {
		setA[v] = struct{}{}
	}
	inBNotA = make([]T, 0)
	for _, v := range b {
		if _, ok := setA[v]; !ok {
			inBNotA = append(inBNotA, v)
		}
	}
	return
}

func Chunk[T any](s []T, size int) [][]T {
	return Split(s, size)
}