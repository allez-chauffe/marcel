package frontend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/Zenika/marcel/config"
	"github.com/Zenika/marcel/httputil"
)

func Start() error {
	var r = mux.NewRouter()

	ConfigureRouter(r)

	log.Infof("Starting frontend server on port %d...", config.Config.Frontend.Port)

	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Frontend.Port), r)
}

func ConfigureRouter(r *mux.Router) {
	var base = httputil.NormalizeBase(config.Config.Frontend.BasePath)

	var b = r.PathPrefix(httputil.TrimTrailingSlash(base)).Subrouter()

	b.HandleFunc("/config", configHandler).Methods("GET")
	b.PathPrefix("/").Handler(fileHandler(base))
}

func configHandler(res http.ResponseWriter, req *http.Request) {
	var apiURI = config.Config.Frontend.APIURI
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

func fileHandler(base string) http.Handler {
	return http.StripPrefix(
		base,
		http.FileServer(
			packr.NewBox("./build/"),
		),
	)
}
