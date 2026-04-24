package url

import (
	"strings"
	"testing"

	. "github.com/Ccmuyu/gopkg/test"
)

func TestParse(t *testing.T) {
	u, err := Parse("https://example.com/path?query=value")
	AssertTrue(t, err == nil)
	AssertEqual(t, u.Host, "example.com")
}

func TestParseQuery(t *testing.T) {
	m, err := ParseQuery("name=john&age=30")
	AssertTrue(t, err == nil)
	AssertEqual(t, m["name"], "john")
	AssertEqual(t, m["age"], "30")
}

func TestBuild(t *testing.T) {
	s := Build("https://example.com", map[string]string{
		"name": "john",
	})
	AssertTrue(t, Contains(s, "name=john"))
}

func TestHasScheme(t *testing.T) {
	AssertTrue(t, HasScheme("https://example.com"))
	AssertTrue(t, !HasScheme("example.com"))
}

func TestGetHost(t *testing.T) {
	AssertEqual(t, GetHost("https://example.com/path"), "example.com")
}

func TestGetPath(t *testing.T) {
	AssertEqual(t, GetPath("https://example.com/path/to/resource"), "/path/to/resource")
}

func TestEncode(t *testing.T) {
	AssertEqual(t, Encode("hello world"), "hello+world")
}

func TestDecode(t *testing.T) {
	AssertEqual(t, Decode("hello+world"), "hello world")
}

func TestJoinPath(t *testing.T) {
	AssertEqual(t, JoinPath("https://example.com", "api", "v1", "users"), "https://example.com/api/v1/users")
}

func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}