package plugins

import (
	"github.com/allez-chauffe/marcel/api/db/internal/db"
)

var store db.Store

func CreateStore(database db.Databse) {
	store = database.CreateStore(func() db.Entity {
		return new(Plugin)
	})
}

func List() ([]Plugin, error) {
	var plugins = []Plugin{}
	return plugins, store.List(&plugins)
}

func Get(eltName string) (*Plugin, error) {
	var p = new(Plugin)
	return p, store.Get(eltName, p)
}

func Exists(eltName string) (bool, error) {
	return store.Exists(eltName)
}

func Insert(p *Plugin) error {
	return store.Insert(p)
}

func Update(p *Plugin) error {
	return store.Update(p)
}

func UpsertAll(plugins []Plugin) error {
	return db.Transactional(store, func(tx db.Transaction) error {
		for _, p := range plugins {
			if err := tx.Upsert(&p); err != nil {
				return err
			}
		}

		return nil
	})
}

func Delete(eltName string) error {
	return store.Delete(eltName)
}
