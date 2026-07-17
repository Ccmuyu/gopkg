# cors

CORS（跨域资源共享）中间件，基于标准库 `net/http`，无第三方依赖。

## 安装

```go
import "github.com/Ccmuyu/gopkg/cors"
```

## 快速开始

```go
// 开发默认：允许任意源
h := cors.Default().Handler(mux)

// 生产：白名单源 + 凭证
c := cors.New(cors.Config{
    AllowOrigins:     []string{"https://app.example.com"},
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
    AllowHeaders:     []string{"Content-Type", "Authorization"},
    ExposeHeaders:    []string{"X-Request-Id"},
    AllowCredentials: true,
    MaxAge:           600,
})
http.ListenAndServe(":8080", c.Middleware()(mux))
```

## 行为说明

- 带 `Access-Control-Request-Method` 的 `OPTIONS` 视为预检：允许则 `204`，否则 `403`。
- `AllowCredentials` 为 true 时忽略 `AllowOrigins` 中的 `"*"`（规范禁止）。
- 未配置 `AllowHeaders` 时，预检会回显 `Access-Control-Request-Headers`。
- 对具体源响应会设置 `Vary: Origin`。

## API 参考

```go
func Default() *CORS
func New(cfg Config) *CORS
func (c *CORS) Middleware() func(http.Handler) http.Handler
func (c *CORS) Handler(next http.Handler) http.Handler

type Config struct {
    AllowOrigins        []string
    AllowMethods        []string
    AllowHeaders        []string
    ExposeHeaders       []string
    AllowCredentials    bool
    MaxAge              int
    AllowPrivateNetwork bool
}
```
