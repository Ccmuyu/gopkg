package ip

import (
	"encoding/json"
	"fmt"
	"net"
)

func IsValid(s string) bool {
	return net.ParseIP(s) != nil
}

var privateBlocks = []*net.IPNet{
	mustCIDR("10.0.0.0/8"),
	mustCIDR("172.16.0.0/12"),
	mustCIDR("192.168.0.0/16"),
	mustCIDR("127.0.0.0/8"),
	mustCIDR("fc00::/7"),
	mustCIDR("::1/128"),
}

func IsPrivate(s string) bool {
	ip := net.ParseIP(s)
	if ip == nil {
		return false
	}
	for _, cidr := range privateBlocks {
		if cidr.Contains(ip) {
			return true
		}
	}
	return false
}

func mustCIDR(s string) *net.IPNet {
	_, cidr, err := net.ParseCIDR(s)
	if err != nil {
		panic("invalid CIDR: " + s)
	}
	return cidr
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

func IntToIP(n int64) string {
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