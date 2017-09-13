package notifier

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Zenika/MARCEL/backend/commons"
	"github.com/gorilla/mux"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

/*Service is a websocket connection handler.
It handles connection request and dispatch them to the correct Media goroutine
*/
type Service struct {
	medias map[int]*Media
}

//NewService create a fresh new Service
func NewService() *Service {
	return new(Service)
}

//HandleMediaConnection Handles a connection request to a given media.
func (s *Service) HandleMediaConnection(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mediaID, err := strconv.Atoi(vars["idMedia"])

	_, err = upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("Websocket establishement failed for media %d : %s", mediaID, err.Error())
		commons.WriteResponse(w, http.StatusInternalServerError, "Failed to establish websocket connection")
		return
	}
}

//RegisterMedia open a new goroutine for the given Media
func (s *Service) RegisterMedia(mediaID int) {
	if _, found := s.medias[mediaID]; found {
		return
	}

	s.medias[mediaID] = newMedia(mediaID)
}

//UnregisterMedia close the Media gotourine and all its clients.
func (s *Service) UnregisterMedia(mediaID int) {
	media, found := s.medias[mediaID]
	if !found {
		return
	}

	media.Close()
	delete(s.medias, mediaID)
}
