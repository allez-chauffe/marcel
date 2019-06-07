package export

import (
	"encoding/json"
	"os"

	"github.com/Zenika/MARCEL/api/db"
)

func export(fetch func() (interface{}, error), outputFile string) error {
	if err := db.OpenRO(); err != nil {
		return err
	}
	defer db.Close()

	f, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := fetch()
	if err != nil {
		return err
	}

	return json.NewEncoder(f).Encode(data)
}
