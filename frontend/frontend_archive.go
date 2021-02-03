// +build noembed

package frontend

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/allez-chauffe/marcel/config"
	"github.com/allez-chauffe/marcel/httputil"
	"github.com/allez-chauffe/marcel/version"
)

func initFs() (fs.FS, error) {
	url := fmt.Sprintf("https://github.com/allez-chauffe/marcel/releases/download/%[1]s/marcel-frontend-%[1]s.tgz", version.Version())
	path := filepath.Join(config.Default().API().DataDir(), "frontend", version.Version())

	if err := httputil.DownloadTgzToDir(url, path); err != nil {
		return nil, err
	}

	return os.DirFS(path), nil
}
