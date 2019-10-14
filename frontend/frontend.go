package frontend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/Zenika/marcel/config"
	"github.com/Zenika/marcel/httputil"
)

func Start() error {
	var r = mux.NewRouter()

	ConfigureRouter(r)

	log.Infof("Frontend server listening on %d...", config.Default().Frontend().Port())

	return http.ListenAndServe(fmt.Sprintf(":%d", config.Default().Frontend().Port()), r)
}

func ConfigureRouter(r *mux.Router) error {
	var base = httputil.NormalizeBase(config.Default().Frontend().BasePath())

	var b = r.PathPrefix(httputil.TrimTrailingSlash(base)).Subrouter()

	b.HandleFunc("/config", configHandler).Methods("GET")
	fh, err := fileHandler(base)
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

func fileHandler(base string) (http.Handler, error) {
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
