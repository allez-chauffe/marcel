package frontend

import (
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/allez-chauffe/marcel/pkg/config"
	xfs "github.com/allez-chauffe/marcel/pkg/io/fs"
	"github.com/allez-chauffe/marcel/pkg/module"
)

// Module creates the frontend module
func Module() *module.Module {
	var fs fs.FS

	// Set default URIs for API in case Frontend is the root module
	module.SetURI("API", config.Default().API().BasePath())

	return &module.Module{
		Name: "Frontend",
		Start: func(_ module.Context, next module.NextFunc) (module.StopFunc, error) {
			var err error
			fs, err = initFs()
			if err != nil {
				return nil, err
			}

			return nil, next()
		},
		HTTP: module.HTTP{
			BasePath: config.Default().Frontend().BasePath(),
			Setup: func(_ module.Context, basePath string, r *mux.Router) {
				r.PathPrefix("/").Handler(fileHandler(basePath, fs))
			},
		},
	}
}

func fileHandler(base string, fs fs.FS) http.Handler {
	return http.StripPrefix(
		base,
		http.FileServer(http.FS(
			xfs.NewTemplater(
				fs,
				[]string{"index.html"},
				map[string]string{"REACT_APP_BASE": base},
			),
		)),
	)
}
