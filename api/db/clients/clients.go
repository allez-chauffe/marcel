package clients

import "github.com/allez-chauffe/marcel/api/db/internal/db"

var DefaultBucket *Bucket

type Bucket struct {
	store db.Store
}

func CreateDefaultBucket() error {
	store, err := db.DB.CreateStore(func() db.Entity {
		return new(Client)
	})

	if err != nil {
		return err
	}

	DefaultBucket = &Bucket{store}
	return nil
}

func Transactional(tx db.Transaction) *Bucket {
	return &Bucket{DefaultBucket.store.Transactional(tx)}
}

func (b *Bucket) Get(id string) (*Client, error) {
	c := &Client{}
	return c, b.store.Get(id, &c)
}

func (b *Bucket) Exists(id string) (bool, error) {
	return b.store.Exists(id)
}

func (b *Bucket) List() ([]Client, error) {
	var clients = []Client{}
	return clients, b.store.List(&clients)
}

func (b *Bucket) Insert(c *Client) error {
	return b.store.Insert(c)
}

func (b *Bucket) Update(c *Client) error {
	return b.store.Update(c)
}

func (b *Bucket) Delete(id string) error {
	return b.store.Delete(id)
}

func (b *Bucket) DeleteAll() error {
	return b.store.DeleteAll()
}
