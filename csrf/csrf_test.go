package csrf

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	. "github.com/Ccmuyu/gopkg/test"
)

func TestGenerateToken(t *testing.T) {
	p := New(Config{})
	tok, err := p.Generate()
	AssertNoError(t, err)
	AssertTrue(t, len(tok) > 0)

	tok2, err := p.Generate()
	AssertNoError(t, err)
	AssertTrue(t, tok != tok2)
}

func TestSafeMethodSetsCookie(t *testing.T) {
	p := New(Config{})
	h := p.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	AssertEqual(t, rr.Code, http.StatusOK)
	cookie := rr.Result().Cookies()
	AssertTrue(t, len(cookie) >= 1)
	AssertEqual(t, cookie[0].Name, "csrf_token")
	AssertTrue(t, cookie[0].Value != "")
}

func TestUnsafeWithoutTokenForbidden(t *testing.T) {
	p := New(Config{})
	h := p.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("should not reach next")
	}))

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	AssertEqual(t, rr.Code, http.StatusForbidden)
}

func TestUnsafeWithMatchingHeader(t *testing.T) {
	p := New(Config{})
	token, err := p.Generate()
	AssertNoError(t, err)

	called := false
	h := p.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusCreated)
	}))

	req := httptest.NewRequest(http.MethodPost, "/submit", nil)
	req.AddCookie(&http.Cookie{Name: "csrf_token", Value: token})
	req.Header.Set("X-CSRF-Token", token)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	AssertTrue(t, called)
	AssertEqual(t, rr.Code, http.StatusCreated)
}

func TestUnsafeWithFormField(t *testing.T) {
	p := New(Config{})
	token, err := p.Generate()
	AssertNoError(t, err)

	form := url.Values{}
	form.Set("csrf_token", token)
	req := httptest.NewRequest(http.MethodPost, "/submit", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.AddCookie(&http.Cookie{Name: "csrf_token", Value: token})

	h := p.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	AssertEqual(t, rr.Code, http.StatusOK)
}

func TestMismatchedToken(t *testing.T) {
	p := New(Config{})
	token, err := p.Generate()
	AssertNoError(t, err)

	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.AddCookie(&http.Cookie{Name: "csrf_token", Value: token})
	req.Header.Set("X-CSRF-Token", "wrong-token")

	h := p.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("should not reach next")
	}))
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	AssertEqual(t, rr.Code, http.StatusForbidden)
}

func TestTrustedOrigin(t *testing.T) {
	p := New(Config{
		TrustedOrigins: []string{"https://app.example.com"},
	})
	token, err := p.Generate()
	AssertNoError(t, err)

	h := p.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.AddCookie(&http.Cookie{Name: "csrf_token", Value: token})
	req.Header.Set("X-CSRF-Token", token)
	req.Header.Set("Origin", "https://evil.example.com")
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	AssertEqual(t, rr.Code, http.StatusForbidden)

	req = httptest.NewRequest(http.MethodPost, "/", nil)
	req.AddCookie(&http.Cookie{Name: "csrf_token", Value: token})
	req.Header.Set("X-CSRF-Token", token)
	req.Header.Set("Origin", "https://app.example.com")
	rr = httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	AssertEqual(t, rr.Code, http.StatusOK)
}

func TestCustomErrorHandler(t *testing.T) {
	var got error
	p := New(Config{
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
			got = err
			w.WriteHeader(http.StatusTeapot)
		},
	})
	h := p.Middleware()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("should not reach next")
	}))

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	AssertEqual(t, rr.Code, http.StatusTeapot)
	AssertTrue(t, errors.Is(got, ErrNoCookie) || errors.Is(got, ErrMissingToken))
}

func TestTokenReusesCookie(t *testing.T) {
	p := New(Config{})
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.AddCookie(&http.Cookie{Name: "csrf_token", Value: "existing"})
	tok, err := p.Token(req)
	AssertNoError(t, err)
	AssertEqual(t, tok, "existing")
}

func TestValidHelpers(t *testing.T) {
	p := New(Config{})
	token, err := p.Generate()
	AssertNoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	AssertTrue(t, errors.Is(p.Valid(req), ErrNoCookie))

	req.AddCookie(&http.Cookie{Name: "csrf_token", Value: token})
	AssertTrue(t, errors.Is(p.Valid(req), ErrMissingToken))

	req.Header.Set("X-CSRF-Token", token)
	AssertNoError(t, p.Valid(req))
}
