package clients

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"

	"github.com/allez-chauffe/marcel/api/commons"
	"github.com/allez-chauffe/marcel/api/db/clients"
)

//getClientFromRequest return (if any) the client configuration based on id fiven in URL
func (s *Service) getClientFromRequest(w http.ResponseWriter, r *http.Request) (*clients.Client, error) {
	vars := mux.Vars(r)
	clientID, found := vars["clientID"]

	if !found {
		log.Errorf("Malformed URL (missing client id)")
		commons.WriteResponse(w, http.StatusBadRequest, "Malformed URL (missing client id)")
		return nil, errors.New("Malformed URL (missing client id)")
	}

	client, err := clients.Get(clientID)
	if err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return nil, err
	}
	if client == nil {
		log.Errorf("Unknown client : %s", clientID)
		commons.WriteResponse(w, http.StatusNotFound, "Client not found")
		return nil, errors.New("Client not found")
	}

	return client, nil
}

//getClientFromRequest return the client configuration parsed from the request body
func (s *Service) getClientFromRequestBody(w http.ResponseWriter, r *http.Request) (*clients.Client, error) {
	client := new(clients.Client)

	if err := json.NewDecoder(r.Body).Decode(client); err != nil {
		commons.WriteResponse(w, http.StatusBadRequest, err.Error())
		return nil, err
	}

	return client, nil
}

func (ws *WSClient) writeMessageWithType(msgType int, msg []byte, logMsg string, errorMsg string) bool {
	if logMsg != "" {
		log.Debugln(logMsg)
	}

	ws.conn.SetWriteDeadline(time.Now().Add(writeWait))
	out, err := ws.conn.NextWriter(msgType)

	if err == nil {
		_, err = out.Write(msg)
	}

	if err == nil {
		err = out.Close()
	}

	if err != nil {
		if errorMsg != "" {
			log.Errorln(errorMsg, err)
		}
		ws.Unregister()
	}

	return err == nil
}

func (s *Service) getClientPayload(client *clients.Client) *ClientPayload {
	_, isConnected := s.wsclients[client.ID]
	return &ClientPayload{client, isConnected}
}

func (s *Service) getClientsPayload() (map[string]*ClientPayload, error) {
	clients, err := clients.List()
	if err != nil {
		return nil, err
	}

	clientsPayload := map[string]*ClientPayload{}
	for _, client := range clients {
		clientsPayload[client.ID] = s.getClientPayload(&client)
	}

	return clientsPayload, nil
}

func (ws *WSClient) writeMessage(msg string, logMsg string, errorMsg string) bool {
	return ws.writeMessageWithType(websocket.TextMessage, []byte(msg), logMsg, errorMsg)
}

func (ws *WSClient) writeEmptyMessage(msgType int, logMsg string, errorMsg string) bool {
	return ws.writeMessageWithType(msgType, []byte{}, logMsg, errorMsg)
}
