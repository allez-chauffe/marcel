package frontend

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Zenika/marcel/config"
	"github.com/Zenika/marcel/httputil"
	"github.com/Zenika/marcel/module"
)

// Module creates the frontend module
func Module() *module.Module {
	var fs http.FileSystem

	// Set default URIs for API in case Frontend is the root module
	module.SetURI("API", config.Default().API().BasePath())

	return &module.Module{
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
			BasePath: config.Default().Frontend().BasePath(),
			Setup: func(basePath string, r *mux.Router) {
				r.PathPrefix("/").Handler(fileHandler(basePath, fs))
			},
		},
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
