// +build nopackr

package backoffice

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/Zenika/marcel/config"
	"github.com/Zenika/marcel/version"
)

func initFs() (http.FileSystem, error) {
	url := fmt.Sprintf("https://github.com/Zenika/marcel/releases/download/%s/marcel-backoffice.tgz", version.Version)

	log.Infof("Fetching %s", url)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	path := filepath.Join(config.Config().API().DataDir(), "backoffice", version.Version)

	if err := os.RemoveAll(path); err != nil {
		return nil, err
	}

	if err := os.MkdirAll(path, os.ModePerm); err != nil { // FIXME is ModePerm OK ?
		return nil, err
	}

	log.Infoln("Uncompressing marcel-backoffice.tgz")

	gz, err := gzip.NewReader(res.Body)
	if err != nil {
		return nil, err
	}
	defer gz.Close()

	tr := tar.NewReader(gz)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		info := hdr.FileInfo()
		filePath := filepath.Join(path, hdr.Name)

		if info.IsDir() {
			if err := os.Mkdir(filePath, os.ModePerm); err != nil { // FIXME is ModePerm OK ?
				return nil, err
			}
		} else {
			if err := copyToFile(filePath, tr); err != nil {
				return nil, err
			}
		}
	}

	return http.Dir(path), nil
}

func copyToFile(name string, r io.Reader) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := io.Copy(f, r); err != nil {
		return err
	}

	return nil
}
