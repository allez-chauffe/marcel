// +build nopkger

package frontend

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/Zenika/marcel/config"
	"github.com/Zenika/marcel/httputil"
	"github.com/Zenika/marcel/version"
)

func initFs() (http.FileSystem, error) {
	url := fmt.Sprintf("https://github.com/Zenika/marcel/releases/download/%[1]s/marcel-frontend-%[1]s.tgz", version.Version())
	path := filepath.Join(config.Default().API().DataDir(), "frontend", version.Version())

	if err := httputil.DownloadTgzToDir(url, path); err != nil {
		return nil, err
	}

	return http.Dir(path), nil
}
