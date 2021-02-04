package clients

import (
	"github.com/allez-chauffe/marcel/api/db/internal/db"
)

var store db.Store

func CreateStore(database db.Databse) {
	store = database.CreateStore(func() db.Entity {
		return new(Client)
	})
}

func Get(id string) (*Client, error) {
	var result = &Client{}
	return result, store.Get(id, &result)
}

func List() ([]Client, error) {
	var clients = []Client{}
	return clients, store.List(&clients)
}

func Insert(c *Client) error {
	return store.Insert(c)
}

func Update(c *Client) error {
	return store.Update(c)
}

func Delete(id string) error {
	return store.Delete(id)
}

func DeleteAll() error {
	return store.DeleteAll()
}
