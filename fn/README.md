# fn

函数式编程辅助工具：谓词组合、元组。

## 安装

```go
import "github.com/Ccmuyu/gopkg/fn"
```

## 快速开始

```go
isPositive := func(v int) bool { return v > 0 }
isEven := func(v int) bool { return v%2 == 0 }

positiveAndEven := fn.And(isPositive, isEven)
positiveAndEven(2)   // true
positiveAndEven(1)   // false

tup := fn.MakeTuple2("hello", 42)
tup.A  // "hello"
tup.B  // 42
```

## API 参考

```go
func And[T any](predicates ...func(T) bool) func(T) bool
func Or[T any](predicates ...func(T) bool) func(T) bool
func Not[T any](predicate func(T) bool) func(T) bool
func MakeTuple2[A, B any](a A, b B) Tuple2[A, B]
func MakeTuple3[A, B, C any](a A, b B, c C) Tuple3[A, B, C]

type Tuple2[A, B any] struct { A A; B B }
type Tuple3[A, B, C any] struct { A A; B B; C C }
```
