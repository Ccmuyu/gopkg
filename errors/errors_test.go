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

type parseError struct {
	msg string
}

func (e *parseError) Error() string {
	return e.msg
}