package backoffice

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/Zenika/marcel/config"
	"github.com/Zenika/marcel/httputil"
	"github.com/Zenika/marcel/module"
)

const index = "/index.html"

func Module() module.Module {
	var base = httputil.NormalizeBase(config.Default().Backoffice().BasePath())
	var fs http.FileSystem

	return module.Module{
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
			BasePath: httputil.TrimTrailingSlash(base),
			Setup: func(r *mux.Router) {
				r.Handle("", http.RedirectHandler(base, http.StatusMovedPermanently))
				r.HandleFunc("/config", configHandler).Methods("GET")
				r.PathPrefix("/").Handler(fileHandler(base, fs))
			},
		},
	}
}

func configHandler(res http.ResponseWriter, req *http.Request) {
	// FIXME utility ?
	var apiURI = config.Default().Backoffice().APIURI()
	if !strings.HasSuffix(apiURI, "/") {
		apiURI += "/"
	}

	var frontendURI = config.Default().Backoffice().FrontendURI()
	if !strings.HasSuffix(frontendURI, "/") {
		frontendURI += "/"
	}

	if err := json.NewEncoder(res).Encode(struct {
		APIURI      string `json:"apiURI"`
		FrontendURI string `json:"frontendURI"`
	}{
		APIURI:      apiURI,
		FrontendURI: frontendURI,
	}); err != nil {
		panic(err)
	}
}

func fileHandler(base string, fs http.FileSystem) http.Handler {
	return http.StripPrefix(
		base,
		http.FileServer(
			httputil.NewNotFoundRewriter(
				httputil.NewTemplater(
					fs,
					[]string{index},
					map[string]string{"REACT_APP_BASE": base},
				),
				index,
			),
		),
	)
}
