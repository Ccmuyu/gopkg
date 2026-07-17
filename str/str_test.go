package str

import (
	"testing"

	. "github.com/Ccmuyu/gopkg/test"
)

func TestTrim(t *testing.T) {
	AssertEqual(t, Trim("  hello  "), "hello")
}

func TestUpper(t *testing.T) {
	AssertEqual(t, Upper("hello"), "HELLO")
}

func TestLower(t *testing.T) {
	AssertEqual(t, Lower("HELLO"), "hello")
}

func TestSplit(t *testing.T) {
	AssertSliceEqual(t, Split("a,b,c", ","), []string{"a", "b", "c"})
	AssertSliceEqual(t, Split("", ","), []string{})
}

func TestSplitN(t *testing.T) {
	AssertSliceEqual(t, SplitN("a,b,c", ",", 2), []string{"a", "b,c"})
}

func TestJoin(t *testing.T) {
	AssertEqual(t, Join([]string{"a", "b", "c"}, ","), "a,b,c")
}

func TestContains(t *testing.T) {
	AssertTrue(t, Contains("hello", "ell"))
	AssertTrue(t, !Contains("hello", "xxx"))
}

func TestHasPrefix(t *testing.T) {
	AssertTrue(t, HasPrefix("hello", "hel"))
	AssertTrue(t, !HasPrefix("hello", "xxx"))
}

func TestHasSuffix(t *testing.T) {
	AssertTrue(t, HasSuffix("hello", "llo"))
	AssertTrue(t, !HasSuffix("hello", "xxx"))
}

func TestReplace(t *testing.T) {
	AssertEqual(t, Replace("hello", "l", "L"), "heLLo")
}

func TestEllipsis(t *testing.T) {
	AssertEqual(t, Ellipsis("hello world", 5), "hello...")
	AssertEqual(t, Ellipsis("hi", 5), "hi")
	AssertEqual(t, Ellipsis("你好世界", 2), "你好...")
	AssertEqual(t, Ellipsis("hello", -1), "...")
}

func TestIsEmpty(t *testing.T) {
	AssertTrue(t, IsEmpty(""))
	AssertTrue(t, !IsEmpty("a"))
}

func TestDefaultIfEmpty(t *testing.T) {
	AssertEqual(t, DefaultIfEmpty("", "default"), "default")
	AssertEqual(t, DefaultIfEmpty("value", "default"), "value")
}

func TestRepeat(t *testing.T) {
	AssertEqual(t, Repeat("ab", 3), "ababab")
}

func TestContainsAny(t *testing.T) {
	AssertTrue(t, ContainsAny("hello", "aeiou"))
	AssertTrue(t, !ContainsAny("hll", "aeiou"))
}

func TestCount(t *testing.T) {
	AssertEqual(t, Count("hello", "l"), 2)
}

func TestAfter(t *testing.T) {
	AssertEqual(t, After("hello@example.com", "@"), "example.com")
	AssertEqual(t, After("hello", "@"), "")
}

func TestBefore(t *testing.T) {
	AssertEqual(t, Before("hello@example.com", "@"), "hello")
	AssertEqual(t, Before("hello", "@"), "")
}

func TestBetween(t *testing.T) {
	AssertEqual(t, Between("{{hello}}", "{{", "}}"), "hello")
	AssertEqual(t, Between("hello", "{{", "}}"), "")
	AssertEqual(t, Between("hello(world)", "(", ")"), "world")
}

func TestReverse(t *testing.T) {
	AssertEqual(t, Reverse("hello"), "olleh")
	AssertEqual(t, Reverse(""), "")
	AssertEqual(t, Reverse("a"), "a")
	AssertEqual(t, Reverse("你好"), "好你")
}

func TestPadLeft(t *testing.T) {
	AssertEqual(t, PadLeft("42", 5, "0"), "00042")
	AssertEqual(t, PadLeft("hello", 3, " "), "hello")
	AssertEqual(t, PadLeft("你好", 4, " "), "  你好")
}

func TestPadRight(t *testing.T) {
	AssertEqual(t, PadRight("42", 5, "0"), "42000")
	AssertEqual(t, PadRight("hello", 3, " "), "hello")
}

func TestTruncate(t *testing.T) {
	AssertEqual(t, Truncate("hello world", 5), "hello")
	AssertEqual(t, Truncate("hi", 5), "hi")
	AssertEqual(t, Truncate("hello", -1), "hello")
	AssertEqual(t, Truncate("你好世界", 2), "你好")
}

func TestCapitalize(t *testing.T) {
	AssertEqual(t, Capitalize("hello"), "Hello")
	AssertEqual(t, Capitalize("HELLO"), "Hello")
	AssertEqual(t, Capitalize(""), "")
}

func TestToCamel(t *testing.T) {
	AssertEqual(t, ToCamel("hello_world"), "helloWorld")
	AssertEqual(t, ToCamel("hello-world"), "helloWorld")
	AssertEqual(t, ToCamel("HelloWorld"), "helloWorld")
	AssertEqual(t, ToCamel("hello world"), "helloWorld")
}

func TestToSnake(t *testing.T) {
	AssertEqual(t, ToSnake("helloWorld"), "hello_world")
	AssertEqual(t, ToSnake("HelloWorld"), "hello_world")
	AssertEqual(t, ToSnake("hello world"), "hello_world")
	AssertEqual(t, ToSnake("hello"), "hello")
}

func TestToKebab(t *testing.T) {
	AssertEqual(t, ToKebab("helloWorld"), "hello-world")
	AssertEqual(t, ToKebab("HelloWorld"), "hello-world")
}

func TestMask(t *testing.T) {
	AssertEqual(t, Mask("12345678901", 3, '*'), "123********")
	AssertEqual(t, Mask("123", 3, '*'), "123")
	AssertEqual(t, Mask("12345", 10, '*'), "12345")
	AssertEqual(t, Mask("hello世界", 2, '*'), "he*****")
	AssertEqual(t, Mask("123", -1, '*'), "***")
}
