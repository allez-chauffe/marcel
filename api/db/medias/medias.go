package medias

import (
	bh "github.com/timshannon/bolthold"

	"github.com/Zenika/MARCEL/api/db/internal/db"
)

func List() ([]Media, error) {
	var medias = []Media{}

	return medias, db.Store.Find(&medias, nil)
}

func Get(id int) (*Media, error) {
	var m = new(Media)

	if err := db.Store.Get(id, m); err != nil {
		if err == bh.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return m, nil
}
