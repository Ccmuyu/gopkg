# slices

切片操作工具：过滤、映射、去重、分组、拆分等。

## 安装

```go
import "github.com/Ccmuyu/gopkg/slices"
```

## 快速开始

```go
nums := []int{1, 2, 3, 4, 5, 2, 3}

slices.Dedup(nums)                         // [1, 2, 3, 4, 5]

evens := slices.Filter(nums, func(v int) bool {
    return v%2 == 0
})                                         // [2, 4, 2]

doubled := slices.Map(nums, func(v int) int {
    return v * 2
})                                         // [2, 4, 6, 8, 10, 4, 6]

flat := slices.FlatMap([][]int{{1,2},{3,4}}, func(v []int) []int {
    return v
})                                         // [1, 2, 3, 4]

merged := slices.Merge([]int{1,2}, []int{3,4})  // [1, 2, 3, 4]

slices.Contains(nums, 3)                   // true
slices.Index(nums, 3)                      // 2

chunks := slices.Split(nums, 2)            // [[1,2],[3,4],[5,2],[3]]

grouped := slices.GroupBy(nums, func(v int) int {
    return v % 2
})                                         // {0: [2, 4, 2], 1: [1, 3, 5, 3]}
```

## API 参考

```go
func Dedup[T comparable](arr []T) []T
func Filter[T any](arr []T, fn func(T) bool) []T
func Map[T any, R any](arr []T, fn func(T) R) []R
func FlatMap[T any, R any](arr []T, fn func(T) []R) []R
func Merge[T any](slices ...[]T) []T
func Contains[T comparable](arr []T, target T) bool
func Index[T comparable](arr []T, target T) int
func Split[T any](s []T, chunk int) [][]T
func Reverse[T any](arr []T)
func GroupBy[T any, K comparable](arr []T, fn func(T) K) map[K][]T
func Reduce[T any, R any](arr []T, initial R, fn func(R, T) R) R
func Some[T any](arr []T, fn func(T) bool) bool
func Every[T any](arr []T, fn func(T) bool) bool
func None[T any](arr []T, fn func(T) bool) bool
func Without[T comparable](arr []T, items ...T) []T
func Intersection[T comparable](a, b []T) []T
func Union[T comparable](slices ...[]T) []T
func Difference[T comparable](a, b []T) (inANotB, inBNotA []T)
func Chunk[T any](s []T, size int) [][]T
func Sort[T cmp.Ordered](arr []T)
func Fill[T any](arr []T, val T)
func Take[T any](arr []T, n int) []T
func Drop[T any](arr []T, n int) []T
```
