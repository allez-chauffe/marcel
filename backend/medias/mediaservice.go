package medias

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/Zenika/MARCEL/backend/commons"
	"github.com/Zenika/MARCEL/backend/notifier"
	"github.com/Zenika/MARCEL/backend/plugins"
	"github.com/gorilla/mux"
)

const MEDIAS_CONFIG_PATH string = "data"
const MEDIAS_CONFIG_FILENAME string = "medias.json"

type Service struct {
	manager  *Manager
	notifier *notifier.Service
}

func NewService(pluginManager *plugins.Manager, notifier *notifier.Service) *Service {
	service := new(Service)

	service.manager = NewManager(pluginManager, notifier, MEDIAS_CONFIG_PATH, MEDIAS_CONFIG_FILENAME)
	service.notifier = notifier

	return service
}

func (m *Service) GetManager() *Manager {
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
	media := m.getMediaFromRequest(w, r)
	if media == nil {
		return
	}

	b, err := json.Marshal(*media)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotFound, err.Error())
		return
	}

	commons.WriteResponse(w, http.StatusOK, (string)(b))
}

// swagger:route POST /medias SaveHandler
//
// SaveIntoDB a media.
// If it's an update of an existing media, it will be first deactivated (all plugins stopped)
//  prior to be activated and saved.
// By default, the media will be activated
//
//     Consumes:
//     - application/json
//
//     Schemes: http, https
func (m *Service) SaveHandler(w http.ResponseWriter, r *http.Request) {
	// 1 : Get content and check structure
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		commons.WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// 2 : if media != new => stop all plugins'backend containers
	var media *Media = NewMedia()
	err = json.Unmarshal(body, &media)
	//if it's a new media (id==0) : create one
	if tmpMedia, _ := m.manager.Get(media.ID); tmpMedia != nil {
		m.manager.Deactivate(tmpMedia)
	} else {
		//it's a new media, let give it an ID
		media.ID = m.manager.GetNextID()
	}

	m.manager.SaveIntoDB(media)
	// 3 : start backend for every plugin instance
	m.manager.Activate(media)

	m.manager.Commit()

	commons.WriteResponse(w, http.StatusOK, "Media correctly saved with ID "+strconv.Itoa(media.ID))
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
	newMedia := m.manager.CreateEmpty()

	//return it to the client
	b, err := json.Marshal(*newMedia)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotFound, err.Error())
		return
	}

	commons.WriteResponse(w, http.StatusOK, (string)(b))
}

// swagger:route GET /medias/{idMedia:[0-9]*}/activate ActivateHandler
//
// If the media was deactivated (IsActive==false), backends for its plugins are started
//
//     Schemes: http, https
func (m *Service) ActivateHandler(w http.ResponseWriter, r *http.Request) {
	media := m.getMediaFromRequest(w, r)

	if media != nil {
		m.manager.Activate(media)
		m.notifier.RegisterMedia(media.ID)

		m.manager.Commit()

		commons.WriteResponse(w, http.StatusOK, "Media is active")
	}
}

// swagger:route GET /medias/{idMedia:[0-9]*}/deactivate DeactivateHandler
//
// If the media was activated (IsActive==true), backends for its plugins are stopped
func (m *Service) DeactivateHandler(w http.ResponseWriter, r *http.Request) {
	media := m.getMediaFromRequest(w, r)

	if media != nil {
		m.manager.Deactivate(media)
		m.notifier.UnregisterMedia(media.ID)

		m.manager.Commit()

		commons.WriteResponse(w, http.StatusOK, "Media has been deactivated")
	}
}

// swagger:route GET /medias/{idMedia:[0-9]*}/restart RestartHandler
//
// restart backends for the plugins of this media
func (m *Service) RestartHandler(w http.ResponseWriter, r *http.Request) {
	media := m.getMediaFromRequest(w, r)
	if media == nil {
		return
	}

	if media.IsActive {
		m.manager.Deactivate(media)
	}
	m.manager.Activate(media)

	m.manager.Commit()

	commons.WriteResponse(w, http.StatusOK, "Media has been correctly restarted")
}

// swagger:route DELETE /medias/{idMedia:[0-9]*} DeleteHandler
//
// Delete this media
func (m *Service) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	media := m.getMediaFromRequest(w, r)

	if media != nil {
		m.manager.Delete(media)

		commons.WriteResponse(w, http.StatusOK, "Media has been correctly deleted")
	}
}

// swagger:route DELETE /medias DeleteAllHandler
//
// Delete all medias
func (m *Service) DeleteAllHandler(w http.ResponseWriter, r *http.Request) {

	for i := len(m.manager.Config.Medias) - 1; i >= 0; i-- {
		media := m.manager.Config.Medias[i]
		m.manager.Deactivate(&media)

		m.manager.RemoveFromDB(&media)
		m.manager.Commit()

	}
	commons.WriteResponse(w, http.StatusOK, "All medias have been correctly deleted")
}

// swagger:route GET /medias{idMedia:[0-9]*}/plugins/{eltName}/{instanceId}/*
//
// Serves static frontend files of the given plugin instance for the given media.
func (m *Service) GetPluginFilesHandler(w http.ResponseWriter, r *http.Request) {
	const sep = string(os.PathSeparator)
	vars := mux.Vars(r)
	eltName := vars["eltName"]
	instanceID := vars["instanceId"]
	filePath := vars["filePath"]

	if filePath == "" {
		filePath = "index.html"
	}

	if media := m.getMediaFromRequest(w, r); media != nil {
		pluginDirectoryPath := m.manager.GetPluginDirectory(media, eltName, instanceID)
		pluginFilePath := pluginDirectoryPath + sep + "frontend" + sep + filePath
		http.ServeFile(w, r, pluginFilePath)
	}
}

func (m *Service) getMediaFromRequest(w http.ResponseWriter, r *http.Request) (media *Media) {
	vars := mux.Vars(r)
	attr := vars["idMedia"]

	idMedia, err := strconv.Atoi(attr)
	if err != nil {
		commons.WriteResponse(w, http.StatusBadRequest, err.Error())
		return nil
	}

	media, err = m.manager.Get(idMedia)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotFound, err.Error())
		return nil
	}

	return media
}
