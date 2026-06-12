package errors

import (
	"testing"

	gtest "github.com/Ccmuyu/gopkg/test"
)

func TestNew(t *testing.T) {
	err := New("test error")
	gtest.AssertTrue(t, err != nil)
	gtest.AssertEqual(t, err.Error(), "test error")
}

func TestNewf(t *testing.T) {
	err := Newf("error: %d", 123)
	gtest.AssertEqual(t, err.Error(), "error: 123")
}

func TestWrap(t *testing.T) {
	orig := New("original")
	wrapped := Wrap(orig, "wrapped")
	gtest.AssertTrue(t, Is(wrapped, orig))
}

func TestWrapf(t *testing.T) {
	orig := New("original")
	wrapped := Wrapf(orig, "wrapped: %s", "info")
	gtest.AssertTrue(t, Is(wrapped, orig))
	gtest.AssertTrue(t, gtest.ContainsStr(wrapped.Error(), "wrapped"))
}

func TestIs(t *testing.T) {
	sentinel := New("sentinel")
	wrapped := Wrap(sentinel, "context")
	gtest.AssertTrue(t, Is(wrapped, sentinel))
}

func TestAs(t *testing.T) {
	var terr *parseError
	err := &parseError{msg: "parse error: invalid"}
	As(Wrap(err, "wrapped"), &terr)
	gtest.AssertTrue(t, terr != nil)
}

func TestJoin(t *testing.T) {
	err1 := New("one")
	err2 := New("two")
	joined := Join(err1, err2)
	gtest.AssertTrue(t, joined != nil)
}

func TestMultiError(t *testing.T) {
	m := &MultiError{}
	gtest.AssertTrue(t, !m.HasError())
	gtest.AssertEqual(t, m.Error(), "")

	m.Append(New("err1"))
	m.Append(nil)
	m.Append(New("err2"))
	gtest.AssertTrue(t, m.HasError())
	gtest.AssertTrue(t, gtest.ContainsStr(m.Error(), "err1"))
	gtest.AssertTrue(t, gtest.ContainsStr(m.Error(), "err2"))
}

func TestPanicToError(t *testing.T) {
	err := PanicToError(func() {
		panic("something went wrong")
	})
	gtest.AssertTrue(t, err != nil)
	gtest.AssertEqual(t, err.Error(), "something went wrong")
}

func TestPanicToErrorNoPanic(t *testing.T) {
	err := PanicToError(func() {
		// no panic
	})
	gtest.AssertTrue(t, err == nil)
}

func TestPanicToErrorWithError(t *testing.T) {
	sentinel := New("sentinel")
	err := PanicToError(func() {
		panic(sentinel)
	})
	gtest.AssertTrue(t, Is(err, sentinel))
}

type parseError struct {
	msg string
}

func (e *parseError) Error() string {
	return e.msg
}