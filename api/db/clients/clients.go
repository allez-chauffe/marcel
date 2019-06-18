package clients

import (
	uuid "github.com/satori/go.uuid"
	bh "github.com/timshannon/bolthold"

	"github.com/Zenika/marcel/api/db/internal/db"
)

func Get(id string) (*Client, error) {
	var c = new(Client)

	if err := db.Store.Get(id, c); err != nil {
		if err == bh.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return c, nil
}

func List() ([]Client, error) {
	var clients = []Client{}

	return clients, db.Store.Find(&clients, nil)
}

func Insert(c *Client) error {
	c.ID = uuid.NewV4().String()

	return db.Store.Insert(c.ID, c)
}

func Update(c *Client) error {
	return db.Store.Update(c.ID, c)
}

func Delete(id string) error {
	return db.Store.Delete(id, &Client{})
}

func DeleteAll() error {
	return db.Store.DeleteMatching(&Client{}, nil)
}
