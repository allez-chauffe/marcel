package httputil

import (
	"net/http"
	"os"
)

type notFoundRewriter struct {
	fs  http.FileSystem
	url string
}

var _ http.FileSystem = (*notFoundRewriter)(nil)

func (r *notFoundRewriter) Open(path string) (http.File, error) {
	f, err := r.fs.Open(path)
	if err != nil && os.IsNotExist(err) {
		return r.fs.Open(r.url)
	}
	return f, err
}

func NewNotFoundRewriter(fs http.FileSystem, url string) http.FileSystem {
	return &notFoundRewriter{fs, url}
}
