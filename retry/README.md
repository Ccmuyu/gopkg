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

## API 参考

```go
func Retry(fn func() error, opts ...Option) error
func WithMaxAttempts(n int) Option
func WithDelay(d time.Duration) Option
func WithBackoff() Option
func WithJitter() Option
```
