package cors

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/Ccmuyu/gopkg/test"
)

func TestDefaultAllowsAnyOrigin(t *testing.T) {
	h := Default().Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/api", nil)
	req.Header.Set("Origin", "https://evil.example")
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	AssertEqual(t, rr.Code, http.StatusOK)
	AssertEqual(t, rr.Header().Get("Access-Control-Allow-Origin"), "*")
	AssertEqual(t, rr.Header().Get("Access-Control-Allow-Credentials"), "")
}

func TestSpecificOriginAndCredentials(t *testing.T) {
	c := New(Config{
		AllowOrigins:     []string{"https://app.example.com"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"X-Request-Id"},
	})
	h := c.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/api", nil)
	req.Header.Set("Origin", "https://app.example.com")
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	AssertEqual(t, rr.Header().Get("Access-Control-Allow-Origin"), "https://app.example.com")
	AssertEqual(t, rr.Header().Get("Access-Control-Allow-Credentials"), "true")
	AssertEqual(t, rr.Header().Get("Access-Control-Expose-Headers"), "X-Request-Id")
	AssertTrue(t, containsVary(rr.Header(), "Origin"))
}

func TestDisallowedOriginNoCORSHeaders(t *testing.T) {
	c := New(Config{AllowOrigins: []string{"https://app.example.com"}})
	h := c.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/api", nil)
	req.Header.Set("Origin", "https://other.example.com")
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	AssertEqual(t, rr.Code, http.StatusOK)
	AssertEqual(t, rr.Header().Get("Access-Control-Allow-Origin"), "")
}

func TestPreflightOK(t *testing.T) {
	c := New(Config{
		AllowOrigins: []string{"https://app.example.com"},
		AllowMethods: []string{http.MethodPut},
		AllowHeaders: []string{"Content-Type", "X-Custom"},
		MaxAge:       600,
	})
	h := c.Middleware()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("preflight should not reach next")
	}))

	req := httptest.NewRequest(http.MethodOptions, "/api", nil)
	req.Header.Set("Origin", "https://app.example.com")
	req.Header.Set("Access-Control-Request-Method", http.MethodPut)
	req.Header.Set("Access-Control-Request-Headers", "Content-Type")
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	AssertEqual(t, rr.Code, http.StatusNoContent)
	AssertEqual(t, rr.Header().Get("Access-Control-Allow-Origin"), "https://app.example.com")
	AssertEqual(t, rr.Header().Get("Access-Control-Allow-Methods"), "PUT")
	AssertEqual(t, rr.Header().Get("Access-Control-Allow-Headers"), "Content-Type, X-Custom")
	AssertEqual(t, rr.Header().Get("Access-Control-Max-Age"), "600")
}

func TestPreflightEchoesRequestHeaders(t *testing.T) {
	c := New(Config{AllowOrigins: []string{"https://app.example.com"}})
	h := c.Handler(http.NotFoundHandler())

	req := httptest.NewRequest(http.MethodOptions, "/api", nil)
	req.Header.Set("Origin", "https://app.example.com")
	req.Header.Set("Access-Control-Request-Method", http.MethodPost)
	req.Header.Set("Access-Control-Request-Headers", "X-Foo, X-Bar")
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	AssertEqual(t, rr.Code, http.StatusNoContent)
	AssertEqual(t, rr.Header().Get("Access-Control-Allow-Headers"), "X-Foo, X-Bar")
}

func TestPreflightForbidden(t *testing.T) {
	c := New(Config{AllowOrigins: []string{"https://app.example.com"}})
	h := c.Handler(http.NotFoundHandler())

	req := httptest.NewRequest(http.MethodOptions, "/api", nil)
	req.Header.Set("Origin", "https://evil.example")
	req.Header.Set("Access-Control-Request-Method", http.MethodPost)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	AssertEqual(t, rr.Code, http.StatusForbidden)
}

func TestCredentialsIgnoresWildcard(t *testing.T) {
	c := New(Config{
		AllowOrigins:     []string{"*", "https://app.example.com"},
		AllowCredentials: true,
	})
	h := c.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Origin", "https://unknown.example")
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	AssertEqual(t, rr.Header().Get("Access-Control-Allow-Origin"), "")

	req = httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Origin", "https://app.example.com")
	rr = httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	AssertEqual(t, rr.Header().Get("Access-Control-Allow-Origin"), "https://app.example.com")
	AssertEqual(t, rr.Header().Get("Access-Control-Allow-Credentials"), "true")
}

func TestPrivateNetwork(t *testing.T) {
	c := New(Config{
		AllowOrigins:        []string{"https://app.example.com"},
		AllowPrivateNetwork: true,
	})
	h := c.Handler(http.NotFoundHandler())

	req := httptest.NewRequest(http.MethodOptions, "/", nil)
	req.Header.Set("Origin", "https://app.example.com")
	req.Header.Set("Access-Control-Request-Method", http.MethodGet)
	req.Header.Set("Access-Control-Request-Private-Network", "true")
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	AssertEqual(t, rr.Header().Get("Access-Control-Allow-Private-Network"), "true")
}

func containsVary(h http.Header, name string) bool {
	for _, v := range h.Values("Vary") {
		if v == name {
			return true
		}
	}
	return false
}
