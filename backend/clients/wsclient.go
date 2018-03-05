package clients

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

//WSClient represent a websocket connection to a client.
type WSClient struct {
	service    *Service
	client     *Client
	conn       *websocket.Conn
	pingTicker *time.Ticker
	send       chan string
	receive    chan string
	close      chan chan struct{}
	closeNotif chan struct{}
}

func (ws *WSClient) run() {
	ws.pingTicker = time.NewTicker(pingPeriod)
	var closeNotif chan bool

	for {
		select {

		case msg := <-ws.send:
			ws.writeMessage(
				msg,
				fmt.Sprintf("Send message to client %s : %s", ws, msg),
				fmt.Sprintf("Write failed %s : ", ws),
			)

		case <-ws.pingTicker.C:
			ws.writeEmptyMessage(
				websocket.PingMessage,
				"", fmt.Sprintf("Ping failed %s : ", ws),
			)

		case msg := <-ws.receive:
			log.Printf("Message received from %s : %s", ws, msg)

		case closeNoif := <-ws.close:
			ws.cleanUp()
			if closeNotif != nil {
				closeNoif <- struct{}{}
			}
			return
		}
	}
}

func (ws *WSClient) listen() {
	//Setting websocket config and ping handling
	ws.conn.SetReadLimit(maxMessageSize)
	ws.conn.SetReadDeadline(time.Now().Add(pongWait))
	ws.conn.SetPongHandler(func(_ string) error {
		ws.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		msgType, message, err := ws.conn.ReadMessage()
		if err != nil {
			ws.Close()
			return
		}

		if msgType == websocket.TextMessage {
			ws.receive <- string(message)
		}
	}
}

func newWSClient(service *Service, client *Client, conn *websocket.Conn) *WSClient {
	ws := &WSClient{
		service: service,
		client:  client,
		conn:    conn,
		send:    make(chan string),
		receive: make(chan string),
		close:   make(chan chan struct{}, 1),
	}

	go ws.run()
	go ws.listen()

	service.register <- ws

	return ws
}

func (ws *WSClient) cleanUp() {
	ws.conn.Close()
	close(ws.send)
	close(ws.close)
	ws.service.unregister <- ws
}

//Close gracefully close and cleanup the client
func (ws *WSClient) Close() {
	ws.close <- nil
}

//CloseAndWait close and cleanup the client.
//An empty struct will be sent throught the returned chan after clean up.
func (ws *WSClient) CloseAndWait() <-chan struct{} {
	closeNotif := make(chan struct{})
	ws.close <- closeNotif
	return closeNotif
}

func (ws *WSClient) String() string {
	return fmt.Sprintf("%s-%s (%s)", ws.client.Name, ws.client.ID, ws.client.Type)
}
