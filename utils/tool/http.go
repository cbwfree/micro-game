package tool

import (
	"net"
	"net/http"
	"strings"
)

// GetHttpRealIP 获取真实IP地址
func GetHttpRealIP(r *http.Request) string {
	// Fall back to legacy behavior
	if ip := r.Header.Get("X-FORWARDED-FOR"); ip != "" {
		return strings.Split(ip, ", ")[0]
	}
	if ip := r.Header.Get("X-REAL-IP"); ip != "" {
		return ip
	}
	ra, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ra
}
