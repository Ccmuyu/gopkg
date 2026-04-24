package ip

import (
	"encoding/json"
	"net"
)

func IsValid(s string) bool {
	return net.ParseIP(s) != nil
}

func IsPrivate(s string) bool {
	ip := net.ParseIP(s)
	if ip == nil {
		return false
	}
	privateBlocks := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
		"127.0.0.0/8",
		"fc00::/7",
		"::1/128",
	}
	for _, block := range privateBlocks {
		_, cidr, _ := net.ParseCIDR(block)
		if cidr.Contains(ip) {
			return true
		}
	}
	return false
}

func ToInt(s string) (int64, error) {
	ip := net.ParseIP(s)
	if ip == nil {
		return 0, nil
	}
	ip4 := ip.To4()
	if ip4 != nil {
		var val int64
		for _, b := range ip4 {
			val = val<<8 + int64(b)
		}
		return val, nil
	}
	return 0, nil
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
	return json.Marshal(map[string]interface{}{
		"ip": s,
	})
}