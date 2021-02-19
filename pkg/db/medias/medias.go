package medias

import (
	"fmt"

	"github.com/allez-chauffe/marcel/pkg/db/internal/db"
	"github.com/sirupsen/logrus"
)

var DefaultStore *Store

type Store struct {
	store db.Store
}

func Transactional(tx db.Transaction) *Store {
	return &Store{DefaultStore.store.Transactional(tx)}
}

func CreateStore() error {
	store, err := db.DB.CreateStore(func() db.Entity {
		return new(Media)
	})

	if err != nil {
		return err
	}

	DefaultStore = &Store{store}
	return nil
}

func (b *Store) List() ([]Media, error) {
	var medias = []Media{}
	return medias, b.store.List(&medias)
}

func (b *Store) Get(id int) (*Media, error) {
	var result = new(Media)
	return result, b.store.Get(id, &result)
}

func (b *Store) Insert(m *Media) (err error) {
	logrus.Debugf("Before ensure")
	return db.EnsureTransaction(b.store, func(store db.Store) error {
		if err = store.Insert(m); err != nil {
			return err
		}
		logrus.Debugf("After insert %v", m)

		// Set new media name if not given
		if m.Name == "" {
			m.Name = fmt.Sprintf("Media %d", m.ID)
			if err := store.Update(m); err != nil {
				return err
			}
		}

		return nil
	})
}

func (b *Store) Update(m *Media) error {
	return b.store.Update(m)
}

func (b *Store) Delete(id int) error {
	return b.store.Delete(id)
}

func (b *Store) UpsertAll(medias []Media) error {
	return db.EnsureTransaction(b.store, func(store db.Store) error {
		for _, m := range medias {
			if err := store.Upsert(&m); err != nil {
				return err
			}
		}

		return nil
	})
}

func (b *Store) Exists(id int) (bool, error) {
	return b.store.Exists(id)
}
