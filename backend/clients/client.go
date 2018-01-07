package clients

import (
	"github.com/satori/go.uuid"
)

//Client represent the mapping between a ID and a readable name given by user
type Client struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	MediaID int    `json:"mediaID"`
}

//ClientJSON respresent the client that will be sent to back-office
type ClientJSON struct {
	*Client
	IsConnected bool `json:"isConnected"`
}

func newClient() *Client {
	return &Client{
		ID:      uuid.Must(uuid.NewV4()).String(),
		Name:    "",
		Type:    "Unkown",
		MediaID: 0,
	}
}
