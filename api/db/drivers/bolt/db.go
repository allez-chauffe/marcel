package bolt

import (
	"fmt"
	"os"
	"reflect"
	"time"

	log "github.com/sirupsen/logrus"
	bh "github.com/timshannon/bolthold"
	bolt "go.etcd.io/bbolt"

	"github.com/allez-chauffe/marcel/api/db/internal/db"
	"github.com/allez-chauffe/marcel/pkg/config"
)

type boltDriver struct{}

var Driver db.Driver = new(boltDriver)

func (driver *boltDriver) Open() (db.Database, error) {
	log.Info("Opening bbolt database...")

	var options = *bolt.DefaultOptions
	options.Timeout = 100 * time.Millisecond

	database, err := bh.Open(os.ExpandEnv(config.Default().API().DB().Bolt().File()), 0644, &bh.Options{
		Options: &options,
	})
	if err != nil {
		return nil, fmt.Errorf("Error while opening bbolt database: %w", err)
	}

	log.Info("bbolt database opened")

	return &boltDatabase{database}, nil
}

type boltDatabase struct {
	bh *bh.Store
}

func (database *boltDatabase) CreateStore(newEntity func() db.Entity) (db.Store, error) {
	return &boltStore{
		&boltStoreConfig{database, newEntity, reflect.TypeOf(newEntity()).Elem().Name()},
		nil,
	}, nil
}

func (database *boltDatabase) Begin() (db.Transaction, error) {
	return database.bh.Bolt().Begin(true)
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
