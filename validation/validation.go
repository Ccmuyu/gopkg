package validation

import (
	"encoding/json"
	"net"
	"net/mail"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"unicode"
)

var (
	emailRegex       = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	phoneRegex       = regexp.MustCompile(`^1[3-9]\d{9}$`)
	urlRegex         = regexp.MustCompile(`^https?://[^\s/$.?#].[^\s]*$`)
	ipv4Regex        = regexp.MustCompile(`^(\d{1,3}\.){3}\d{1,3}$`)
	uuidRegex        = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)
	macRegex         = regexp.MustCompile(`^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$`)
	hexColorRegex    = regexp.MustCompile(`^#([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$`)
	postalCodeRegex  = regexp.MustCompile(`^\d{6}$`)
	creditCardRegex  = regexp.MustCompile(`^[0-9]{13,19}$`)
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

func IsUUID(s string) bool {
	return uuidRegex.MatchString(s)
}

func IsMAC(s string) bool {
	return macRegex.MatchString(s)
}

func IsJSON(s string) bool {
	var js any
	return json.Unmarshal([]byte(s), &js) == nil
}

func IsHexColor(s string) bool {
	return hexColorRegex.MatchString(s)
}

func IsPostalCode(s string) bool {
	return postalCodeRegex.MatchString(s)
}

func IsPort(port int) bool {
	return port > 0 && port <= 65535
}

func IsLatitude(lat float64) bool {
	return lat >= -90 && lat <= 90
}

func IsLongitude(lng float64) bool {
	return lng >= -180 && lng <= 180
}

func IsCreditCard(s string) bool {
	if !creditCardRegex.MatchString(s) {
		return false
	}
	return luhnCheck(s)
}

func luhnCheck(s string) bool {
	var sum int
	parity := len(s) % 2
	for i, r := range s {
		digit := int(r - '0')
		if i%2 == parity {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	return sum%10 == 0
}

var (
	regexCache   sync.Map
	regexCacheMu sync.Mutex
	regexCount   int
)

const regexCacheMax = 1024

func IsMatch(s string, pattern string) bool {
	if cached, ok := regexCache.Load(pattern); ok {
		return cached.(*regexp.Regexp).MatchString(s)
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}
	regexCacheMu.Lock()
	if regexCount >= regexCacheMax {
		regexCacheMu.Unlock()
		return re.MatchString(s)
	}
	regexCount++
	regexCacheMu.Unlock()
	regexCache.Store(pattern, re)
	return re.MatchString(s)
}

func IsEmailAddr(s string) bool {
	_, err := mail.ParseAddress(s)
	return err == nil
}

var privateBlocks = []*net.IPNet{
	mustCIDR("10.0.0.0/8"),
	mustCIDR("172.16.0.0/12"),
	mustCIDR("192.168.0.0/16"),
	mustCIDR("127.0.0.0/8"),
}

func IsPrivateIP(s string) bool {
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