package backoffice

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gobuffalo/packr"

	"github.com/Zenika/marcel/config"
)

type BasedBox struct {
	packr.Box
	BasePath string
}

func (b BasedBox) Open(name string) (http.File, error) {
	return b.Box.Open(strings.TrimPrefix(name, strings.TrimSuffix(b.BasePath, "/")))
}

var Box = packr.NewBox("./build/")

func Start() error {
	base := "/test/"

	http.Handle(base, http.FileServer(BasedBox{Box, base}))

	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
}
