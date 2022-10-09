package api

import (
	"net/http"
	log "github.com/sirupsen/logrus"
)


func LoggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	    log.Infof("%s request %s", r.Method, r.RequestURI)
		h.ServeHTTP(w, r)
	})
}