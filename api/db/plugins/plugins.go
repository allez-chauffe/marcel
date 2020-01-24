package plugins

import (
	uuid "github.com/satori/go.uuid"
	bh "github.com/timshannon/bolthold"
	bolt "go.etcd.io/bbolt"

	"github.com/Zenika/marcel/api/db/internal/db"
)

func List() ([]Plugin, error) {
	var plugins = []Plugin{}

	return plugins, db.Store.Find(&plugins, nil)
}

func Get(id string) (*Plugin, error) {
	var p = new(Plugin)

	if err := db.Store.Get(id, p); err != nil {
		if err == bh.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return p, nil
}

func GetByPath(path string) (*Plugin, error) {
	var plugins []Plugin

	if err := db.Store.Find(&plugins, bh.Where("Path").Eq(path).Index("Path")); err != nil {
		return nil, err
	}

	if len(plugins) == 0 {
		return nil, nil
	}

	return &plugins[0], nil
}

func Insert(p *Plugin) error {
	return db.Store.Insert(p.Path, p)
}

func Update(p *Plugin) error {
	return db.Store.Update(p.ID, p)
}

func UpsertAll(plugins []Plugin) error {
	return db.Store.Bolt().Update(func(tx *bolt.Tx) error {
		for _, p := range plugins {
			if p.ID == "" {
				p.ID = uuid.NewV4().String()
			}
			if err := db.Store.TxUpsert(tx, p.ID, &p); err != nil {
				return err
			}
		}

		return nil
	})
}

func Delete(id string) error {
	return db.Store.Delete(id, &Plugin{})
}
