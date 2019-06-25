package backoffice

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"

	"github.com/Zenika/marcel/config"
	"github.com/Zenika/marcel/httputil"
)

const index = "/index.html"

var (
	backofficeConfig = struct {
		BackendURI  string `json:"backendURI"`
		FrontendURI string `json:"frontendURI"`
	}{
		BackendURI:  "/api/v1/",
		FrontendURI: "/front/",
	}
)

func Start(base string) error {
	base = httputil.NormalizeBase(base)

	r := mux.NewRouter()

	ConfigureRouter(r.PathPrefix(httputil.TrimTrailingSlash(base)).Subrouter(), base)

	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.API.Port), r)
}

func ConfigureRouter(r *mux.Router, base string) {
	base = httputil.NormalizeBase(base)

	r.Handle("", http.RedirectHandler(base, http.StatusMovedPermanently))
	r.HandleFunc("/config", configHandler).Methods("GET")
	r.PathPrefix("/").Handler(fileHandler(base))
}

func configHandler(res http.ResponseWriter, req *http.Request) {
	if err := json.NewEncoder(res).Encode(backofficeConfig); err != nil {
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
