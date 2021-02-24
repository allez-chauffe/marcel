package clients

import "github.com/allez-chauffe/marcel/pkg/db/internal/db"

type Store struct {
	store db.StoreBase
}

func CreateStore(database db.Client) (*Store, error) {
	store, err := database.CreateStore(func() db.Entity {
		return new(Client)
	})
	if err != nil {
		return nil, err
	}

	return &Store{store}, nil
}

func (s *Store) Transactional(tx db.Transaction) *Store {
	return &Store{s.store.Transactional(tx)}
}

func (s *Store) Get(id string) (*Client, error) {
	c := &Client{}
	return c, s.store.Get(id, &c)
}

func (s *Store) Exists(id string) (bool, error) {
	return s.store.Exists(id)
}

func (s *Store) List() ([]Client, error) {
	var clients = []Client{}
	return clients, s.store.List(&clients)
}

func (s *Store) Insert(c *Client) error {
	return s.store.Insert(c)
}

func (s *Store) Update(c *Client) error {
	return s.store.Update(c)
}

func (s *Store) Delete(id string) error {
	return s.store.Delete(id)
}

func (s *Store) DeleteAll() error {
	return s.store.DeleteAll()
}
