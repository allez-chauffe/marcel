package clients

import (
	"encoding/json"
	"net/http"

	"github.com/Pallinder/go-randomdata"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"

	"github.com/Zenika/MARCEL/api/auth"
	"github.com/Zenika/MARCEL/api/commons"
	"github.com/Zenika/MARCEL/config"
)

//Service is the websocket connection handler
type Service struct {
	wsclients  wsclients
	manager    *Manager
	register   chan *WSClient
	unregister chan *WSClient
}

type wsclients map[string]*WSClient

type connRequest struct {
	conn   *websocket.Conn
	client Client
}

type newClientRequest struct {
	name    string
	mediaID int
}

//Create a new
func NewService() *Service {
	service := &Service{
		make(wsclients),
		newManager(config.Config.DataPath, config.Config.ClientsFile),
		make(chan *WSClient),
		make(chan *WSClient),
	}

	go service.run()

	return service
}

//WSConnectionHandler Handles a connection request from a given client.
func (s *Service) WSConnectionHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckPermissions(r, nil) {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	client, found := s.getClientFromRequest(w, r)
	if !found {
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
	if !auth.CheckPermissions(r, nil) {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	log.Debugln("Getting client configuration")

	client, exists := s.getClientFromRequest(w, r)
	if !exists {
		return
	}

	_, isConnected := s.wsclients[client.ID]

	commons.WriteJsonResponse(w, ClientJSON{client, isConnected})
}

//GetAllHandler send the list of all registered clients.
func (s *Service) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckPermissions(r, nil) {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	commons.WriteJsonResponse(w, s.getClientsJson())
}

//CreateHandler create a new client entry in the database.
func (s *Service) CreateHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckPermissions(r, nil) {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	params := &newClientRequest{}
	if err := json.NewDecoder(r.Body).Decode(params); err != nil {
		commons.WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if params.name == "" {
		params.name = randomdata.SillyName()
	}

	if params.mediaID < 0 {
		params.mediaID = 0
	}

	client := s.manager.addNewClient(params.name, params.mediaID)
	log.Debugf("Created a new client : %v", client)
	commons.WriteJsonResponse(w, ClientJSON{client, false})
}

//DeleteHandler delete a client from the database
func (s *Service) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckPermissions(r, nil) {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	client, exists := s.getClientFromRequest(w, r)

	if !exists {
		return
	}

	s.manager.deleteClient(client.ID)
	s.unregister <- s.wsclients[client.ID]
	log.Debugf("Deleted client %s-%s (%s)", client.Name, client.ID, client.Type)
	commons.WriteResponse(w, http.StatusNoContent, "")
}

func (s *Service) DeleteAllHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckPermissions(r, nil) {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	s.manager.deleteAllClients()
	for _, ws := range s.wsclients {
		s.unregister <- ws
	}

	log.Debugln("All client deleted and disconnected")
	commons.WriteResponse(w, http.StatusNoContent, "")
}

//UpdateHandler update a client configuration in the database
func (s *Service) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckPermissions(r, nil) {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	client, ok := s.getClientFromRequestBody(w, r)
	if !ok {
		return
	}

	_, exists := s.manager.Get(client.ID)
	if !exists {
		commons.WriteResponse(w, http.StatusNotFound, "Client not found")
		return
	}

	savedClient := s.manager.updateClient(client)
	s.SendByID(client.ID, "update")
	// <-s.WaitForClose(client.ID)
	commons.WriteJsonResponse(w, savedClient)
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

//GetManager returns the client manager of the service
func (s *Service) GetManager() *Manager {
	return s.manager
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
