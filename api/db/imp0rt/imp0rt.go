package imp0rt

import (
	"encoding/json"
	"os"

	"github.com/Zenika/MARCEL/api/db"
)

func imp0rt(inputFile string, value interface{}, save func() error) error {
	if err := db.Open(); err != nil {
		return err
	}
	defer db.Close()

	f, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(value); err != nil {
		return err
	}

	return save()
}
