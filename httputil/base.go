package httputil

import "strings"

func NormalizeBase(base string) string {
	if !strings.HasPrefix(base, "/") {
		base = "/" + base
	}
	if !strings.HasSuffix(base, "/") {
		base = base + "/"
	}
	return base
}

func TrimTrailingSlash(base string) string {
	return strings.TrimSuffix(base, "/")
}
