// +build !nopkger

package backoffice

import (
	"net/http"

	"github.com/markbates/pkger"

	// Imports pkger version of backoffice
	_ "github.com/Zenika/marcel/pkged"
)

func initFs() (http.FileSystem, error) {
	return pkger.Dir("/backoffice/build"), nil
}
