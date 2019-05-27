package plugins

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/Zenika/MARCEL/api/auth"
	"github.com/Zenika/MARCEL/api/commons"
	"github.com/Zenika/MARCEL/api/db/plugins"
)

var (
	pluginsTempDir     string
	initPluginsTempDir sync.Once
)

// swagger:route GET /plugins GetAllHandler
//
// Gets information of all plugins
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckPermissions(r, nil, "user", "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	plugins, err := plugins.List()
	if err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	commons.WriteJsonResponse(w, plugins)
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
func GetHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckPermissions(r, nil, "user", "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	vars := mux.Vars(r)
	eltName := vars["eltName"]

	plugin, err := plugins.Get(eltName)
	if err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	if plugin == nil {
		commons.WriteResponse(w, http.StatusNotFound, "")
		return
	}

	commons.WriteJsonResponse(w, plugin)
}

type AddPluginBody struct {
	URL string `json:"url"`
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckPermissions(r, nil, "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	body := &AddPluginBody{}
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		commons.WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	log.Infof("Plugin registration requested for %s", body.URL)

	plugin, tempDir, err := FetchFromGit(body.URL)
	defer os.RemoveAll(tempDir)
	if err != nil {
		log.Error(err)
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	exists, err := plugins.Exists(plugin.EltName)
	if err != nil {
		log.Error(err)
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	if exists {
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

	if err := plugins.Insert(plugin); err != nil {
		log.Error(err)
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Infof("Plugin successfuly registered : %s (%s)", plugin.EltName, plugin.Name)
	commons.WriteJsonResponse(w, plugin)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckPermissions(r, nil, "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	vars := mux.Vars(r)
	eltName := vars["eltName"]

	plugin, err := plugins.Get(eltName)
	if err != nil {
		log.Error(err)
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	if plugin == nil {
		commons.WriteResponse(w, http.StatusNotFound, "")
		return
	}

	log.Infof("Plugin update requested for %s", eltName)

	plugin, tempDir, err := FetchFromGit(plugin.URL)
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

	if err := plugins.Update(plugin); err != nil {
		log.Error(err)
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Infof("Plugin successfuly updated: %s (%s)", plugin.EltName, plugin.Name)
	commons.WriteJsonResponse(w, plugin)
}
