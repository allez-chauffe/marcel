package medias

import (
	"fmt"

	bh "github.com/timshannon/bolthold"
	bolt "go.etcd.io/bbolt"

	"github.com/Zenika/marcel/api/db/internal/db"
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

func nextID(tx *bolt.Tx) (int, error) {
	agg, err := db.Store.TxFindAggregate(tx, &Media{}, nil)
	if err != nil {
		return 0, err
	}

	// We find the new ID by adding 1 to the ID of the last created media
	if len(agg) != 0 && agg[0].Count() != 0 {
		lastMedia := new(Media)
		agg[0].Max("ID", lastMedia)
		return lastMedia.ID + 1, nil
	}

	return 0, nil
}

func Insert(m *Media) error {
	return db.Store.Bolt().Update(func(tx *bolt.Tx) error {
		ID, err := nextID(tx)
		if err != nil {
			return err
		}

		m.ID = ID
		m.Name = fmt.Sprintf("Media %d", m.ID)

		return db.Store.TxInsert(tx, m.ID, m)
	})
}

func Update(m *Media) error {
	return db.Store.Update(m.ID, m)
}

func Delete(id int) error {
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
