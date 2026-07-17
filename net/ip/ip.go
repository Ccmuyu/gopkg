package ip

import (
	"encoding/json"
	"fmt"
	"net"
)

func IsValid(s string) bool {
	return net.ParseIP(s) != nil
}

// IsPrivate 判断地址是否不应被视为公网目标，包括私网、回环、
// 链路本地和未指定地址。可用于网络访问控制的基础检查。
func IsPrivate(s string) bool {
	ip := net.ParseIP(s)
	if ip == nil {
		return false
	}
	if ip4 := ip.To4(); ip4 != nil {
		ip = ip4
	}
	return ip.IsPrivate() || ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsUnspecified()
}

func ToInt(s string) (int64, error) {
	ip := net.ParseIP(s)
	if ip == nil {
		return 0, fmt.Errorf("invalid IP: %s", s)
	}
	ip4 := ip.To4()
	if ip4 != nil {
		var val int64
		for _, b := range ip4 {
			val = val<<8 + int64(b)
		}
		return val, nil
	}
	return 0, fmt.Errorf("IPv6 not supported: %s", s)
}

// IntToIP 将 [0, 2^32-1] 范围内的整数转换为 IPv4；越界时返回空字符串。
func IntToIP(n int64) string {
	if n < 0 || n > 1<<32-1 {
		return ""
	}
	return net.IP{
		byte(n >> 24),
		byte(n >> 16),
		byte(n >> 8),
		byte(n),
	}.To4().String()
}

func ToJSON(s string) ([]byte, error) {
	return json.Marshal(map[string]any{
		"ip": s,
	})
}
