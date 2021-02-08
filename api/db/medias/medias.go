package medias

import (
	"fmt"

	"github.com/allez-chauffe/marcel/api/db/internal/db"
)

var DefaultBucket *Bucket

type Bucket struct {
	store db.Store
}

func Transactional(tx db.Transaction) *Bucket {
	return &Bucket{DefaultBucket.store.Transactional(tx)}
}

func CreateDefaultBucket() error {
	store, err := db.DB.CreateStore(func() db.Entity {
		return new(Media)
	})

	if err != nil {
		return err
	}

	DefaultBucket = &Bucket{store}
	return nil
}

func (b *Bucket) List() ([]Media, error) {
	var medias = []Media{}
	return medias, b.store.List(&medias)
}

func (b *Bucket) Get(id int) (*Media, error) {
	var result = new(Media)
	return result, b.store.Get(id, &result)
}

func (b *Bucket) Insert(m *Media) (err error) {
	return db.EnsureTransaction(b.store, func(store db.Store) error {
		if err = store.Insert(m); err != nil {
			return err
		}

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

func (b *Bucket) Update(m *Media) error {
	return b.store.Update(m)
}

func (b *Bucket) Delete(id int) error {
	return b.store.Delete(id)
}

func (b *Bucket) UpsertAll(medias []Media) error {
	return db.EnsureTransaction(b.store, func(store db.Store) error {
		for _, m := range medias {
			if err := store.Upsert(&m); err != nil {
				return err
			}
		}

		return nil
	})
}

func (b *Bucket) Exists(id int) (bool, error) {
	return b.store.Exists(id)
}
