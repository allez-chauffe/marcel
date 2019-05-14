package backoffice

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/packr"
)

var Box = packr.NewBox("./build/")

func ListenAndServe(port uint, pBase string) error {
	base := pBase
	if base == "" {
		base = "/"
	}

	http.Handle("/", http.FileServer(Box))

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
