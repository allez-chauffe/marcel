package clients

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Zenika/MARCEL/backend/commons"
	"github.com/gorilla/websocket"
)

const (
	configPath     string = "data"
	configFileName string = "clients.json"
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

//Create a new
func NewService() *Service {
	service := &Service{
		make(wsclients),
		newManager(configPath, configFileName),
		make(chan *WSClient),
		make(chan *WSClient),
	}

	go service.run()

	return service
}

//WSConnectionHandler Handles a connection request from a given client.
func (s *Service) WSConnectionHandler(w http.ResponseWriter, r *http.Request) {
	client, found := s.getClientFromRequest(w, r)
	if !found {
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Websocket establishement failed for client %s : %s", client.ID, err.Error())
		commons.WriteResponse(w, http.StatusInternalServerError, "Failed to establish websocket connection")
		return
	}

	newWSClient(s, client, conn)
}

//GetHandler send the request client configuration.
func (s *Service) GetHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting client configuration")

	client, exists := s.getClientFromRequest(w, r)
	if !exists {
		return
	}

	_, isConnected := s.wsclients[client.ID]

	commons.WriteJsonResponse(w, ClientJSON{client, isConnected})
}

//GetAllHandler send the list of all registered clients.
func (s *Service) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting all clients configuration")

	clients := s.manager.GetAll()

	result := map[string]ClientJSON{}
	for id, client := range clients {
		_, isConnected := s.wsclients[client.ID]
		result[id] = ClientJSON{client, isConnected}
	}

	commons.WriteJsonResponse(w, result)
}

//CreateHandler create a new client entry in the database.
func (s *Service) CreateHandler(w http.ResponseWriter, r *http.Request) {
	body := map[string]interface{}{}
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		commons.WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	nameI, ok := body["name"]
	name, ok := nameI.(string)
	if !ok || name == "" {
		name = "New Client"
	}

	mediaIDI, ok := body["mediaID"]
	mediaID, ok := mediaIDI.(int)
	if !ok || mediaID < 0 {
		mediaID = 0
	}

	client := s.manager.addNewClient(name, mediaID)
	log.Printf("Created a new client : %v", client)
	commons.WriteJsonResponse(w, ClientJSON{client, false})
}

//DeleteHandler delete a client from the database
func (s *Service) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	client, exists := s.getClientFromRequest(w, r)

	if !exists {
		return
	}

	s.manager.deleteClient(client.ID)
	log.Printf("Deleted client %s-%s (%s)", client.Name, client.ID, client.Type)
	commons.WriteResponse(w, http.StatusNoContent, "")
}

func (s *Service) DeleteAllHandler(w http.ResponseWriter, r *http.Request) {
	s.manager.deleteAllClients()
	for _, ws := range s.wsclients {
		s.unregister <- ws
	}
	log.Println("All client deleted and disconnected")
	commons.WriteResponse(w, http.StatusNoContent, "")
}

//UpdateHandler update a client configuration in the database
func (s *Service) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	client, ok := s.getClientFromRequestBody(w, r)
	if !ok {
		return
	}

	_, exists := s.manager.Get(client.ID)
	if !exists {
		commons.WriteResponse(w, http.StatusNotFound, "Client not found")
		return
	}

	s.manager.updateClient(client)
	s.SendByID(client.ID, "update")
	<-s.WaitForClose(client.ID)
	commons.WriteResponse(w, http.StatusNoContent, "")
}

//Run is a goroutine managing connected client list (s.clients)
func (s *Service) run() {
	for {
		select {
		case ws := <-s.register:
			log.Printf("Client connected : %s-%s (%s)", ws.client.Name, ws.client.ID, ws.client.Type)
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
	log.Printf("Sending %q to all clients of media %d", msg, mediaID)
	for _, ws := range s.wsclients {
		if ws.client.MediaID == mediaID {
			log.Printf("Sending %q to %s", msg, ws)
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

	log.Printf("Sending %q to client %s", msg, ws)
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
