package db

import (
	"github.com/allez-chauffe/marcel/api/db/clients"
	bhDriver "github.com/allez-chauffe/marcel/api/db/drivers/bolthold"
	"github.com/allez-chauffe/marcel/api/db/internal/db"
	"github.com/allez-chauffe/marcel/api/db/medias"
	"github.com/allez-chauffe/marcel/api/db/plugins"
	"github.com/allez-chauffe/marcel/api/db/users"
)

var DB db.Databse

func Open() error {
	return open(false)
}

func OpenRO() error {
	return open(true)
}

func open(readonly bool) error {
	DB = getDatabaseDriver()

	if err := DB.Open(readonly); err != nil {
		return err
	}

	// Initialise every stores
	clients.CreateStore(DB)
	medias.CreateStore(DB)
	plugins.CreateStore(DB)
	users.CreateStore(DB)

	return nil
}

func Close() error {
	return DB.Close()
}

func getDatabaseDriver() db.Databse {
	// TODO: Select the driver in config
	return bhDriver.New()
}