package httputil

import (
	"io/fs"
	"os"
)

type notFoundRewriter struct {
	fs  fs.FS
	url string
}

var _ fs.FS = (*notFoundRewriter)(nil)

func (r *notFoundRewriter) Open(path string) (fs.File, error) {
	f, err := r.fs.Open(path)
	if err != nil && os.IsNotExist(err) {
		return r.fs.Open(r.url)
	}
	return f, err
}

func NewNotFoundRewriter(fs fs.FS, url string) fs.FS {
	return &notFoundRewriter{fs, url}
}
