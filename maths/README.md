# maths

常用数学计算工具：比大小、绝对值、求和、平均、幂运算、素数判断、最大公约数、最小公倍数、斐波那契数列等。

## 安装

```go
import "github.com/Ccmuyu/gopkg/maths"
```

## 快速开始

```go
maths.Min(3, 7)                    // 3
maths.Max(3, 7)                    // 7
maths.Clamp(15, 1, 10)             // 10
maths.Abs(-5)                      // 5
maths.Sum(1, 2, 3, 4, 5)           // 15
maths.Average(1, 2, 3, 4, 5)       // 3.0
maths.Pow(2, 10)                    // 1024
maths.IsPrime(17)                   // true
maths.GCD(12, 8)                    // 4
maths.LCM(4, 6)                     // 12
maths.Fibonacci(5)                  // [0, 1, 1, 2, 3]
```

## API 参考

```go
func Min[T cmp.Ordered](a, b T) T
func Max[T cmp.Ordered](a, b T) T
func Clamp[T cmp.Ordered](val, low, high T) T
func Abs[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64](x T) T
func Sum[T Numeric](nums ...T) T
func Average[T Numeric](nums ...T) float64
func Pow(base, exp int) int
func IsPrime(n int) bool
func GCD(a, b int) int
func LCM(a, b int) int
func Fibonacci(n int) []int
```
