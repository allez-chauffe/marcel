package medias

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/Zenika/marcel/api/auth"
	"github.com/Zenika/marcel/api/clients"
	"github.com/Zenika/marcel/api/commons"
	"github.com/Zenika/marcel/api/db/medias"
	"github.com/Zenika/marcel/config"
)

type Service struct {
	clientsService *clients.Service
}

func NewService(clientsService *clients.Service) *Service {
	service := new(Service)

	service.clientsService = clientsService

	return service
}

// GetAllHandler gets information of all medias
func (m *Service) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckPermissions(r, nil, "user", "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	medias, err := medias.List()
	if err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	commons.WriteJsonResponse(w, medias)
}

// GetHandler gets information of a media
func (m *Service) GetHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckPermissions(r, nil, "user", "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	if media := m.getMediaFromRequest(w, r); media != nil {
		commons.WriteJsonResponse(w, media)
	}
}

// SaveHandler saves a media
func (m *Service) SaveHandler(w http.ResponseWriter, r *http.Request) {
	media := &medias.Media{}
	if err := json.NewDecoder(r.Body).Decode(media); err != nil {
		commons.WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	tmpMedia, err := medias.Get(media.ID)
	if err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	if tmpMedia == nil {
		commons.WriteResponse(w, http.StatusNotFound, "")
		return
	}

	if !auth.CheckPermissions(r, []string{tmpMedia.Owner}, "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	if err := activate(media); err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	media.IsActive = true

	if err := medias.Update(media); err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	commons.WriteJsonResponse(w, media)
	m.clientsService.SendByMedia(media.ID, "update")
}

// CreateHandler creates a new empty media
func (m *Service) CreateHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckPermissions(r, nil) {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	media := medias.New(auth.GetAuth(r).Subject)

	if err := medias.Insert(media); err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	commons.WriteJsonResponse(w, media)
}

// ActivateHandler activates a media
func (m *Service) ActivateHandler(w http.ResponseWriter, r *http.Request) {
	media := m.getMediaFromRequest(w, r)
	if media == nil {
		return
	}

	if !auth.CheckPermissions(r, []string{media.Owner}, "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	if err := activate(media); err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	media.IsActive = true

	if err := medias.Update(media); err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	commons.WriteResponse(w, http.StatusOK, "Media is active")
	m.clientsService.SendByMedia(media.ID, "update")
}

// DeactivateHandler deactivates a media
func (m *Service) DeactivateHandler(w http.ResponseWriter, r *http.Request) {
	media := m.getMediaFromRequest(w, r)
	if media == nil {
		return
	}

	if !auth.CheckPermissions(r, []string{media.Owner}, "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	media.IsActive = false

	if err := medias.Update(media); err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	commons.WriteResponse(w, http.StatusOK, "Media has been deactivated")
	m.clientsService.SendByMedia(media.ID, "update")
}

// DeleteHandler deletes a media
func (m *Service) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	media := m.getMediaFromRequest(w, r)
	if media == nil {
		return
	}

	if !auth.CheckPermissions(r, []string{media.Owner}, "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	if err := medias.Delete(media.ID); err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := os.RemoveAll(filepath.Join(config.Config.API.MediasDir(), strconv.Itoa(media.ID))); err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	commons.WriteResponse(w, http.StatusOK, "Media has been correctly deleted")
}

// GetPluginFilesHandler Serves static frontend files of the given plugin instance for the given media.
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
		pluginDirectoryPath := getPluginDirectory(media, eltName, instanceID)
		pluginFilePath := filepath.Join(pluginDirectoryPath, "frontend", filePath)
		http.ServeFile(w, r, pluginFilePath)
	}
}

func (m *Service) getMediaFromRequest(w http.ResponseWriter, r *http.Request) (media *medias.Media) {
	vars := mux.Vars(r)
	attr := vars["idMedia"]

	idMedia, err := strconv.Atoi(attr)
	if err != nil {
		commons.WriteResponse(w, http.StatusBadRequest, err.Error())
		return nil
	}

	media, err = medias.Get(idMedia)
	if err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return nil
	}
	if media == nil {
		commons.WriteResponse(w, http.StatusNotFound, "")
		return nil
	}

	return media
}

func activate(media *medias.Media) error {
	errorMessages := ""

	for _, mp := range media.Plugins {
		// duplicate plugin files into "medias/{idMedia}/{plugins_EltName}/{idInstance}"
		mpPath := getPluginDirectory(media, mp.EltName, mp.InstanceID)
		if err := copyNewInstanceOfPlugin(media, &mp, mpPath); err != nil {
			log.Errorln(err.Error())
			//Don't return an error now, we need to activate the other plugins
			errorMessages += err.Error() + "\n"
		}
	}

	if errorMessages != "" {
		return errors.New(errorMessages)
	}

	return nil
}

func copyNewInstanceOfPlugin(media *medias.Media, mp *medias.MediaPlugin, path string) error {
	return commons.CopyDir(filepath.Join(config.Config.API.PluginsDir(), mp.EltName, "frontend"), filepath.Join(path, "frontend"))
}

func getPluginDirectory(media *medias.Media, eltName string, instanceID string) string {
	return filepath.Join(config.Config.API.MediasDir(), strconv.Itoa(media.ID), eltName, instanceID)
}
