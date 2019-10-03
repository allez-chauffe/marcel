// +build !nopackr

package backoffice

import (
	"net/http"

	"github.com/gobuffalo/packr"
)

func initFs() (http.FileSystem, error) {
	return packr.NewBox("../backoffice/build/"), nil
}
