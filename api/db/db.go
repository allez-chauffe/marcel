package db

import (
	bolt "go.etcd.io/bbolt"

	"github.com/Zenika/MARCEL/api/db/internal/db"
	"github.com/Zenika/MARCEL/config"
)

func Open() error {
	var err error
	db.DB, err = bolt.Open(config.Config.DBFile, 0644, nil)
	if err != nil {
		return err
	}

	// FIXME wait for signal and close

	return nil
}
