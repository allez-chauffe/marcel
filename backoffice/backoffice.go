package backoffice

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"

	"github.com/Zenika/marcel/config"
	"github.com/Zenika/marcel/httputil"
)

const index = "/index.html"

func Start() error {
	var r = mux.NewRouter()

	ConfigureRouter(r)

	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Backoffice.Port), r)
}

func ConfigureRouter(r *mux.Router) {
	var base = httputil.NormalizeBase(config.Config.Backoffice.BasePath)

	var b = r.PathPrefix(httputil.TrimTrailingSlash(base)).Subrouter()

	b.Handle("", http.RedirectHandler(base, http.StatusMovedPermanently))
	b.HandleFunc("/config", configHandler).Methods("GET")
	b.PathPrefix("/").Handler(fileHandler(base))
}

func configHandler(res http.ResponseWriter, req *http.Request) {
	var apiURI = config.Config.Backoffice.APIURI
	if !strings.HasSuffix(apiURI, "/") {
		apiURI += "/"
	}

	if err := json.NewEncoder(res).Encode(struct {
		APIURI      string `json:"apiURI"`
		FrontendURI string `json:"frontendURI"`
	}{
		APIURI:      apiURI,
		FrontendURI: config.Config.Backoffice.FrontendURI,
	}); err != nil {
		panic(err)
	}
}

func fileHandler(base string) http.Handler {
	return http.StripPrefix(
		base,
		http.FileServer(
			httputil.NewNotFoundRewriter(
				httputil.NewTemplater(
					packr.NewBox("./build/"),
					[]string{index},
					map[string]string{"REACT_APP_BASE_URL": base},
				),
				index,
			),
		),
	)
}
