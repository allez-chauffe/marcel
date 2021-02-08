package db

import (
	"fmt"

	"github.com/allez-chauffe/marcel/api/db/clients"
	"github.com/allez-chauffe/marcel/api/db/drivers/bolt"
	"github.com/allez-chauffe/marcel/api/db/drivers/postgres"
	"github.com/allez-chauffe/marcel/api/db/internal/db"
	"github.com/allez-chauffe/marcel/api/db/medias"
	"github.com/allez-chauffe/marcel/api/db/plugins"
	"github.com/allez-chauffe/marcel/api/db/users"
	"github.com/allez-chauffe/marcel/config"
)

func Open() error {
	database, err := driver().Open()
	if err != nil {
		return err
	}

	db.DB = database

	// Initialise every stores
	if err := clients.CreateDefaultBucket(); err != nil {
		return err
	}
	if err := medias.CreateDefaultBucket(); err != nil {
		return err
	}
	if err := plugins.CreateDefaultBucket(); err != nil {
		return err
	}
	if err := users.CreateDefaultBucket(); err != nil {
		return err
	}

	return nil
}

func Close() error {
	return db.DB.Close()
}

func Begin() (*Tx, error) {
	tx, err := db.DB.Begin()
	if err != nil {
		return nil, err
	}

	return &Tx{tx}, nil
}

func Transactional(task func(*Tx) error) (err error) {
	return db.Transactional(func(tx db.Transaction) error {
		return task(&Tx{tx})
	})
}

func driver() db.Driver {
	switch config.Default().API().DB().Driver() {
	case "bolt":
		return bolt.Driver
	case "postgres":
		return postgres.Driver
	default:
		panic(fmt.Errorf("Unknown database driver %s", driver))
	}
}
