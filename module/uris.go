package module

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var uris = make(map[string]string)

// URI gives the URI of a module.
func URI(name string) string {
	return uris[name]
}

// SetURI sets the URI of a module (for distant modules).
func SetURI(name, basePath string) {
	uris[name] = basePath
}

func urisHandler(res http.ResponseWriter, req *http.Request) {
	if err := json.NewEncoder(res).Encode(uris); err != nil {
		log.WithError(err).Error("Could not JSON encode URIS")
		res.WriteHeader(http.StatusInternalServerError)
	}
}
