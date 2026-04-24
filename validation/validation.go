package validation

import (
	"net"
	"net/mail"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var (
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	phoneRegex = regexp.MustCompile(`^1[3-9]\d{9}$`)
	urlRegex   = regexp.MustCompile(`^https?://[^\s/$.?#].[^\s]*$`)
	ipv4Regex  = regexp.MustCompile(`^(\d{1,3}\.){3}\d{1,3}$`)
)

func IsEmail(s string) bool {
	if len(s) > 320 {
		return false
	}
	return emailRegex.MatchString(s)
}

func IsPhone(s string) bool {
	return phoneRegex.MatchString(s)
}

func IsURL(s string) bool {
	return urlRegex.MatchString(s)
}

func IsIP(s string) bool {
	return net.ParseIP(s) != nil
}

func IsIPv4(s string) bool {
	if !ipv4Regex.MatchString(s) {
		return false
	}
	parts := strings.Split(s, ".")
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		if n < 0 || n > 255 {
			return false
		}
	}
	return true
}

func IsAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func IsAlphanumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func IsMatch(s string, pattern string) bool {
	matched, _ := regexp.MatchString(pattern, s)
	return matched
}

func IsEmailAddr(s string) bool {
	_, err := mail.ParseAddress(s)
	return err == nil
}

func IsPrivateIP(s string) bool {
	ip := net.ParseIP(s)
	if ip == nil {
		return false
	}
	privateBlocks := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
		"127.0.0.0/8",
	}
	for _, block := range privateBlocks {
		_, cidr, _ := net.ParseCIDR(block)
		if cidr.Contains(ip) {
			return true
		}
	}
	return false
}