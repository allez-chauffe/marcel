package httputil

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func DownloadTgzToDir(url, path string) error {
	log.Infof("Fetching %s...", url)
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("GET %s returned status %d", url, res.StatusCode)
	}

	if err := os.RemoveAll(path); err != nil {
		return err
	}

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}

	log.Infof("Uncompressing %s...", url)

	gz, err := gzip.NewReader(res.Body)
	if err != nil {
		return err
	}
	defer gz.Close()

	tr := tar.NewReader(gz)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		info := hdr.FileInfo()
		filePath := filepath.Join(path, hdr.Name)

		if info.IsDir() {
			if err := os.Mkdir(filePath, os.ModePerm); err != nil {
				return err
			}
		} else {
			if err := copyToFile(filePath, tr); err != nil {
				return err
			}
		}
	}

	return nil
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
