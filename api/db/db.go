package db

import (
	"github.com/timshannon/bolthold"

	"github.com/Zenika/MARCEL/api/db/internal/db"
	"github.com/Zenika/MARCEL/api/db/users"
	"github.com/Zenika/MARCEL/config"
)

func Open() error {
	var err error
	db.Store, err = bolthold.Open(config.Config.DBFile, 0644, nil)
	if err != nil {
		return err
	}

	if err := users.EnsureOneUser(); err != nil {
		return err
	}

	// FIXME wait for signal and close ?

	return nil
}
