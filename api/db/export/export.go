package export

import (
	"encoding/json"
	"io"
	"os"

	"github.com/Zenika/marcel/api/db"
)

func export(fetch func() (interface{}, error), pretty bool, outputFile string) error {
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

	encoder := json.NewEncoder(w)
	if pretty {
		encoder.SetIndent("", "  ")
	}
	return encoder.Encode(data)
}
