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

type DB struct {
	client  db.Client
	clients *clients.Store
	medias  *medias.Store
	plugins *plugins.Store
	users   *users.Store
}

func Open() (database *DB, err error) {
	database = new(DB)

	database.client, err = driver().Open()
	if err != nil {
		return nil, err
	}

	if database.clients, err = clients.CreateStore(database.client); err != nil {
		return nil, err
	}
	if database.medias, err = medias.CreateStore(database.client); err != nil {
		return nil, err
	}
	if database.plugins, err = plugins.CreateStore(database.client); err != nil {
		return nil, err
	}
	if database.users, err = users.CreateStore(database.client); err != nil {
		return nil, err
	}

	return
}

func (database *DB) Close() error {
	return database.client.Close()
}

func (database *DB) Begin() (*Tx, error) {
	tx, err := database.client.Begin()
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
