package plugins

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/Zenika/MARCEL/backend/auth/middleware"
	"github.com/Zenika/MARCEL/backend/commons"
	"github.com/Zenika/MARCEL/backend/config"
)

var (
	pluginsTempDir     string
	initPluginsTempDir sync.Once
)

type Service struct {
	Manager *Manager
}

func NewService() *Service {
	var p = new(Service)

	p.Manager = NewManager(config.Config.DataPath, config.Config.PluginsFile)

	return p
}

func (s *Service) GetManager() *Manager {
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
	if !middleware.CheckPermissions(r, nil) {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	commons.WriteJsonResponse(w, s.Manager.GetConfig())
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
	if !middleware.CheckPermissions(r, nil) {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	commons.WriteJsonResponse(w, m.Manager.GetAll())
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
	if !middleware.CheckPermissions(r, nil) {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	vars := mux.Vars(r)
	eltName := vars["eltName"]

	plugin, err := s.Manager.Get(eltName)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotFound, err.Error())
		return
	}

	commons.WriteJsonResponse(w, plugin)
}

type AddPluginBody struct {
	URL string `json:"url"`
}

func (s *Service) AddHandler(w http.ResponseWriter, r *http.Request) {
	// if !middleware.CheckPermissions(r, nil, "admin") {
	// 	commons.WriteResponse(w, http.StatusForbidden, "")
	// 	return
	// }

	body := &AddPluginBody{}
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		commons.WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	log.Infof("Plugin registration requested for %s", body.URL)

	plugin, tempDir, err := s.Manager.FetchFromGit(body.URL)
	defer os.RemoveAll(tempDir)
	if err != nil {
		log.Error(err)
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if s.Manager.Exists(plugin.EltName) {
		log.Errorf("The plugin '%s' already exists", plugin.EltName)
		commons.WriteResponse(w, http.StatusBadRequest, fmt.Sprintf("The plugin '%s' already exists", plugin.EltName))
		return
	}

	log.Debugf("Moving temporary directory (%s) to plugin's folder (%s)", tempDir, plugin.GetDirectory())
	if err := os.Rename(tempDir, plugin.GetDirectory()); err != nil {
		log.Errorf("Error while moving temporary directory (%s) to plugin's folder (%s) : %s", tempDir, plugin.GetDirectory(), err)
		commons.WriteResponse(w, http.StatusInternalServerError, "Error while saving plugin's files")
		return
	}

	s.Manager.Add(plugin)

	log.Infof("Plugin successfuly registered : %s (%s)", plugin.EltName, plugin.Name)
	commons.WriteJsonResponse(w, plugin)
}

func (s *Service) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	// if !middleware.CheckPermissions(r, nil, "admin") {
	// 	commons.WriteResponse(w, http.StatusForbidden, "")
	// 	return
	// }

	vars := mux.Vars(r)
	eltName := vars["eltName"]

	plugin, err := s.Manager.Get(eltName)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotFound, err.Error())
		return
	}

	log.Infof("Plugin update requested for %s", eltName)

	plugin, tempDir, err := s.Manager.FetchFromGit(plugin.URL)
	// The temp dir cleanup should be done before handling because it can be created even if an error occured
	defer os.RemoveAll(tempDir)
	if err != nil {
		log.Error(err)
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Debugf("Removing old plugin's directory (%s)", plugin.GetDirectory())
	if err := os.RemoveAll(plugin.GetDirectory()); err != nil {
		log.Errorf("Error while removing old plugin directory : %s", err)
		commons.WriteResponse(w, http.StatusInternalServerError, "Error while updating plugin's files")
		return
	}

	log.Debugf("Moving temporary directory (%s) to plugin's directory (%s)", tempDir, plugin.GetDirectory())
	if err := os.Rename(tempDir, plugin.GetDirectory()); err != nil {
		log.Errorf("Error while moving temporary directory (%s) to plugin's directory (%s) : %s", tempDir, plugin.GetDirectory(), err)
		commons.WriteResponse(w, http.StatusInternalServerError, "Error while updating plugin's files")
		return
	}

	s.Manager.Replace(plugin)

	log.Infof("Plugin successfuly updated: %s (%s)", plugin.EltName, plugin.Name)
	commons.WriteJsonResponse(w, plugin)
}
