# set

泛型集合实现，提供 Set 数据结构的常用操作。

## 安装

```go
import "github.com/Ccmuyu/gopkg/set"
```

## 快速开始

```go
s := set.New(1, 2, 3, 2, 1)
s.Contains(1)                     // true
s.Len()                           // 3
s.Add(4, 5)
s.Remove(1)
s.ToSlice()                       // [2, 3, 4, 5]

s2 := set.New(3, 4, 5)
s.Union(s2)                       // {2, 3, 4, 5}
s.Intersection(s2)                // {3, 4}
s.Difference(s2)                  // {2, 5}
```

## API 参考

```go
func New[T comparable](items ...T) Set[T]
func FromSlice[T comparable](items []T) Set[T]
func (s Set[T]) Add(items ...T)
func (s Set[T]) Remove(item T)
func (s Set[T]) Contains(item T) bool
func (s Set[T]) Len() int
func (s Set[T]) ToSlice() []T
func (s Set[T]) ForEach(fn func(T))
func (s Set[T]) Union(other Set[T]) Set[T]
func (s Set[T]) Intersection(other Set[T]) Set[T]
func (s Set[T]) Difference(other Set[T]) Set[T]
```
