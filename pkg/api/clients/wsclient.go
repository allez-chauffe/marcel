package clients

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"

	"github.com/allez-chauffe/marcel/pkg/db/clients"
)

//WSClient represent a websocket connection to a client.
type WSClient struct {
	service   *Service
	client    *clients.Client
	conn      *websocket.Conn
	send      chan string
	waitClose chan chan bool
	receive   chan string
	close     chan error
}

func (ws *WSClient) run() {
	pingTicker := time.NewTicker(pingPeriod)
	var closeNotif chan bool

	for {
		select {

		case msg := <-ws.send:
			ws.writeMessage(
				msg,
				fmt.Sprintf("Send message to client %s : %s", ws, msg),
				fmt.Sprintf("Write failed %s : ", ws),
			)

		case closeNotif = <-ws.waitClose:

		case <-pingTicker.C:
			ws.writeEmptyMessage(
				websocket.PingMessage,
				"", fmt.Sprintf("Ping failed %s : ", ws),
			)

		case msg := <-ws.receive:
			log.Debugf("Message received from %s : %s", ws, msg)

		case <-ws.close:
			close(ws.send)
			close(ws.close)
			pingTicker.Stop()
			ws.writeEmptyMessage(
				websocket.CloseMessage,
				fmt.Sprintf("Closing client %s", ws), "",
			)
			select {
			case closeNotif <- true:
			default:
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
			ws.Unregister()
			return
		}
		if msgType == websocket.TextMessage {
			msg := string(message)
			switch msg {
			default:
				ws.receive <- string(msg)
			}
		}
	}
}

func newWSClient(service *Service, client *clients.Client, conn *websocket.Conn) *WSClient {
	ws := &WSClient{
		service,
		client,
		conn,
		make(chan string),
		make(chan chan bool),
		make(chan string),
		make(chan error, 1),
	}

	go ws.run()
	go ws.listen()

	service.register <- ws

	return ws
}

//Unregister wsclient from its service (the call the Close will be made by the service)
func (ws *WSClient) Unregister() {
	ws.service.unregister <- ws
}

//Close gracefully close and cleanup the client
func (ws *WSClient) Close() {
	select {
	case ws.close <- nil:
	default:
	}
}

func (ws *WSClient) String() string {
	return fmt.Sprintf("%s-%s (%s)", ws.client.Name, ws.client.ID, ws.client.Type)
}
