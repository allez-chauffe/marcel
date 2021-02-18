package catchall

import (
	"errors"
	"io/fs"
)

type catchAll struct {
	fs          fs.FS
	defaultPath string
}

var _ fs.FS = (*catchAll)(nil)

func (r *catchAll) Open(path string) (fs.File, error) {
	f, err := r.fs.Open(path)
	if errors.Is(err, fs.ErrNotExist) {
		return r.fs.Open(r.defaultPath)
	}
	return f, err
}

func New(fs fs.FS, defaultPath string) fs.FS {
	return &catchAll{fs, defaultPath}
}
