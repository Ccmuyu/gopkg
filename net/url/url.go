package url

import (
	"net/url"
	"strings"
)

func Parse(s string) (*url.URL, error) {
	return url.Parse(s)
}

func ParseQuery(s string) (map[string]string, error) {
	m, err := url.ParseQuery(s)
	if err != nil {
		return nil, err
	}
	result := make(map[string]string)
	for k, v := range m {
		result[k] = v[0]
	}
	return result, nil
}

func Build(base string, params map[string]string) string {
	u, err := url.Parse(base)
	if err != nil {
		return ""
	}
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func HasScheme(s string) bool {
	u, err := url.Parse(s)
	return err == nil && u.Scheme != ""
}

func GetHost(s string) string {
	u, err := url.Parse(s)
	if err != nil {
		return ""
	}
	return u.Host
}

func GetPath(s string) string {
	u, err := url.Parse(s)
	if err != nil {
		return ""
	}
	return u.Path
}

func Encode(s string) string {
	return url.QueryEscape(s)
}

func Decode(s string) string {
	decoded, err := url.QueryUnescape(s)
	if err != nil {
		return s
	}
	return decoded
}

func JoinPath(base string, parts ...string) string {
	u, err := url.Parse(base)
	if err != nil {
		return ""
	}
	urlParts := append([]string{u.Path}, parts...)
	u.Path = strings.Join(urlParts, "/")
	return u.String()
}