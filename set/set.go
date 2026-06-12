package set

type Set[T comparable] map[T]struct{}

func New[T comparable](items ...T) Set[T] {
	s := make(Set[T], len(items))
	for _, item := range items {
		s[item] = struct{}{}
	}
	return s
}

func FromSlice[T comparable](items []T) Set[T] {
	return New(items...)
}

func (s Set[T]) Add(items ...T) {
	for _, item := range items {
		s[item] = struct{}{}
	}
}

func (s Set[T]) Remove(item T) {
	delete(s, item)
}

func (s Set[T]) Contains(item T) bool {
	_, ok := s[item]
	return ok
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) ToSlice() []T {
	result := make([]T, 0, len(s))
	for item := range s {
		result = append(result, item)
	}
	return result
}

func (s Set[T]) ForEach(fn func(T)) {
	for item := range s {
		fn(item)
	}
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	result := make(Set[T], len(s)+len(other))
	for item := range s {
		result[item] = struct{}{}
	}
	for item := range other {
		result[item] = struct{}{}
	}
	return result
}

func (s Set[T]) Intersection(other Set[T]) Set[T] {
	if len(s) > len(other) {
		return other.Intersection(s)
	}
	result := make(Set[T])
	for item := range s {
		if other.Contains(item) {
			result[item] = struct{}{}
		}
	}
	return result
}

func (s Set[T]) Difference(other Set[T]) Set[T] {
	result := make(Set[T])
	for item := range s {
		if !other.Contains(item) {
			result[item] = struct{}{}
		}
	}
	return result
}
