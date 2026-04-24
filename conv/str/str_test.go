package str

import (
	"testing"

	. "github.com/Ccmuyu/gopkg/test"
)

func TestAtoi(t *testing.T) {
	AssertEqual(t, Atoi("123"), 123)
	AssertEqual(t, Atoi(""), 0)
}

func TestAtoi64(t *testing.T) {
	AssertEqual(t, Atoi64("123"), int64(123))
	AssertEqual(t, Atoi64(""), int64(0))
}

func TestAtoi32(t *testing.T) {
	AssertEqual(t, Atoi32("123"), int32(123))
}

func TestItoa(t *testing.T) {
	AssertEqual(t, Itoa(123), "123")
}

func TestItoa64(t *testing.T) {
	AssertEqual(t, Itoa64(123), "123")
}

func TestParseFloat(t *testing.T) {
	AssertEqual(t, ParseFloat("3.14"), 3.14)
	AssertEqual(t, ParseFloat(""), 0.0)
}

func TestFormatFloat(t *testing.T) {
	AssertEqual(t, FormatFloat(3.14), "3.14")
}

func TestParseBool(t *testing.T) {
	AssertTrue(t, ParseBool("true"))
	AssertTrue(t, !ParseBool(""))
}

func TestFormatBool(t *testing.T) {
	AssertEqual(t, FormatBool(true), "true")
	AssertEqual(t, FormatBool(false), "false")
}