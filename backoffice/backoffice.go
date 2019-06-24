package backoffice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"

	"github.com/Zenika/marcel/config"
	"github.com/Zenika/marcel/httputil"
)

var (
	fileHandler = http.FileServer(httputil.NewNotFoundRewriter(&templater{packr.NewBox("./build/")}, "/index.html"))

	backofficeConfig = struct {
		BackendURI  string `json:"backendURI"`
		FrontendURI string `json:"frontendURI"`
	}{
		BackendURI:  "/api/v1/",
		FrontendURI: "/front/",
	}
)

func Start(base string) error {
	base = strings.TrimSuffix(base, "/")

	r := mux.NewRouter()

	baseRoute := r.PathPrefix(base)
	baseRouter := baseRoute.Subrouter()

	baseRouter.Handle("", http.RedirectHandler(base+"/", http.StatusMovedPermanently))
	baseRouter.HandleFunc("/config", configHandler).Methods("GET")
	baseRouter.PathPrefix("/").Handler(http.StripPrefix(base, fileHandler))

	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), r)
}

func configHandler(res http.ResponseWriter, req *http.Request) {
	if err := json.NewEncoder(res).Encode(backofficeConfig); err != nil {
		panic(err)
	}
}

type templater struct {
	http.FileSystem
}

var _ http.FileSystem = (*templater)(nil)

func (fs *templater) Open(path string) (http.File, error) {
	if path == "/index.html" {
		f, err := fs.FileSystem.Open(path)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		var buf = new(bytes.Buffer)
		if _, err = io.Copy(buf, f); err != nil {
			return nil, err
		}

		tmpl, err := template.New("/index.html").Parse(string(buf.Bytes()))
		if err != nil {
			return nil, err
		}

		buf.Reset()

		if err = tmpl.Execute(buf, map[string]string{"REACT_APP_BASE_URL": "/test/"}); err != nil {
			return nil, err
		}

		content := buf.Bytes()

		info, err := f.Stat()
		if err != nil {
			return nil, err
		}

		return &file{*bytes.NewReader(content), fileInfo{info, int64(len(content))}}, nil
	}

	return fs.FileSystem.Open(path)
}

type file struct {
	bytes.Reader
	info fileInfo
}

func (*file) Close() error {
	return nil
}

func (*file) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *file) Stat() (os.FileInfo, error) {
	return f.info, nil
}

var _ http.File = (*file)(nil)

type fileInfo struct {
	info os.FileInfo
	size int64
}

var _ os.FileInfo = fileInfo{}

func (i fileInfo) Name() string {
	return i.info.Name()
}

func (i fileInfo) Size() int64 {
	return i.size
}

func (i fileInfo) Mode() os.FileMode {
	return i.info.Mode()
}

func (i fileInfo) ModTime() time.Time {
	return i.info.ModTime()
}

func (i fileInfo) IsDir() bool {
	return i.info.IsDir()
}

func (i fileInfo) Sys() interface{} {
	return i.info.Sys()
}
