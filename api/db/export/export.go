package export

import (
	"encoding/json"
	"io"
	"os"

	"github.com/Zenika/marcel/api/db"
)

func export(fetch func() (interface{}, error), outputFile string) error {
	if err := db.OpenRO(); err != nil {
		return err
	}
	defer db.Close()

	var w io.WriteCloser
	if outputFile == "" {
		w = os.Stdout
	} else {
		var err error
		if w, err = os.Create(outputFile); err != nil {
			return err
		}
		defer w.Close()
	}

	data, err := fetch()
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(data)
}
