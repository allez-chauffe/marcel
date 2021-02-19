// +build !noembed

package backoffice

import (
	"embed"
	"io/fs"
)

//go:embed build
var backofficeFS embed.FS

func initFs() (fs.FS, error) {
	return fs.Sub(backofficeFS, "build")
}
