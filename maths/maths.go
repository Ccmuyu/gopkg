package maths

import "cmp"

func Min[T cmp.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Clamp[T cmp.Ordered](val, low, high T) T {
	if val < low {
		return low
	}
	if val > high {
		return high
	}
	return val
}

func Abs[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func Sum[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64](nums ...T) T {
	var sum T
	for _, n := range nums {
		sum += n
	}
	return sum
}

func Average[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64](nums ...T) float64 {
	if len(nums) == 0 {
		return 0
	}
	var sum T
	for _, n := range nums {
		sum += n
	}
	return float64(sum) / float64(len(nums))
}

func Pow(base, exp int) int {
	if exp < 0 {
		return 0
	}
	result := 1
	for exp > 0 {
		if exp&1 == 1 {
			result *= base
		}
		base *= base
		exp >>= 1
	}
	return result
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return a / GCD(a, b) * b
}

func Fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}
	result := make([]int, n)
	result[0], result[1] = 0, 1
	for i := 2; i < n; i++ {
		result[i] = result[i-1] + result[i-2]
	}
	return result
}

func Round(f float64) float64 {
	if f < 0 {
		return -Round(-f)
	}
	return float64(int(f + 0.5))
}

func Floor(f float64) float64 {
	if f < 0 {
		return float64(int(f) - 1)
	}
	return float64(int(f))
}

func Ceil(f float64) float64 {
	if f < 0 {
		return float64(int(f))
	}
	if f == float64(int(f)) {
		return f
	}
	return float64(int(f) + 1)
}

func MinSlice[T cmp.Ordered](arr []T) T {
	if len(arr) == 0 {
		var zero T
		return zero
	}
	min := arr[0]
	for _, v := range arr[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

func MaxSlice[T cmp.Ordered](arr []T) T {
	if len(arr) == 0 {
		var zero T
		return zero
	}
	max := arr[0]
	for _, v := range arr[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func Factorial(n int) int {
	if n < 0 {
		return 0
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func Median[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64](nums ...T) float64 {
	n := len(nums)
	if n == 0 {
		return 0
	}
	sorted := make([]T, n)
	copy(sorted, nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if sorted[i] > sorted[j] {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}
	if n%2 == 0 {
		return float64(sorted[n/2-1]+sorted[n/2]) / 2
	}
	return float64(sorted[n/2])
}
