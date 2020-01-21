package httputil

import "strings"

// NormalizeBase adds a leading slash and a trailing slash if missing.
func NormalizeBase(base string) string {
	if base == "" {
		return ""
	}
	if !strings.HasPrefix(base, "/") {
		base = "/" + base
	}
	if !strings.HasSuffix(base, "/") {
		base = base + "/"
	}
	return base
}

// TrimTrailingSlash trims the trailing slash if present.
func TrimTrailingSlash(base string) string {
	return strings.TrimSuffix(base, "/")
}
