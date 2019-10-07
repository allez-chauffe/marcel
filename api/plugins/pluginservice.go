package plugins

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/Zenika/marcel/api/auth"
	"github.com/Zenika/marcel/api/commons"
	"github.com/Zenika/marcel/api/db/plugins"
	"github.com/Zenika/marcel/config"
)

// Initialize unsures that the plugins directory exists
func Initialize() {
	pluginsPath, err := filepath.Abs(config.Default().API().PluginsDir())
	if err != nil {
		log.Fatalf("Error while parsing plugins directory path: %s", err)
	}

	if stat, err := os.Stat(pluginsPath); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(config.Default().API().PluginsDir(), os.ModePerm); err != nil {
				log.Fatalf("Error while trying to create plugins directory '%s': %s", pluginsPath, err)
			}

			log.Debugf("Plugins directory '%s' created", pluginsPath)
			return
		}
	} else if !stat.IsDir() {
		log.Fatalf("The plugins path '%s' is not a directory", pluginsPath)
	}

	log.Debugf("Using plugins directory %s", pluginsPath)
}

// GetAllHandler gets information of all plugins
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

// GetHandler gets information of a plugin
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

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckPermissions(r, nil, "user", "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	vars := mux.Vars(r)
	eltName := vars["eltName"]

	log.Debugf("Plugin deletion requested: %s", eltName)

	plugin, err := plugins.Get(eltName)
	if err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	if plugin == nil {
		commons.WriteResponse(w, http.StatusNotFound, "")
		return
	}

	if err := os.RemoveAll(plugin.GetDirectory()); err != nil {
		if os.IsNotExist(err) {
			log.Warnf("The %s plugin's folder doesn't exists. Ignoring it.", plugin.EltName)
		} else {
			log.Errorf("Error while removing %s plugin's folder %s: %s", plugin.EltName, plugin.GetDirectory(), err.Error())
			commons.WriteResponse(w, http.StatusInternalServerError, "Error while removing plugin's files")
			return
		}
	}

	if err := plugins.Delete(eltName); err != nil {
		log.Errorf("Error while removing %s plugin from database: %s", plugin.EltName, err.Error())
		commons.WriteResponse(w, http.StatusInternalServerError, "Error while removing plugin from database")
		return
	}

	w.WriteHeader(http.StatusNoContent)

	log.Infof("Plugin deleted : %s", plugin.EltName)
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
