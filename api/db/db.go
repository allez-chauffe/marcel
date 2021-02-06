package db

import (
	"github.com/allez-chauffe/marcel/api/db/clients"
	bhDriver "github.com/allez-chauffe/marcel/api/db/drivers/bolthold"
	"github.com/allez-chauffe/marcel/api/db/internal/db"
	"github.com/allez-chauffe/marcel/api/db/medias"
	"github.com/allez-chauffe/marcel/api/db/plugins"
	"github.com/allez-chauffe/marcel/api/db/users"
)

func Open() error {
	return open(false)
}

func OpenRO() error {
	return open(true)
}

func open(readonly bool) error {
	db.DB = getDatabaseDriver()

	if err := db.DB.Open(readonly); err != nil {
		return err
	}

	// Initialise every stores
	clients.CreateDefaultBucket()
	medias.CreateDefaultBucket()
	plugins.CreateDefaultBucket()
	users.CreateDefaultBucket()

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

func getDatabaseDriver() db.Database {
	// TODO: Select the driver in config
	return bhDriver.New()
}
