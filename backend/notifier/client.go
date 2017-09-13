package notifier

import (
	"github.com/gorilla/websocket"
)

//Client represent a websocket connection to a client.
type Client struct {
	media *Media
	conn  *websocket.Conn
	close chan error
}

func newClient(media *Media, conn *websocket.Conn) *Client {
	client := &Client{
		media,
		conn,
		make(chan error, 1),
	}
	go client.run()
	return client
}

func (c *Client) run() {
	for {
		select {
		case <-c.close:
			c.conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}
	}
}

//Close gracefully close and cleanup the client
func (c *Client) Close() {
	c.close <- nil
}
