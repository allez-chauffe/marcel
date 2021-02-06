package plugins

import (
	"github.com/allez-chauffe/marcel/api/db/internal/db"
)

var DefaultBucket *Bucket

type Bucket struct {
	store db.Store
}

func CreateDefaultBucket() {
	DefaultBucket = &Bucket{
		db.DB.CreateStore(func() db.Entity {
			return new(Plugin)
		}),
	}
}

func Transactional(tx db.Transaction) *Bucket {
	return &Bucket{DefaultBucket.store.Transactional(tx)}
}

func (b *Bucket) List() ([]Plugin, error) {
	var plugins = []Plugin{}
	return plugins, b.store.List(&plugins)
}

func (b *Bucket) Get(eltName string) (*Plugin, error) {
	var p = new(Plugin)
	return p, b.store.Get(eltName, &p)
}

func (b *Bucket) Exists(eltName string) (bool, error) {
	return b.store.Exists(eltName)
}

func (b *Bucket) Insert(p *Plugin) error {
	return b.store.Insert(p)
}

func (b *Bucket) Update(p *Plugin) error {
	return b.store.Update(p)
}

func (b *Bucket) UpsertAll(plugins []Plugin) error {
	return db.EnsureTransaction(b.store, func(store db.Store) error {
		for _, p := range plugins {
			if err := store.Upsert(&p); err != nil {
				return err
			}
		}

		return nil
	})
}

func (b *Bucket) Delete(eltName string) error {
	return b.store.Delete(eltName)
}
