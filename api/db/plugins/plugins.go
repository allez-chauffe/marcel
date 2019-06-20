package plugins

import (
	bh "github.com/timshannon/bolthold"
	bolt "go.etcd.io/bbolt"

	"github.com/Zenika/marcel/api/db/internal/db"
)

func List() ([]Plugin, error) {
	var plugins = []Plugin{}

	return plugins, db.Store.Find(&plugins, nil)
}

func Get(eltName string) (*Plugin, error) {
	var p = new(Plugin)

	if err := db.Store.Get(eltName, p); err != nil {
		if err == bh.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return p, nil
}

func Exists(eltName string) (bool, error) {
	if err := db.Store.Get(eltName, &Plugin{}); err != nil {
		if err == bh.ErrNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func Insert(p *Plugin) error {
	return db.Store.Insert(p.EltName, p)
}

func Update(p *Plugin) error {
	return db.Store.Update(p.EltName, p)
}

func UpsertAll(plugins []Plugin) error {
	return db.Store.Bolt().Update(func(tx *bolt.Tx) error {
		for _, p := range plugins {
			if err := db.Store.TxUpsert(tx, p.EltName, &p); err != nil {
				return err
			}
		}

		return nil
	})
}

func Delete(eltName string) error {
	return db.Store.Delete(eltName, &Plugin{})
}
