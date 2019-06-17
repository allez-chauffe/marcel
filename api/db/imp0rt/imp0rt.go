package imp0rt

import (
	"encoding/json"
	"io"
	"os"

	"github.com/Zenika/MARCEL/api/db"
)

func imp0rt(inputFile string, value interface{}, save func() error) error {
	if err := db.Open(); err != nil {
		return err
	}
	defer db.Close()

	var r io.ReadCloser
	if inputFile == "" {
		r = os.Stdin
	} else {
		var err error
		if r, err = os.Open(inputFile); err != nil {
			return err
		}
		defer r.Close()
	}

	if err := json.NewDecoder(r).Decode(value); err != nil {
		return err
	}

	return save()
}
