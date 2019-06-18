package backoffice

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"

	"github.com/Zenika/marcel/config"
)

var (
	fileHandler = http.FileServer(&notFoundToIndex{packr.NewBox("./build/")})

	backofficeConfig = struct {
		BackendURI  string `json:"backendURI"`
		FrontendURI string `json:"frontendURI"`
	}{
		BackendURI:  "/api/v1/",
		FrontendURI: "/front/",
	}
)

func Start(base string) error {
	base = strings.TrimSuffix(base, "/")

	r := mux.NewRouter()

	baseRoute := r.PathPrefix(base)
	baseRouter := baseRoute.Subrouter()

	baseRouter.Handle("", http.RedirectHandler(base+"/", http.StatusMovedPermanently))
	baseRouter.HandleFunc("/config", configHandler).Methods("GET")
	baseRouter.PathPrefix("/").Handler(http.StripPrefix(base, fileHandler))

	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), r)
}

func configHandler(res http.ResponseWriter, req *http.Request) {
	if err := json.NewEncoder(res).Encode(backofficeConfig); err != nil {
		panic(err)
	}
}

type notFoundToIndex struct {
	http.FileSystem
}

var _ http.FileSystem = (*notFoundToIndex)(nil)

func (fs *notFoundToIndex) Open(path string) (http.File, error) {
	f, err := fs.FileSystem.Open(path)
	if err != nil && os.IsNotExist(err) {
		return fs.FileSystem.Open("/index.html")
	}
	return f, err
}
