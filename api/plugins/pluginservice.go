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

func Initialize() error {
	return initializePluginsDir()
}

func initializePluginsDir() error {
	path, err := filepath.Abs(config.Default().API().PluginsDir())
	if err != nil {
		return fmt.Errorf("Parse plugins directory path %#v: %w", path, err)
	}

	if stat, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(config.Default().API().PluginsDir(), os.ModePerm); err != nil {
				return fmt.Errorf("Create plugins directory %#v: %w", path, err)
			}
		} else {
			return fmt.Errorf("Stat plugins directory %#v: %w", path, err)
		}
	} else if !stat.IsDir() {
		return fmt.Errorf("%#v is not a directory", path)
	}

	log.Debugf("Using plugins directory %s", path)

	return nil
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
	id := vars["id"]

	log.Debugf("Plugin deletion requested: %s", id)

	plugin, err := plugins.Get(id)
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
			log.Warnf("The %s plugin's folder doesn't exists. Ignoring it.", plugin.ID)
		} else {
			log.Errorf("Error while removing %s plugin's folder %s: %s", plugin.ID, plugin.GetDirectory(), err.Error())
			commons.WriteResponse(w, http.StatusInternalServerError, "Error while removing plugin's files")
			return
		}
	}

	if err := plugins.Delete(id); err != nil {
		log.Errorf("Error while removing %s plugin from database: %s", plugin.ID, err.Error())
		commons.WriteResponse(w, http.StatusInternalServerError, "Error while removing plugin from database")
		return
	}

	w.WriteHeader(http.StatusNoContent)

	log.Infof("Plugin deleted : %s", plugin.ID)
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

	log.Infof("Adding plugin %#v", body.URL)

	plugin, tempDir, err := fetchFromGit(body.URL)
	defer os.RemoveAll(tempDir)
	if err != nil {
		log.Error(err)
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
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

	log.Infof("Added plugin %#v", plugin.URL)
	commons.WriteJsonResponse(w, plugin)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckPermissions(r, nil, "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	plugin, err := plugins.Get(id)
	if err != nil {
		log.Error(err)
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	if plugin == nil {
		commons.WriteResponse(w, http.StatusNotFound, "")
		return
	}

	log.Infof("Plugin update requested for %s", id)

	plugin, tempDir, err := fetchFromGit(plugin.URL)
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

	log.Infof("Plugin successfuly updated: %s", plugin.URL)
	commons.WriteJsonResponse(w, plugin)
}
