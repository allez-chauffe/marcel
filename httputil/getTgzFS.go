package httputil

import (
	"compress/gzip"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/nlepage/go-tarfs"
	log "github.com/sirupsen/logrus"
)

func GetTgzFS(url string) (fs.FS, error) {
	log.Infof("Fetching %s...", url)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %s returned status %d", url, res.StatusCode)
	}

	log.Infof("Uncompressing %s...", url)

	gz, err := gzip.NewReader(res.Body)
	if err != nil {
		return nil, err
	}
	defer gz.Close()

	return tarfs.New(gz)
}
