# retry

重试机制：支持最大次数、固定延迟、退避、抖动。

## 安装

```go
import "github.com/Ccmuyu/gopkg/retry"
```

## 快速开始

```go
err := retry.Retry(func() error {
    return doSomething()
}, retry.WithMaxAttempts(5),
   retry.WithDelay(100*time.Millisecond),
   retry.WithBackoff(),
   retry.WithJitter())
```

需要取消或截止时间时使用 `RetryCtx`。已取消的 context 不会执行回调。
`maxAttempts <= 0` 会归一化为 1，确保操作至少执行一次。

```go
err := retry.RetryCtx(ctx, func() error {
    return doSomething()
}, retry.WithMaxAttempts(5))
```

## API 参考

```go
func Retry(fn func() error, opts ...Option) error
func RetryCtx(ctx context.Context, fn func() error, opts ...Option) error
func WithMaxAttempts(n int) Option
func WithDelay(d time.Duration) Option
func WithBackoff() Option
func WithJitter() Option
```
