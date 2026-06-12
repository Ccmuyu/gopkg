# errors

错误处理工具：创建、包装、判断和合并错误。

## 安装

```go
import "github.com/Ccmuyu/gopkg/errors"
```

## 快速开始

```go
err := errors.New("something went wrong")
err = errors.Wrap(err, "additional context")
err = errors.Wrapf(err, "user %d not found", 42)

errors.Is(err, targetErr)   // true/false
errors.As(err, &targetType) // true/false

joined := errors.Join(err1, err2, err3)

// MultiError 聚合
m := &errors.MultiError{}
m.Append(err1)
m.Append(err2)
m.HasError()  // true

// PanicToError 安全执行
err := errors.PanicToError(func() {
    panic("something")
})
```

## API 参考

```go
func New(msg string) error
func Newf(format string, args ...any) error
func Wrap(err error, msg string) error
func Wrapf(err error, format string, args ...any) error
func Is(err, target error) bool
func As(err error, target any) bool
func Join(errs ...error) error
func PanicToError(fn func()) (err error)

type MultiError struct { Errors []error }
func (m *MultiError) Error() string
func (m *MultiError) Append(err error)
func (m *MultiError) HasError() bool
```
