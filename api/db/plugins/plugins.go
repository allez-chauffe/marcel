package plugins

import (
	"github.com/allez-chauffe/marcel/api/db/internal/db"
)

var DefaultStore *Store

type Store struct {
	store db.Store
}

func CreateStore() error {
	store, err := db.DB.CreateStore(func() db.Entity {
		return new(Plugin)
	})

	if err != nil {
		return err
	}

	DefaultStore = &Store{store}
	return nil
}

func Transactional(tx db.Transaction) *Store {
	return &Store{DefaultStore.store.Transactional(tx)}
}

func (b *Store) List() ([]Plugin, error) {
	var plugins = []Plugin{}
	return plugins, b.store.List(&plugins)
}

func (b *Store) Get(eltName string) (*Plugin, error) {
	var p = new(Plugin)
	return p, b.store.Get(eltName, &p)
}

func (b *Store) Exists(eltName string) (bool, error) {
	return b.store.Exists(eltName)
}

func (b *Store) Insert(p *Plugin) error {
	return b.store.Insert(p)
}

func (b *Store) Update(p *Plugin) error {
	return b.store.Update(p)
}

func (b *Store) UpsertAll(plugins []Plugin) error {
	return db.EnsureTransaction(b.store, func(store db.Store) error {
		for _, p := range plugins {
			if err := store.Upsert(&p); err != nil {
				return err
			}
		}

		return nil
	})
}

func (b *Store) Delete(eltName string) error {
	return b.store.Delete(eltName)
}
