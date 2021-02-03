package db

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	bh "github.com/timshannon/bolthold"
	bolt "go.etcd.io/bbolt"

	"github.com/allez-chauffe/marcel/api/db/internal/db"
	"github.com/allez-chauffe/marcel/config"
)

// Open opens bbolt database in read/write mode
func Open() error {
	return open(false)
}

// OpenRO opens bbolt database in read only mode
func OpenRO() error {
	return open(true)
}

func open(readOnly bool) error {
	log.Info("Opening bbolt database...")

	var options = *bolt.DefaultOptions
	options.ReadOnly = readOnly
	options.Timeout = 100 * time.Millisecond

	var err error
	if db.Store, err = bh.Open(os.ExpandEnv(config.Default().API().DBFile()), 0644, &bh.Options{
		Options: &options,
	}); err != nil {
		return fmt.Errorf("Error while opening bbolt database: %w", err)
	}

	log.Info("bbolt database opened")

	return nil
}

// Close closes bbolt database connection
func Close() error {
	log.Info("Closing bbolt database...")

	err := db.Store.Close()
	if err != nil {
		return fmt.Errorf("Error while closing bbolt database: %w", err)
	}

	log.Info("bbolt database closed")

	return nil
}
