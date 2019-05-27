package medias

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/Zenika/MARCEL/api/auth"
	"github.com/Zenika/MARCEL/api/clients"
	"github.com/Zenika/MARCEL/api/commons"
	"github.com/Zenika/MARCEL/config"
)

type Service struct {
	manager        *Manager
	clientsService *clients.Service
}

func NewService(clientsService *clients.Service) *Service {
	service := new(Service)

	service.manager = NewManager(clientsService, config.Config.DataPath, config.Config.MediasFile)
	service.clientsService = clientsService

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
	if !auth.CheckPermissions(r, nil) {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	commons.WriteJsonResponse(w, m.manager.GetConfiguration())
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
	if !auth.CheckPermissions(r, nil) {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	commons.WriteJsonResponse(w, m.manager.GetAll())
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
	if !auth.CheckPermissions(r, nil) {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	if media := m.getMediaFromRequest(w, r); media != nil {
		commons.WriteJsonResponse(w, media)
	}
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
	media := &Media{}
	if err := json.NewDecoder(r.Body).Decode(media); err != nil {
		commons.WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	tmpMedia, _ := m.manager.Get(media.ID)
	if tmpMedia == nil {
		commons.WriteResponse(w, http.StatusNotFound, "")
		return
	}

	if !auth.CheckPermissions(r, []string{tmpMedia.Owner}, "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	m.manager.Deactivate(tmpMedia)
	m.manager.SaveIntoDB(media)
	m.manager.Activate(media)
	m.manager.Commit()

	commons.WriteJsonResponse(w, media)
	m.clientsService.SendByMedia(media.ID, "update")
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
	if !auth.CheckPermissions(r, nil) {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	//get a new media
	newMedia := m.manager.CreateEmpty(auth.GetAuth(r).Subject)

	commons.WriteJsonResponse(w, newMedia)
}

// swagger:route GET /medias/{idMedia:[0-9]*}/activate ActivateHandler
//
// If the media was deactivated (IsActive==false), backends for its plugins are started
//
//     Schemes: http, https
func (m *Service) ActivateHandler(w http.ResponseWriter, r *http.Request) {
	media := m.getMediaFromRequest(w, r)
	if media != nil {
		return
	}

	if !auth.CheckPermissions(r, []string{media.Owner}, "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	m.manager.Activate(media)

	m.manager.Commit()

	commons.WriteResponse(w, http.StatusOK, "Media is active")
	m.clientsService.SendByMedia(media.ID, "update")

}

// swagger:route GET /medias/{idMedia:[0-9]*}/deactivate DeactivateHandler
//
// If the media was activated (IsActive==true), backends for its plugins are stopped
func (m *Service) DeactivateHandler(w http.ResponseWriter, r *http.Request) {
	media := m.getMediaFromRequest(w, r)
	if media == nil {
		return
	}

	if !auth.CheckPermissions(r, []string{media.Owner}, "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	m.manager.Deactivate(media)

	m.manager.Commit()

	commons.WriteResponse(w, http.StatusOK, "Media has been deactivated")
	m.clientsService.SendByMedia(media.ID, "update")
}

// swagger:route GET /medias/{idMedia:[0-9]*}/restart RestartHandler
//
// restart backends for the plugins of this media
func (m *Service) RestartHandler(w http.ResponseWriter, r *http.Request) {
	media := m.getMediaFromRequest(w, r)
	if media == nil {
		return
	}

	if !auth.CheckPermissions(r, []string{media.Owner}, "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	if media.IsActive {
		m.manager.Deactivate(media)
	}
	m.manager.Activate(media)

	m.manager.Commit()

	commons.WriteResponse(w, http.StatusOK, "Media has been correctly restarted")
	m.clientsService.SendByMedia(media.ID, "update")
}

// swagger:route DELETE /medias/{idMedia:[0-9]*} DeleteHandler
//
// Delete this media
func (m *Service) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	media := m.getMediaFromRequest(w, r)
	if media == nil {
		return
	}

	if !auth.CheckPermissions(r, []string{media.Owner}, "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	m.manager.Delete(media)
	commons.WriteResponse(w, http.StatusOK, "Media has been correctly deleted")
}

// swagger:route DELETE /medias DeleteAllHandler
//
// Delete all medias
func (m *Service) DeleteAllHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckPermissions(r, nil, "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

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
	if !auth.CheckPermissions(r, nil) {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	vars := mux.Vars(r)
	eltName := vars["eltName"]
	instanceID := vars["instanceId"]
	filePath := vars["filePath"]

	if filePath == "" {
		filePath = "index.html"
	}

	if media := m.getMediaFromRequest(w, r); media != nil {
		pluginDirectoryPath := m.manager.GetPluginDirectory(media, eltName, instanceID)
		pluginFilePath := filepath.Join(pluginDirectoryPath, "frontend", filePath)
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
