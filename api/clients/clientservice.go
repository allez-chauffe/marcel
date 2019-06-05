package clients

import (
	"encoding/json"
	"net/http"

	"github.com/Pallinder/go-randomdata"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"

	"github.com/Zenika/MARCEL/api/auth"
	"github.com/Zenika/MARCEL/api/commons"
	"github.com/Zenika/MARCEL/api/db/clients"
)

type ClientPayload struct {
	*clients.Client
	IsConnected bool `json:"isConnected"`
}

//Service is the websocket connection handler
type Service struct {
	wsclients  wsclients
	register   chan *WSClient
	unregister chan *WSClient
}

type wsclients map[string]*WSClient

type connRequest struct {
	conn   *websocket.Conn
	client clients.Client
}

func NewService() *Service {
	service := &Service{
		make(wsclients),
		make(chan *WSClient),
		make(chan *WSClient),
	}

	go service.run()

	return service
}

// WSConnectionHandler Handles a connection request from a given client.
func (s *Service) WSConnectionHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckPermissions(r, nil, "user", "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	client, err := s.getClientFromRequest(w, r)
	if err != nil {
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("Websocket establishement failed for client %s : %s", client.ID, err.Error())
		commons.WriteResponse(w, http.StatusInternalServerError, "Failed to establish websocket connection")
		return
	}

	newWSClient(s, client, conn)
}

//GetHandler send the requested client configuration.
func (s *Service) GetHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckPermissions(r, nil, "user", "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	log.Debugln("Getting client configuration")

	client, err := s.getClientFromRequest(w, r)
	if err != nil {
		return
	}

	_, isConnected := s.wsclients[client.ID]

	commons.WriteJsonResponse(w, ClientPayload{client, isConnected})
}

//GetAllHandler send the list of all registered clients.
func (s *Service) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckPermissions(r, nil, "user", "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	payload, err := s.getClientsPayload()
	if err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	commons.WriteJsonResponse(w, payload)
}

//CreateHandler create a new client entry in the database.
func (s *Service) CreateHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckPermissions(r, nil, "user", "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	client := new(clients.Client)
	if err := json.NewDecoder(r.Body).Decode(client); err != nil {
		commons.WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if client.Name == "" {
		client.Name = randomdata.SillyName()
	}

	if client.MediaID < 0 {
		client.MediaID = 0
	}

	if err := clients.Insert(client); err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Debugf("Created a new client : %v", client)

	commons.WriteJsonResponse(w, ClientPayload{client, false})
}

//DeleteHandler delete a client from the database
func (s *Service) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckPermissions(r, nil, "user", "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	client, err := s.getClientFromRequest(w, r)
	if err != nil {
		return
	}

	if err := clients.Delete(client.ID); err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	s.unregister <- s.wsclients[client.ID]

	log.Debugf("Deleted client %s-%s (%s)", client.Name, client.ID, client.Type)

	commons.WriteResponse(w, http.StatusNoContent, "")
}

func (s *Service) DeleteAllHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckPermissions(r, nil, "user", "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	if err := clients.DeleteAll(); err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	for _, ws := range s.wsclients {
		s.unregister <- ws
	}

	log.Debugln("All client deleted and disconnected")
	commons.WriteResponse(w, http.StatusNoContent, "")
}

//UpdateHandler update a client configuration in the database
func (s *Service) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckPermissions(r, nil, "user", "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	client, err := s.getClientFromRequestBody(w, r)
	if err != nil {
		return
	}

	if err := clients.Update(client); err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	s.SendByID(client.ID, "update")
	//FIXME what is this ?
	// <-s.WaitForClose(client.ID)

	commons.WriteJsonResponse(w, client)
}

//Run is a goroutine managing connected client list (s.clients)
func (s *Service) run() {
	for {
		select {
		case ws := <-s.register:
			log.Debugf("Client connected : %s-%s (%s)", ws.client.Name, ws.client.ID, ws.client.Type)
			s.wsclients[ws.client.ID] = ws
		case ws := <-s.unregister:
			if _, exists := s.wsclients[ws.client.ID]; exists {
				delete(s.wsclients, ws.client.ID)
				ws.Close()
			}
		}
	}
}

//SendByMedia sends a message to each client connected to the given media
func (s *Service) SendByMedia(mediaID int, msg string) {
	log.Debugf("Sending %q to all clients of media %d", msg, mediaID)
	for _, ws := range s.wsclients {
		if ws.client.MediaID == mediaID {
			log.Debugf("Sending %q to %s", msg, ws)
			select {
			case ws.send <- msg:
			default:
			}
		}
	}
}

func (s *Service) SendByID(clientID string, msg string) {
	ws, connected := s.wsclients[clientID]
	if !connected {
		return
	}

	log.Debugf("Sending %q to client %s", msg, ws)
	select {
	case ws.send <- msg:
	default:
	}
}

func (s *Service) WaitForClose(clientID string) <-chan bool {
	ws, connected := s.wsclients[clientID]

	if !connected {
		wait := make(chan bool, 1)
		wait <- true
		return wait
	}

	wait := make(chan bool)
	ws.waitClose <- wait
	return wait
}
