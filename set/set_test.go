package set

import (
	"testing"

	. "github.com/Ccmuyu/gopkg/test"
)

func TestNew(t *testing.T) {
	s := New(1, 2, 3)
	AssertEqual(t, s.Len(), 3)
	AssertTrue(t, s.Contains(1))
	AssertTrue(t, s.Contains(2))
	AssertTrue(t, s.Contains(3))
	AssertTrue(t, !s.Contains(4))
}

func TestNewEmpty(t *testing.T) {
	s := New[int]()
	AssertEqual(t, s.Len(), 0)
}

func TestFromSlice(t *testing.T) {
	s := FromSlice([]int{1, 2, 3, 2, 1})
	AssertEqual(t, s.Len(), 3)
}

func TestAdd(t *testing.T) {
	s := New[int]()
	s.Add(1, 2, 3)
	AssertEqual(t, s.Len(), 3)
	s.Add(1)
	AssertEqual(t, s.Len(), 3)
}

func TestRemove(t *testing.T) {
	s := New(1, 2, 3)
	s.Remove(2)
	AssertEqual(t, s.Len(), 2)
	AssertTrue(t, !s.Contains(2))
	s.Remove(99)
	AssertEqual(t, s.Len(), 2)
}

func TestContains(t *testing.T) {
	s := New("a", "b")
	AssertTrue(t, s.Contains("a"))
	AssertTrue(t, s.Contains("b"))
	AssertTrue(t, !s.Contains("c"))
}

func TestToSlice(t *testing.T) {
	s := New(1, 2, 3)
	sl := s.ToSlice()
	AssertEqual(t, len(sl), 3)
}

func TestForEach(t *testing.T) {
	s := New(1, 2, 3)
	sum := 0
	s.ForEach(func(v int) {
		sum += v
	})
	AssertEqual(t, sum, 6)
}

func TestUnion(t *testing.T) {
	s1 := New(1, 2, 3)
	s2 := New(3, 4, 5)
	u := s1.Union(s2)
	AssertEqual(t, u.Len(), 5)
	AssertTrue(t, u.Contains(1))
	AssertTrue(t, u.Contains(5))
}

func TestIntersection(t *testing.T) {
	s1 := New(1, 2, 3, 4)
	s2 := New(3, 4, 5, 6)
	i := s1.Intersection(s2)
	AssertEqual(t, i.Len(), 2)
	AssertTrue(t, i.Contains(3))
	AssertTrue(t, i.Contains(4))
	AssertTrue(t, !i.Contains(1))
}

func TestIntersectionDisjoint(t *testing.T) {
	s1 := New(1, 2)
	s2 := New(3, 4)
	i := s1.Intersection(s2)
	AssertEqual(t, i.Len(), 0)
}

func TestDifference(t *testing.T) {
	s1 := New(1, 2, 3, 4)
	s2 := New(3, 4, 5)
	d := s1.Difference(s2)
	AssertEqual(t, d.Len(), 2)
	AssertTrue(t, d.Contains(1))
	AssertTrue(t, d.Contains(2))
	AssertTrue(t, !d.Contains(3))
}

func TestStringSet(t *testing.T) {
	s := New("foo", "bar", "baz")
	AssertEqual(t, s.Len(), 3)
	AssertTrue(t, s.Contains("foo"))
	s.Remove("bar")
	AssertEqual(t, s.Len(), 2)
	AssertTrue(t, !s.Contains("bar"))
}
