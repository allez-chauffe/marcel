// +build noembed

package backoffice

import (
	"fmt"
	"io/fs"

	xfs "github.com/allez-chauffe/marcel/pkg/io/fs"
	"github.com/allez-chauffe/marcel/pkg/version"
)

func initFs() (fs.FS, error) {
	return xfs.NewHTTPTgz(fmt.Sprintf("https://github.com/allez-chauffe/marcel/releases/download/%[1]s/marcel-backoffice-%[1]s.tgz", version.Version()))
}
