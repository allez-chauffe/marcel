package notifier

import "log"

type clients map[*Client]bool

//Media represent a Media goroutine, waiting for new clients and notify them when nedeed.
type Media struct {
	mediaID    int
	clients    clients
	register   chan *Client
	unregister chan *Client
	close      chan error
}

func newMedia(mediaID int) *Media {
	media := &Media{
		mediaID,
		make(clients),
		make(chan *Client),
		make(chan *Client, 1),
		make(chan error, 1),
	}
	go media.run()
	return media
}

func (m *Media) run() {
	for {
		select {
		case client := <-m.register:
			log.Printf("Register client for media %d", m.mediaID)
			m.clients[client] = true

		case client := <-m.unregister:
			if _, ok := m.clients[client]; ok {
				log.Printf("Unregister client for media %d", m.mediaID)
				delete(m.clients, client)
				client.Close()
			}

		case <-m.close:
			for client := range m.clients {
				client.Close()
			}
			close(m.register)
			close(m.unregister)
			close(m.close)
			return
		}
	}
}

//Close clean up channels and close all connected clients
func (m *Media) Close() {
	m.close <- nil
}
