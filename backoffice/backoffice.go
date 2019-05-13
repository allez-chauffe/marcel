package backoffice

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/packr/v2"
)

var Box = packr.New("backoffice", "./build/")

func ListenAndServe(port uint, pBase string) error {
	base := pBase
	if base == "" {
		base = "/"
	}

	http.Handle("/", http.FileServer(Box))

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
