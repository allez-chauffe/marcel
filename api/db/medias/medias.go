package medias

import (
	uuid "github.com/satori/go.uuid"
	bh "github.com/timshannon/bolthold"
	bolt "go.etcd.io/bbolt"

	"github.com/Zenika/marcel/api/db/internal/db"
)

func List() ([]Media, error) {
	var medias = []Media{}

	return medias, db.Store.Find(&medias, nil)
}

func Get(id string) (*Media, error) {
	var m = new(Media)

	if err := db.Store.Get(id, m); err != nil {
		if err == bh.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return m, nil
}

func Insert(m *Media) error {
	m.ID = uuid.NewV4().String()
	return db.Store.Insert(m.ID, m)
}

func Update(m *Media) error {
	return db.Store.Update(m.ID, m)
}

func Delete(id string) error {
	return db.Store.Delete(id, &Media{})
}

func UpsertAll(medias []Media) error {
	return db.Store.Bolt().Update(func(tx *bolt.Tx) error {
		for _, m := range medias {
			if err := db.Store.TxUpsert(tx, m.ID, &m); err != nil {
				return err
			}
		}

		return nil
	})
}
