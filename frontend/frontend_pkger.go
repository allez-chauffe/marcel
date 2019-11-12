// +build !nopkger

package frontend

import (
	"net/http"

	"github.com/markbates/pkger"

	// Imports pkger version of frontend
	_ "github.com/Zenika/marcel/pkged"
)

func initFs() (http.FileSystem, error) {
	return pkger.Dir("/frontend/build"), nil
}
