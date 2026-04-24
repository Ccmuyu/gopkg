package errors

import (
	"testing"

	. "github.com/Ccmuyu/gopkg/test"
)

func TestNew(t *testing.T) {
	err := New("test error")
	AssertTrue(t, err != nil)
	AssertEqual(t, err.Error(), "test error")
}

func TestNewf(t *testing.T) {
	err := Newf("error: %d", 123)
	AssertEqual(t, err.Error(), "error: 123")
}

func TestWrap(t *testing.T) {
	orig := New("original")
	wrapped := Wrap(orig, "wrapped")
	AssertTrue(t, Is(wrapped, orig))
}

func TestWrapf(t *testing.T) {
	orig := New("original")
	wrapped := Wrapf(orig, "wrapped: %s", "info")
	AssertTrue(t, Is(wrapped, orig))
	AssertTrue(t, Contains(wrapped.Error(), "wrapped"))
}

func TestIs(t *testing.T) {
	sentinel := New("sentinel")
	wrapped := Wrap(sentinel, "context")
	AssertTrue(t, Is(wrapped, sentinel))
}

func TestAs(t *testing.T) {
	var terr *parseError
	err := &parseError{msg: "parse error: invalid"}
	As(Wrap(err, "wrapped"), &terr)
	AssertTrue(t, terr != nil)
}

func TestJoin(t *testing.T) {
	err1 := New("one")
	err2 := New("two")
	joined := Join(err1, err2)
	AssertTrue(t, joined != nil)
}

func Contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

type parseError struct {
	msg string
}

func (e *parseError) Error() string {
	return e.msg
}