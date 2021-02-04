package medias

import (
	"fmt"

	"github.com/allez-chauffe/marcel/api/db/internal/db"
)

var store db.Store

func CreateStore(database db.Databse) {
	store = database.CreateStore(func() db.Entity{
		return new(Media)
	})
}

func List() ([]Media, error) {
	var medias = []Media{}
	return medias, store.List(&medias)
}

func Get(id int) (*Media, error) {
	var result = new(Media)
	return result, store.Get(id, &result)
}

func Insert(m *Media) error {
	return db.Transactional(store, func(tx db.Transaction) error {
		m.ID = -1 // Let auto increment set the id
		if err := tx.Insert(m); err != nil {
			return err
		}

		// Set new media name with inserted ID
		m.Name = fmt.Sprintf("Media %d", m.ID)
		if err := tx.Update(m); err != nil {
			return err
		}

		return nil
	})
}

func Update(m *Media) error {
	return store.Update(m)
}

func Delete(id int) error {
	return store.Delete(id)
}

func UpsertAll(medias []Media) error {
	return db.Transactional(store, func(tx db.Transaction) error {
		for _, m := range medias {
			if err := tx.Upsert(&m); err != nil {
				return err
			}
		}

		return nil
	})
}
