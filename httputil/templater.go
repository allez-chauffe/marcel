package httputil

import (
	"bytes"
	"io"
	"io/fs"
	"net/http"
	"os"
	"text/template"
	"time"
)

type templater struct {
	fs       fs.FS
	includes map[string]bool
	data     interface{}
}

var _ fs.FS = (*templater)(nil)

func (t *templater) Open(path string) (fs.File, error) {
	if !t.includes[path] {
		return t.fs.Open(path)
	}

	// FIXME add a cache

	f, err := t.fs.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var buf = new(bytes.Buffer)
	if _, err = io.Copy(buf, f); err != nil {
		return nil, err
	}

	tmpl, err := template.New(path).Parse(string(buf.Bytes()))
	if err != nil {
		return nil, err
	}

	buf.Reset()

	if err = tmpl.Execute(buf, t.data); err != nil {
		return nil, err
	}

	info, err := f.Stat()
	if err != nil {
		return nil, err
	}

	return newBfile(buf.Bytes(), info), nil
}

func NewTemplater(fs fs.FS, includes []string, data interface{}) fs.FS {
	t := &templater{fs, make(map[string]bool, len(includes)), data}

	for _, path := range includes {
		t.includes[path] = true
	}

	return t
}

type bfile struct {
	bytes.Reader
	info os.FileInfo
}

func (*bfile) Close() error {
	return nil
}

func (*bfile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *bfile) Stat() (os.FileInfo, error) {
	return f.info, nil
}

var _ http.File = (*bfile)(nil)

type bfileInfo struct {
	info os.FileInfo
	size int64
}

var _ os.FileInfo = bfileInfo{}

func (i bfileInfo) Name() string {
	return i.info.Name()
}

func (i bfileInfo) Size() int64 {
	return i.size
}

func (i bfileInfo) Mode() os.FileMode {
	return i.info.Mode()
}

func (i bfileInfo) ModTime() time.Time {
	return i.info.ModTime()
}

func (i bfileInfo) IsDir() bool {
	return i.info.IsDir()
}

func (i bfileInfo) Sys() interface{} {
	return i.info.Sys()
}

func newBfile(b []byte, info os.FileInfo) http.File {
	return &bfile{
		*bytes.NewReader(b),
		bfileInfo{
			info,
			int64(len(b)),
		},
	}
}
