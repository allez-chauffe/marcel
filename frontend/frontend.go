package frontend

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/Zenika/marcel/config"
	"github.com/Zenika/marcel/httputil"
	"github.com/Zenika/marcel/module"
)

func Module() module.Module {
	var base = httputil.NormalizeBase(config.Default().Frontend().BasePath())
	var fs http.FileSystem

	return module.Module{
		Name: "Frontend",
		Start: func(next module.StartNextFunc) (module.StopFunc, error) {
			var err error
			fs, err = initFs()
			if err != nil {
				return nil, err
			}

			return nil, next()
		},
		Http: &module.Http{
			BasePath: httputil.TrimTrailingSlash(base),
			Setup: func(r *mux.Router) {
				r.HandleFunc("/config", configHandler).Methods("GET")
				r.PathPrefix("/").Handler(fileHandler(base, fs))
			},
		},
	}
}

func ConfigureRouter(r *mux.Router) error {
	var base = httputil.NormalizeBase(config.Default().Frontend().BasePath())

	var b = r.PathPrefix(httputil.TrimTrailingSlash(base)).Subrouter()

	b.HandleFunc("/config", configHandler).Methods("GET")
	fh, err := fileHandlerOld(base)
	if err != nil {
		return err
	}
	b.PathPrefix("/").Handler(fh)

	return nil
}

func configHandler(res http.ResponseWriter, req *http.Request) {
	var apiURI = config.Default().Frontend().APIURI()
	if !strings.HasSuffix(apiURI, "/") {
		apiURI += "/"
	}

	if err := json.NewEncoder(res).Encode(struct {
		APIURI string `json:"apiURI"`
	}{
		APIURI: apiURI,
	}); err != nil {
		panic(err)
	}
}

func fileHandler(base string, fs http.FileSystem) http.Handler {
	return http.StripPrefix(
		base,
		http.FileServer(
			httputil.NewTemplater(
				fs,
				[]string{"/index.html"},
				map[string]string{"REACT_APP_BASE": base},
			),
		),
	)
}

func fileHandlerOld(base string) (http.Handler, error) {
	fs, err := initFs()
	if err != nil {
		return nil, err
	}
	return http.StripPrefix(
		base,
		http.FileServer(
			httputil.NewTemplater(
				fs,
				[]string{"/index.html"},
				map[string]string{"REACT_APP_BASE": base},
			),
		),
	), nil
}
