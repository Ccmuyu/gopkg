package fn

func And[T any](predicates ...func(T) bool) func(T) bool {
	return func(v T) bool {
		for _, p := range predicates {
			if !p(v) {
				return false
			}
		}
		return true
	}
}

func Or[T any](predicates ...func(T) bool) func(T) bool {
	return func(v T) bool {
		for _, p := range predicates {
			if p(v) {
				return true
			}
		}
		return false
	}
}

func Not[T any](predicate func(T) bool) func(T) bool {
	return func(v T) bool {
		return !predicate(v)
	}
}

type Tuple2[A, B any] struct {
	A A
	B B
}

func MakeTuple2[A, B any](a A, b B) Tuple2[A, B] {
	return Tuple2[A, B]{A: a, B: b}
}

type Tuple3[A, B, C any] struct {
	A A
	B B
	C C
}

func MakeTuple3[A, B, C any](a A, b B, c C) Tuple3[A, B, C] {
	return Tuple3[A, B, C]{A: a, B: b, C: c}
}
