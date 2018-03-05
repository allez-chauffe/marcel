package clients

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Zenika/MARCEL/backend/commons"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

//getClientFromRequest return (if any) the client configuration based on id fiven in URL
func (s *Service) getClientFromRequest(w http.ResponseWriter, r *http.Request) (*Client, bool) {
	vars := mux.Vars(r)
	clientID, found := vars["clientID"]

	if !found {
		log.Printf("Malformed URL (missing client id)")
		commons.WriteResponse(w, http.StatusBadRequest, "Malformed URL (missing client id)")
		return nil, false
	}

	client, clientFound := s.manager.Get(clientID)
	if !clientFound {
		log.Printf("Unknown client : %s", clientID)
		commons.WriteResponse(w, http.StatusNotFound, "Client not found")
		return nil, false
	}

	return client, true
}

//getClientFromRequest return the client configuration parsed from the request body
func (s *Service) getClientFromRequestBody(w http.ResponseWriter, r *http.Request) (*Client, bool) {
	client := &Client{}
	if err := json.NewDecoder(r.Body).Decode(client); err != nil {
		commons.WriteResponse(w, http.StatusBadRequest, err.Error())
		return nil, false
	}

	return client, true
}

func (ws *WSClient) writeMessageWithType(msgType int, msg []byte, logMsg string, errorMsg string) bool {
	if logMsg != "" {
		log.Println(logMsg)
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
			log.Println(errorMsg, err)
		}
		ws.Close()
	}

	return err == nil
}

func (s *Service) getClientJson(client *Client) *ClientJSON {
	_, isConnected := s.wsclients[client.ID]
	return &ClientJSON{client, isConnected}
}

func (s *Service) getClientsJson() map[string]*ClientJSON {
	clients := map[string]*ClientJSON{}
	for id, client := range s.manager.GetAll() {
		clients[id] = s.getClientJson(client)
	}
	return clients
}

func (ws *WSClient) writeMessage(msg string, logMsg string, errorMsg string) bool {
	return ws.writeMessageWithType(websocket.TextMessage, []byte(msg), logMsg, errorMsg)
}

func (ws *WSClient) writeEmptyMessage(msgType int, logMsg string, errorMsg string) bool {
	return ws.writeMessageWithType(msgType, []byte{}, logMsg, errorMsg)
}
