package templater

import (
	"bytes"
	"io"
	"io/fs"
	"text/template"
)

type templater struct {
	fs       fs.FS
	includes map[string]bool
	data     interface{}
	cache    map[string]fs.File
}

var _ fs.FS = (*templater)(nil)

func (tfs *templater) Open(path string) (fs.File, error) {
	if !tfs.includes[path] {
		return tfs.fs.Open(path)
	}

	if f, ok := tfs.cache[path]; ok {
		return f, nil
	}

	f, err := tfs.fs.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return nil, err
	}

	var buf = bytes.NewBuffer(make([]byte, 0, info.Size()))
	if _, err = io.Copy(buf, f); err != nil {
		return nil, err
	}

	tmpl, err := template.New(path).Parse(string(buf.Bytes()))
	if err != nil {
		return nil, err
	}

	buf.Reset()

	if err = tmpl.Execute(buf, tfs.data); err != nil {
		return nil, err
	}

	bf := newBfile(buf.Bytes(), info)
	tfs.cache[path] = bf

	return bf, nil
}

// New creates a new templater fs.FS around wrapped fs.FS.
// Templater will execute files in includes as templates with data.
func New(wrapped fs.FS, includes []string, data interface{}) fs.FS {
	tfs := &templater{wrapped, make(map[string]bool, len(includes)), data, make(map[string]fs.File, len(includes))}

	for _, path := range includes {
		tfs.includes[path] = true
	}

	return tfs
}

type bfile struct {
	*bytes.Reader
	fs.FileInfo
	size int64
}

var _ fs.File = (*bfile)(nil)

func (f *bfile) Stat() (fs.FileInfo, error) {
	return f, nil
}

func (f *bfile) Close() error {
	_, err := f.Reader.Seek(0, io.SeekStart)
	return err
}

var _ fs.FileInfo = (*bfile)(nil)

func (f *bfile) Size() int64 {
	return f.size
}

func newBfile(b []byte, info fs.FileInfo) fs.File {
	return &bfile{bytes.NewReader(b), info, int64(len(b))}
}
