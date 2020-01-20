package backoffice

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Zenika/marcel/config"
	"github.com/Zenika/marcel/httputil"
	"github.com/Zenika/marcel/module"
)

const index = "/index.html"

// Module creates backoffice module
func Module() *module.Module {
	var fs http.FileSystem

	return &module.Module{
		Name: "Backoffice",
		Start: func(next module.NextFunc) (module.StopFunc, error) {
			var err error
			fs, err = initFs()
			if err != nil {
				return nil, err
			}

			return nil, next()
		},
		HTTP: module.HTTP{
			BasePath:      config.Default().Backoffice().BasePath(),
			RedirectSlash: true,
			Setup: func(basePath string, r *mux.Router) {
				r.PathPrefix("/").Handler(fileHandler(basePath, fs))
			},
		},
	}
}

func fileHandler(basePath string, fs http.FileSystem) http.Handler {
	return http.StripPrefix(
		basePath,
		http.FileServer(
			httputil.NewNotFoundRewriter(
				httputil.NewTemplater(
					fs,
					[]string{index},
					map[string]string{"REACT_APP_BASE": basePath},
				),
				index,
			),
		),
	)
}
