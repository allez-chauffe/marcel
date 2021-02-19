// +build !noembed

package frontend

import (
	"embed"
	"io/fs"
)

//go:embed build
var frontendFS embed.FS

func initFs() (fs.FS, error) {
	return fs.Sub(frontendFS, "build")
}
