package db

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	bh "github.com/timshannon/bolthold"
	bolt "go.etcd.io/bbolt"

	"github.com/Zenika/marcel/api/db/internal/db"
	"github.com/Zenika/marcel/config"
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
	options.Timeout = 100 * time.Millisecond

	var err error
	if db.Store, err = bh.Open(os.ExpandEnv(config.Default().API().DBFile()), 0644, &bh.Options{
		Options: &options,
	}); err != nil {
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
