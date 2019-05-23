package db

import (
	log "github.com/sirupsen/logrus"
	"github.com/timshannon/bolthold"

	"github.com/Zenika/MARCEL/api/db/internal/db"
	"github.com/Zenika/MARCEL/api/db/users"
	"github.com/Zenika/MARCEL/config"
)

func Open() error {
	var err error
	if db.Store, err = bolthold.Open(config.Config.DBFile, 0644, nil); err != nil {
		return err
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
