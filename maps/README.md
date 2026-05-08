# maps

Map 常用工具函数，基于 Go 1.18+ 泛型实现，类型安全、无第三方依赖。

## 安装

```go
import "github.com/Ccmuyu/gopkg/maps"
```

## 快速开始

```go
m := map[string]int{"a": 1, "b": 2, "c": 3}

keys := maps.Keys(m)       // ["a", "b", "c"]
vals := maps.Values(m)     // [1, 2, 3]
maps.HasKey(m, "a")        // true

picked := maps.Pick(m, "a", "c")      // {"a": 1, "c": 3}
omitted := maps.Omit(m, "b")          // {"a": 1, "c": 3}
inverted := maps.Invert(m)            // {1: "a", 2: "b", 3: "c"}

filtered := maps.Filter(m, func(k string, v int) bool {
    return v > 1
})                                     // {"b": 2, "c": 3}

mapped := maps.Map(m, func(k string, v int) int {
    return v * 2
})                                     // {"a": 2, "b": 4, "c": 6}

keyMapped := maps.MapKey(m, func(k string) string {
    return "pre_" + k
})                                     // {"pre_a": 1, "pre_b": 2, "pre_c": 3}

merged := maps.Merge(m, map[string]int{"d": 4})
// {"a": 1, "b": 2, "c": 3, "d": 4}

mergedWith := maps.MergeWith(func(v1, v2 int) int {
    return v1 + v2
}, map[string]int{"a": 1, "b": 2}, map[string]int{"b": 3, "c": 4})
// {"a": 1, "b": 5, "c": 4}

maps.IsEqual(m, map[string]int{"a": 1, "b": 2, "c": 3})  // true
```

## API 参考

### 查询

```go
func Keys[K comparable, V any](m map[K]V) []K
func Values[K comparable, V any](m map[K]V) []V
func HasKey[K comparable, V any](m map[K]V, key K) bool
func Get[K comparable, V any](m map[K]V, key K) V
func GetOk[K comparable, V any](m map[K]V, key K) (V, bool)
func IsEqual[K, V comparable](m1, m2 map[K]V) bool
```

### 变换

```go
func Map[K comparable, V any, R any](m map[K]V, fn func(K, V) R) map[K]R
func MapKey[K comparable, V any, RK comparable](m map[K]V, fn func(K) RK) map[RK]V
func Filter[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V
```

### 合并

```go
func Merge[K comparable, V any](maps ...map[K]V) map[K]V
func MergeWith[K comparable, V any](mergeFn func(V, V) V, maps ...map[K]V) map[K]V
```

### 子集操作

```go
func Pick[K comparable, V any](m map[K]V, keys ...K) map[K]V
func Omit[K comparable, V any](m map[K]V, keys ...K) map[K]V
```

### 其他

```go
func Set[K comparable, V any](m map[K]V, key K, value V)
func Clone[K comparable, V any](m map[K]V) map[K]V
func Invert[K comparable, V comparable](m map[K]V) map[V]K
func ForEach[K comparable, V any](m map[K]V, fn func(K, V))
```

### MergeWith 详解

`Merge` 在遇到相同 key 时直接覆盖（后者的值胜出），而 `MergeWith` 允许自定义合并策略：

```go
// 取最大值
maps.MergeWith(func(v1, v2 int) int {
    if v1 > v2 { return v1 }; return v2
}, m1, m2)

// 字符串拼接
maps.MergeWith(func(v1, v2 string) string {
    return v1 + v2
}, m1, m2)
```

### Pick / Omit 详解

`Pick` 选取指定 key，`Omit` 排除指定 key：

```go
m := map[string]int{"a": 1, "b": 2, "c": 3}

maps.Pick(m, "a", "c")       // {"a": 1, "c": 3}
maps.Omit(m, "b")            // {"a": 1, "c": 3}
```

### Invert 详解

键值互换，值必须为 comparable：

```go
m := map[string]int{"a": 1, "b": 2}
maps.Invert(m)               // {1: "a", 2: "b"}
```

## 测试

```bash
go test ./maps/ -v
```
