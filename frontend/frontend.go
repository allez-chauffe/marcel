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
	var absoluteBase = httputil.NormalizeBase(config.Default().Frontend().AbsoluteBasePath())
	var fs http.FileSystem

	return module.Module{
		Name: "Frontend",
		Start: func(next module.NextFunc) (module.StopFunc, error) {
			var err error
			fs, err = initFs()
			if err != nil {
				return nil, err
			}

			return nil, next()
		},
		HTTP: module.HTTP{
			BasePath: base,
			Setup: func(r *mux.Router) {
				r.HandleFunc("/config", configHandler).Methods("GET")
				r.PathPrefix("/").Handler(fileHandler(absoluteBase, fs))
			},
		},
	}
}

func configHandler(res http.ResponseWriter, req *http.Request) {
	var apiURI = config.Default().API().AbsoluteBasePath()
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
