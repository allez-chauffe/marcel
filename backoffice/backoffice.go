package backoffice

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

const index = "/index.html"

func Start() error {
	var r = mux.NewRouter()

	ConfigureRouter(r)

	log.Infof("Backoffice server listening on %d...", config.Default().Backoffice().Port())

	return http.ListenAndServe(fmt.Sprintf(":%d", config.Default().Backoffice().Port()), r)
}

func ConfigureRouter(r *mux.Router) error {
	var base = httputil.NormalizeBase(config.Default().Backoffice().BasePath())

	var b = r.PathPrefix(httputil.TrimTrailingSlash(base)).Subrouter()

	b.Handle("", http.RedirectHandler(base, http.StatusMovedPermanently))
	b.HandleFunc("/config", configHandler).Methods("GET")
	fh, err := fileHandler(base)
	if err != nil {
		return err
	}
	b.PathPrefix("/").Handler(fh)

	return nil
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

func fileHandler(base string) (http.Handler, error) {
	fs, err := initFs()
	if err != nil {
		return nil, err
	}
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
	), nil
}
