// +build noembed

package backoffice

import (
	"fmt"
	"io/fs"

	"github.com/allez-chauffe/marcel/httputil"
	"github.com/allez-chauffe/marcel/version"
)

func initFs() (fs.FS, error) {
	return httputil.GetTgzFS(fmt.Sprintf("https://github.com/allez-chauffe/marcel/releases/download/%[1]s/marcel-backoffice-%[1]s.tgz", version.Version()))
}
