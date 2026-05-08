# sorts

排序工具：支持自定义比较器和稳定排序。

## 安装

```go
import "github.com/Ccmuyu/gopkg/sorts"
```

## 快速开始

```go
nums := []int{3, 1, 4, 1, 5, 9, 2, 6}

sorts.Sort(nums, true)                     // 升序 [1, 1, 2, 3, 4, 5, 6, 9]
sorts.Sort(nums, false)                    // 降序 [9, 6, 5, 4, 3, 2, 1, 1]

sorts.SortFunc(nums, func(a, b int) int {
    return a - b
})                                         // 自定义排序

sorts.SortFuncStable(users, func(a, b User) int {
    return a.Age - b.Age
})                                         // 稳定排序
```

## API 参考

```go
func Sort[T Comparable](arr []T, asc bool)
func SortFunc[T any](arr []T, cmp func(a, b T) int)
func SortFuncStable[T any](arr []T, cmp func(a, b T) int)
```
