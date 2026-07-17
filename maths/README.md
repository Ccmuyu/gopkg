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
maths.Round(3.5)                    // 4.0
maths.Floor(3.7)                    // 3.0
maths.Ceil(3.2)                     // 4.0
maths.Factorial(5)                  // 120
maths.Median(1, 2, 3, 4, 5)        // 3.0
```

`Abs`、`GCD` 和 `LCM` 在数学结果超出返回类型可表示范围时采用最大值饱和语义。

## API 参考

```go
func Min[T cmp.Ordered](a, b T) T
func Max[T cmp.Ordered](a, b T) T
func Clamp[T cmp.Ordered](val, low, high T) T
func Abs[T Numeric](x T) T
func Sum[T Numeric](nums ...T) T
func Average[T Numeric](nums ...T) float64
func Pow(base, exp int) int
func IsPrime(n int) bool
func GCD(a, b int) int
func LCM(a, b int) int
func Fibonacci(n int) []int
func Round(f float64) float64
func Floor(f float64) float64
func Ceil(f float64) float64
func MinSlice[T cmp.Ordered](arr []T) T
func MaxSlice[T cmp.Ordered](arr []T) T
func Factorial(n int) int
func Median[T Numeric](nums ...T) float64
```
