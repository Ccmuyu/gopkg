package env

import (
	"os"
	"strings"
)

func Get(key string, defaultVal string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return defaultVal
}

func MustGet(key string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		panic("required environment variable not set: " + key)
	}
	return v
}

func Set(key, value string) error {
	return os.Setenv(key, value)
}

func Unset(key string) error {
	return os.Unsetenv(key)
}

func EnvironMap() map[string]string {
	result := make(map[string]string)
	for _, entry := range os.Environ() {
		parts := strings.SplitN(entry, "=", 2)
		if len(parts) == 2 {
			result[parts[0]] = parts[1]
		}
	}
	return result
}

func Has(key string) bool {
	_, ok := os.LookupEnv(key)
	return ok
}
