// Package ip provides utilities for IP address detection.
// Reference: https://github.com/yunkon-kim/golang-playground/blob/master/inquire-public-ip/inquire-public-ip.go
package ip

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

// GetClientIP extracts the visitor's IP from an HTTP request.
func GetClientIP(r *http.Request) string {
	// Check X-Forwarded-For (set by reverse proxies)
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		parts := strings.Split(forwarded, ",")
		return strings.TrimSpace(parts[0])
	}

	// Check X-Real-Ip header
	if realIP := r.Header.Get("X-Real-Ip"); realIP != "" {
		return realIP
	}

	// Fall back to RemoteAddr
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}

// GetPublicIP queries external services to get this server's public IP.
func GetPublicIP() (string, error) {
	urls := []string{
		"https://api.ipify.org?format=text",
		"https://ifconfig.co/",
		"http://api.ident.me",
	}

	for _, url := range urls {
		result, err := fetchIP(url)
		if err != nil {
			continue
		}
		return result, nil
	}

	return "", fmt.Errorf("failed to get public IP from all sources")
}

// fetchIP sends a GET request to the given URL and parses the IP from the response.
func fetchIP(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	trimmed := strings.TrimSpace(string(data))
	if net.ParseIP(trimmed) == nil {
		return "", fmt.Errorf("invalid IP format: %s", trimmed)
	}

	return trimmed, nil
}
