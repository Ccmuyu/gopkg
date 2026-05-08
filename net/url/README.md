# net/url

URL 解析与构建工具。

## 安装

```go
import "github.com/Ccmuyu/gopkg/net/url"
```

## 快速开始

```go
u, _ := url.Parse("https://example.com/path?key=val")
url.ParseQuery("key=val&foo=bar")          // {"key":"val","foo":"bar"}, nil
url.Build("https://example.com", map[string]string{"q":"hello"})
// https://example.com?q=hello

url.HasScheme("https://example.com")       // true
url.GetHost("https://example.com/path")    // "example.com"
url.GetPath("https://example.com/path")    // "/path"
url.Encode("a b")                          // "a+b"
url.Decode("a+b")                          // "a b"
url.JoinPath("https://example.com", "api", "v1") // https://example.com/api/v1
```

## API 参考

```go
func Parse(s string) (*url.URL, error)
func ParseQuery(s string) (map[string]string, error)
func Build(base string, params map[string]string) string
func HasScheme(s string) bool
func GetHost(s string) string
func GetPath(s string) string
func Encode(s string) string
func Decode(s string) string
func JoinPath(base string, parts ...string) string
```
