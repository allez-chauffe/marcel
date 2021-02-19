package export

import (
	"encoding/json"
	"io"
	"os"

	"github.com/allez-chauffe/marcel/pkg/db"
)

func export(fetch func() (interface{}, error), outputFile string, pretty bool) error {
	if err := db.Open(); err != nil {
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
