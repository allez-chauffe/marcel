package bolt

import (
	"fmt"
	"os"
	"reflect"
	"time"

	log "github.com/sirupsen/logrus"
	bh "github.com/timshannon/bolthold"
	bolt "go.etcd.io/bbolt"

	"github.com/allez-chauffe/marcel/pkg/config"
	"github.com/allez-chauffe/marcel/pkg/db/internal/db"
)

type driver struct{}

// Driver is the bbolt driver implementation.
var Driver db.Driver = new(driver)

func (d *driver) Open() (db.Client, error) {
	log.Info("Opening bbolt database...")

	var options = *bolt.DefaultOptions
	options.Timeout = 100 * time.Millisecond

	bhs, err := bh.Open(os.ExpandEnv(config.Default().API().DB().Bolt().File()), 0644, &bh.Options{
		Options: &options,
	})
	if err != nil {
		return nil, fmt.Errorf("Error while opening bbolt database: %w", err)
	}

	log.Info("bbolt database opened")

	return &client{bhs}, nil
}

type client struct {
	bh *bh.Store
}

func (c *client) Store(newEntity func() db.Entity) (db.Store, error) {
	return &store{c, newEntity, reflect.TypeOf(newEntity()).Elem().Name(), nil}, nil
}

func (c *client) Begin() (db.Transaction, error) {
	return c.bh.Bolt().Begin(true)
}

func (c *client) Close() error {
	log.Info("Closing bbolt database...")

	err := c.bh.Close()
	if err != nil {
		return fmt.Errorf("Error while closing bbolt database: %w", err)
	}

	log.Info("bbolt database closed")

	return nil
}
