package plugins

import (
	"net/http"
	"encoding/json"
	"github.com/Zenika/MARCEL/backend/commons"
	"github.com/gorilla/mux"
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
func (s *Service) GetConfigHandler(w http.ResponseWriter, r *http.Request) {

	c := s.Manager.GetConfiguration()
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

// swagger:route GET /plugins/{idMedia} GetHandler
//
// Gets information of a plugin
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
// swagger:parameters idPlugin
func (s *Service) GetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eltName := vars["eltName"]

	plugin, err := s.Manager.Get(eltName)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotFound)
		return
	}

	b, err := json.Marshal(*plugin)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotFound)
		return
	}

	w.Write([]byte(b))
}

func (s* Service) UploadHandler(w http.ResponseWriter, r *http.Request) {
	// 1 : open zip file
	// 2 : parse description file
	// 3 : check there's no plugin already installed with same name or reject
	// 4 : unzip into /plugins folder
	// 5 : save into plugins.json file

	w.Write([]byte("Upload en cours..."))
}