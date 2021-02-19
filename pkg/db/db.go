package db

import (
	"fmt"

	"github.com/allez-chauffe/marcel/pkg/config"
	"github.com/allez-chauffe/marcel/pkg/db/clients"
	"github.com/allez-chauffe/marcel/pkg/db/drivers/bolt"
	"github.com/allez-chauffe/marcel/pkg/db/drivers/postgres"
	"github.com/allez-chauffe/marcel/pkg/db/internal/db"
	"github.com/allez-chauffe/marcel/pkg/db/medias"
	"github.com/allez-chauffe/marcel/pkg/db/plugins"
	"github.com/allez-chauffe/marcel/pkg/db/users"
)

func Open() error {
	database, err := driver().Open()
	if err != nil {
		return err
	}

	db.DB = database

	// Initialize every stores
	if err := clients.CreateStore(); err != nil {
		return err
	}
	if err := medias.CreateStore(); err != nil {
		return err
	}
	if err := plugins.CreateStore(); err != nil {
		return err
	}
	if err := users.CreateStore(); err != nil {
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
		panic(fmt.Errorf("Unknown database driver %s", config.Default().API().DB().Driver()))
	}
}
