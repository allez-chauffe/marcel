package medias

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/Zenika/MARCEL/backend/commons"
	"io/ioutil"
)

const MEDIAS_CONFIG_PATH string = "data"
const MEDIAS_CONFIG_FILENAME string = "medias.json"

type Service struct {
	manager *Manager
}

func NewService() *Service {
	service := new(Service)

	service.manager = NewManager(MEDIAS_CONFIG_PATH, MEDIAS_CONFIG_FILENAME)

	return service
}

func (m *Service) GetManager() (*Manager) {
	return m.manager
}

// swagger:route GET /medias/config GetConfigHandler
//
// Gets information of all medias
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
func (m *Service) GetConfigHandler(w http.ResponseWriter, r *http.Request) {

	c := m.manager.GetConfiguration()
	b, err := json.Marshal(c)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotFound, err.Error())
		return
	}

	commons.WriteResponse(w, http.StatusOK, (string)(b))
}

// swagger:route GET /medias GetAllHandler
//
// Gets information of all medias
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
func (m *Service) GetAllHandler(w http.ResponseWriter, r *http.Request) {

	media := m.manager.GetAll()
	b, err := json.Marshal(media)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotFound, err.Error())
		return
	}

	commons.WriteResponse(w, http.StatusOK, (string)(b))
}

// swagger:route GET /medias/{idMedia} GetHandler
//
// Gets information of a media
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
// swagger:parameters idMedia
func (m *Service) GetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	attr := vars["idMedia"]

	idMedia, err := strconv.Atoi(attr)
	if err != nil {
		commons.WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	media, err := m.manager.Get(idMedia)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotFound, err.Error())
		return
	}

	b, err := json.Marshal(*media)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotFound, err.Error())
		return
	}

	commons.WriteResponse(w, http.StatusOK, (string)(b))
}

// swagger:route POST /medias PostHandler
//
// Posts information for a media to be saved.
// If it's an update of an existing media, it will be first stopped (all plugins stopped)
//  priori to be activated and saved.
// By default, the media will be activated
//
//     Consumes:
//     - application/json
//
//     Schemes: http, https
func (m *Service) PostHandler(w http.ResponseWriter, r *http.Request) {
	// 1 : Get content and check structure
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		commons.WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// 2 : if media != new => stop all plugins'backend containers
	var media *Media = NewMedia()
	err = json.Unmarshal(body, &media)
	if tmpMedia, _ := m.manager.Get(media.ID); tmpMedia != nil {
		m.manager.Deactivate(tmpMedia)
	}


	m.manager.Save(media)
	// 3 : start backend for every plugin instance
	m.manager.Activate(media)

	m.manager.Commit()
}

// swagger:route GET /medias CreateHandler
//
// Gets information of all medias
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
func (m *Service) CreateHandler(w http.ResponseWriter, r *http.Request) {
	//get a new media
	newMedia := m.manager.Create()

	//return it to the client
	b, err := json.Marshal(*newMedia)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotFound, err.Error())
		return
	}

	commons.WriteResponse(w, http.StatusOK, (string)(b))
}

func (m *Service) ActivateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	attr := vars["idMedia"]

	idMedia, err := strconv.Atoi(attr)
	if err != nil {
		commons.WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	media, err := m.manager.Get(idMedia)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotFound, err.Error())
		return
	}

	if !media.IsActive {
		m.manager.Activate(media)
		m.manager.Commit()
	}
}

func (m *Service) DeactivateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	attr := vars["idMedia"]

	idMedia, err := strconv.Atoi(attr)
	if err != nil {
		commons.WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	media, err := m.manager.Get(idMedia)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotFound, err.Error())
		return
	}

	if media.IsActive {
		m.manager.Deactivate(media)
		m.manager.Commit()
	}
}

func (m *Service) RestartHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	attr := vars["idMedia"]

	idMedia, err := strconv.Atoi(attr)
	if err != nil {
		commons.WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	media, err := m.manager.Get(idMedia)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotFound, err.Error())
		return
	}

	if media.IsActive {
		m.manager.Deactivate(media)
	}
	m.manager.Activate(media)

	m.manager.Commit()
}