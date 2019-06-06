package db

import (
	log "github.com/sirupsen/logrus"
	bh "github.com/timshannon/bolthold"
	bolt "go.etcd.io/bbolt"

	"github.com/Zenika/MARCEL/api/db/internal/db"
	"github.com/Zenika/MARCEL/api/db/users"
	"github.com/Zenika/MARCEL/config"
)

func Open() error {
	return open(false)
}

func OpenRO() error {
	return open(true)
}

func open(readOnly bool) error {
	var options = *bolt.DefaultOptions
	options.ReadOnly = readOnly

	var err error
	if db.Store, err = bh.Open(config.Config.DBFile, 0644, &bh.Options{
		Options: &options,
	}); err != nil {
		return err
	}

	if readOnly {
		return nil
	}

	if err := users.EnsureOneUser(); err != nil {
		return err
	}

	return nil
}

func Close() error {
	log.Info("Closing database...")

	err := db.Store.Close()
	if err != nil {
		return err
	}

	log.Info("Database closed")

	return nil
}
