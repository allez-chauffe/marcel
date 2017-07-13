package plugins

import (
	"net/http"
	"encoding/json"
	"github.com/Zenika/MARCEL/backend/commons"
)

const PLUGINS_CONFIG_PATH string = "data"
const PLUGINS_CONFIG_FILENAME string = "plugins.json"

type Service struct {
	Manager *Manager
}

func NewService() *Service {
	var p = new(Service)

	c := NewConfiguration()
	p.Manager = NewManager(PLUGINS_CONFIG_PATH, PLUGINS_CONFIG_FILENAME, c)

	return p
}


func (s *Service) GetManager() (*Manager) {
	return s.Manager
}


// swagger:route GET /plugins/config GetConfigHandler
//
// Gets information of all plugins
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
func (m *Service) GetConfigHandler(w http.ResponseWriter, r *http.Request) {

	c := m.Manager.GetConfiguration()
	b, err := json.Marshal(c)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotFound)
		return
	}

	w.Write([]byte(b))
}

// swagger:route GET /plugins GetAllHandler
//
// Gets information of all plugins
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
func (m *Service) GetAllHandler(w http.ResponseWriter, r *http.Request) {

	media := m.Manager.GetAll()
	b, err := json.Marshal(media)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotFound)
		return
	}

	w.Write([]byte(b))
}