package backoffice

import (
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/allez-chauffe/marcel/config"
	"github.com/allez-chauffe/marcel/module"
	xfs "github.com/allez-chauffe/marcel/pkg/io/fs"
)

const index = "index.html"

// Module creates backoffice module
func Module() *module.Module {
	var fs fs.FS

	// Set default URIs for API and Frontend
	// in case Backoffice is the root module
	module.SetURI("API", config.Default().API().BasePath())
	module.SetURI("Frontend", config.Default().Frontend().BasePath())

	return &module.Module{
		Name: "Backoffice",
		Start: func(_ module.Context, next module.NextFunc) (module.StopFunc, error) {
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
			Setup: func(_ module.Context, basePath string, r *mux.Router) {
				r.PathPrefix("/").Handler(fileHandler(basePath, fs))
			},
		},
	}
}

func fileHandler(basePath string, fs fs.FS) http.Handler {
	return http.StripPrefix(
		basePath,
		http.FileServer(http.FS(
			xfs.NewCatchAll(
				xfs.NewTemplater(
					fs,
					[]string{index},
					map[string]string{"REACT_APP_BASE": basePath},
				),
				index,
			),
		)),
	)
}
