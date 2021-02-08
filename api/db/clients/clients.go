package clients

import "github.com/allez-chauffe/marcel/api/db/internal/db"

var DefaultStore *Store

type Store struct {
	store db.Store
}

func CreateStore() error {
	store, err := db.DB.CreateStore(func() db.Entity {
		return new(Client)
	})
	if err != nil {
		return err
	}

	DefaultStore = &Store{store}

	return nil
}

func Transactional(tx db.Transaction) *Store {
	return &Store{DefaultStore.store.Transactional(tx)}
}

func (b *Store) Get(id string) (*Client, error) {
	c := &Client{}
	return c, b.store.Get(id, &c)
}

func (b *Store) Exists(id string) (bool, error) {
	return b.store.Exists(id)
}

func (b *Store) List() ([]Client, error) {
	var clients = []Client{}
	return clients, b.store.List(&clients)
}

func (b *Store) Insert(c *Client) error {
	return b.store.Insert(c)
}

func (b *Store) Update(c *Client) error {
	return b.store.Update(c)
}

func (b *Store) Delete(id string) error {
	return b.store.Delete(id)
}

func (b *Store) DeleteAll() error {
	return b.store.DeleteAll()
}
