# csrf

基于**双重提交 Cookie** 的 CSRF 防护：服务端下发 Cookie，客户端在不安全请求中通过请求头或表单回传同一令牌。

## 安装

```go
import "github.com/Ccmuyu/gopkg/csrf"
```

## 快速开始

```go
p := csrf.New(csrf.Config{
    Secure:   true,                 // 生产环境建议开启
    SameSite: http.SameSiteLaxMode,
    TrustedOrigins: []string{       // 可选：额外校验 Origin
        "https://app.example.com",
    },
})

mux := http.NewServeMux()
mux.HandleFunc("/form", handleForm)
http.ListenAndServe(":8080", p.Middleware()(mux))
```

前端示例：

```js
// Cookie 名默认 csrf_token；需 HTTPOnly=false（默认）才能由 JS 读取
const token = document.cookie
  .split("; ")
  .find((c) => c.startsWith("csrf_token="))
  ?.split("=")[1];

fetch("/api/submit", {
  method: "POST",
  headers: { "X-CSRF-Token": token },
  credentials: "include",
  body: JSON.stringify(data),
});
```

## 行为说明

- 安全方法（`GET` / `HEAD` / `OPTIONS` / `TRACE`）：仅确保响应带上 CSRF Cookie。
- 不安全方法：要求 Cookie 与 `X-CSRF-Token`（或表单字段 `csrf_token`）常量时间相等。
- 配置 `TrustedOrigins` 后，还会校验 `Origin`（缺失则尝试 `Referer`）。
- 校验失败默认返回 **403**；可用 `ErrorHandler` 自定义。

## API 参考

```go
func New(cfg Config) *Protector
func (p *Protector) Generate() (string, error)
func (p *Protector) Token(r *http.Request) (string, error)
func (p *Protector) SetCookie(w http.ResponseWriter, token string)
func (p *Protector) Valid(r *http.Request) error
func (p *Protector) Middleware() func(http.Handler) http.Handler
func (p *Protector) Handler(next http.Handler) http.Handler

var (
    ErrMissingToken error
    ErrInvalidToken error
    ErrNoCookie     error
)
```
