// +build !nopackr

package frontend

import (
	"net/http"

	"github.com/gobuffalo/packr"
)

func initFs() (http.FileSystem, error) {
	return packr.NewBox("../frontend/build/"), nil
}
