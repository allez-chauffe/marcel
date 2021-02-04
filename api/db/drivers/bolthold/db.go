package bolt

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	bh "github.com/timshannon/bolthold"
	bolt "go.etcd.io/bbolt"

	"github.com/Zenika/marcel/api/db/internal/db"
	"github.com/Zenika/marcel/config"
)

type boltDatabase struct {
	bh *bh.Store
}

func New() db.Databse {
	return &boltDatabase{
		new(bh.Store),
	}
}

func (database *boltDatabase) Open(readOnly bool) error {
	log.Info("Opening bbolt database...")

	var options = *bolt.DefaultOptions
	options.ReadOnly = readOnly
	options.Timeout = 100 * time.Millisecond

	var err error
	if database.bh, err = bh.Open(os.ExpandEnv(config.Default().API().DBFile()), 0644, &bh.Options{
		Options: &options,
	}); err != nil {
		return fmt.Errorf("Error while opening bbolt database: %w", err)
	}

	log.Info("bbolt database opened")

	return nil
}

func (database *boltDatabase) Close() error {
	log.Info("Closing bbolt database...")

	err := database.bh.Close()
	if err != nil {
		return fmt.Errorf("Error while closing bbolt database: %w", err)
	}

	log.Info("bbolt database closed")

	return nil
}