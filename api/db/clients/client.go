package clients

import (
	uuid "github.com/satori/go.uuid"
)

type Client struct {
	ID      string `json:"id" boltholdKey:"ID" structs:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	MediaID int    `json:"mediaID"`
}

func New() *Client {
	return &Client{
		ID:   uuid.NewV4().String(),
		Type: "Unkown",
	}
}

func (c *Client) GetID() interface{} {
	return c.ID
}

func (c *Client) SetID(id interface{}) {
	switch cleanID := id.(type) {
	case string:
		c.ID = cleanID
	case []byte:
		c.ID = string(cleanID)
	default:
		panic("Unsuported id type")
	}
}
