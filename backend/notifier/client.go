package notifier

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

//Client represent a websocket connection to a client.
type Client struct {
	media *Media
	conn  *websocket.Conn
	send  chan []byte
	close chan error
}

func newClient(media *Media, conn *websocket.Conn) *Client {
	client := &Client{
		media,
		conn,
		make(chan []byte),
		make(chan error, 1),
	}
	go client.run()
	return client
}

func (c *Client) run() {
	for {
		select {

		case msg := <-c.send:
			log.Println("Notify client : " + string(msg[:]))
			if err := c.writeMessage(websocket.TextMessage, msg); err != nil {
				log.Println("Write failed : ", err)
				c.media.unregister <- c
				return
			}

		case <-c.close:
			close(c.send)
			close(c.close)
			c.writeMessage(websocket.CloseMessage, []byte{})
			return
		}
	}
}

//Close gracefully close and cleanup the client
func (c *Client) Close() {
	c.close <- nil
}

func (c *Client) writeMessage(msgType int, msg []byte) error {
	c.conn.SetWriteDeadline(time.Now().Add(writeWait))
	w, err := c.conn.NextWriter(msgType)
	if err != nil {
		return err
	}

	w.Write(msg)
	return w.Close()
}
