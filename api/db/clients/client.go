package clients

import (
	uuid "github.com/satori/go.uuid"
)

type Client struct {
	ID      string `json:"id"`
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
