package errors

import (
	"errors"
	"fmt"
)

func New(msg string) error {
	return errors.New(msg)
}

func Newf(format string, args ...any) error {
	return fmt.Errorf(format, args...)
}

func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", msg, err)
}

func Wrapf(err error, format string, args ...any) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), err)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target any) bool {
	return errors.As(err, target)
}

func Join(errs ...error) error {
	return errors.Join(errs...)
}

type MultiError struct {
	Errors []error
}

func (m *MultiError) Error() string {
	if len(m.Errors) == 0 {
		return ""
	}
	msg := m.Errors[0].Error()
	for _, err := range m.Errors[1:] {
		msg += "; " + err.Error()
	}
	return msg
}

func (m *MultiError) Append(err error) {
	if err != nil {
		m.Errors = append(m.Errors, err)
	}
}

func (m *MultiError) HasError() bool {
	return len(m.Errors) > 0
}

// Unwrap 返回内部错误列表，使 errors.Is / errors.As 能够遍历（Go 1.20+）。
func (m *MultiError) Unwrap() []error {
	return m.Errors
}

// ErrorOrNil 在无错误时返回 nil，避免「非 nil 接口包裹空 *MultiError」的陷阱。
// 建议以此作为函数返回值：return m.ErrorOrNil()。
func (m *MultiError) ErrorOrNil() error {
	if m == nil || len(m.Errors) == 0 {
		return nil
	}
	return m
}

func PanicToError(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case error:
				err = v
			default:
				err = fmt.Errorf("%v", v)
			}
		}
	}()
	fn()
	return
}