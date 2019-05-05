package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/Zenika/MARCEL/backend/auth/middleware"
	"github.com/Zenika/MARCEL/backend/commons"
	"github.com/Zenika/MARCEL/backend/config"
	"gopkg.in/src-d/go-billy.v4/osfs"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/storage/memory"
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

	commons.WriteJsonResponse(w, s.Manager.GetConfiguration())
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

	log.Debugf("Cloning %s...", body.URL)

	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL:          body.URL,
		SingleBranch: true,
		NoCheckout:   true,
		Depth:        1,
		Tags:         git.NoTags,
	})
	if err != nil {
		log.Errorf("Error while cloning %s: %s", body.URL, err)
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	remote, err := repo.Remote("origin")
	if err != nil {
		log.Errorf("Error retrieving origin remote from %s: %s", body.URL, err)
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Debug("Fetching tags...")
	refs, err := remote.List(&git.ListOptions{})
	if err != nil {
		log.Errorf("Error fetching tags from %s: %s", body.URL, err)
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	var tags []plumbing.ReferenceName
	var master plumbing.ReferenceName
	for _, ref := range refs {
		name := ref.Name()
		if name.IsTag() {
			tags = append(tags, name)
		}

		if name.IsBranch() && name.Short() == "master" {
			master = name
		}
	}

	var ref plumbing.ReferenceName
	if len(tags) != 0 {
		ref = tags[0]
	} else if master != "" {
		ref = master
		log.Warnf("No tags were found on %s. Using default reference (%s)", body.URL, ref.Short())
	} else {
		err := "The repository %s has no tags and no master branch."
		log.Errorf(err, body.URL, err)
		commons.WriteResponse(w, http.StatusInternalServerError, err)
		return
	}

	tempDir, err := ioutil.TempDir(config.Config.PluginsPath, "new_plugin")
	if err != nil {
		log.Errorf("Error while trying to create temporary directory : %s", err)
		commons.WriteResponse(w, http.StatusInternalServerError, "Error while trying to creat temprorary directory")
		return
	}

	log.Debugf("Cloning %s into %s ...", ref.Short(), tempDir)
	repo, err = git.Clone(memory.NewStorage(), osfs.New(tempDir), &git.CloneOptions{
		URL:           body.URL,
		SingleBranch:  true,
		NoCheckout:    true,
		Depth:         1,
		Tags:          git.NoTags,
		ReferenceName: ref,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Debug("Checking out manifest...")

	wt, err := repo.Worktree()
	if err != nil {
		log.Errorf("Error while getting WorkTree of %s: %s", body.URL, err)
		commons.WriteResponse(w, http.StatusInternalServerError, "Error while checking out manifest : Unable to open WorkTree")
		return
	}

	if err = wt.Checkout(&git.CheckoutOptions{Branch: ref}); err != nil {
		log.Errorf("Error while checking out manifest of %s: %s", body.URL, err)
		commons.WriteResponse(w, http.StatusInternalServerError, fmt.Sprintf("Error while checking out manifest : %s", err))
		return
	}

	manifest, err := wt.Filesystem.Open("marcel.json")
	if err != nil {
		log.Errorf("Error while opening manifest of %s: %s", body.URL, err)
		commons.WriteResponse(w, http.StatusInternalServerError, fmt.Sprintf("Error while opening manifest : %s", err))
		return
	}
	defer manifest.Close()

	plugin := &Plugin{}
	if err := json.NewDecoder(manifest).Decode(plugin); err != nil {
		log.Errorf("Error while reading manifest of %s: %s", body.URL, err)
		commons.WriteResponse(w, http.StatusBadRequest, fmt.Sprintf("Error while reading manifest : %s", err))
		return
	}

	exists, err := s.Manager.Exists(plugin.EltName)
	if err != nil {
		log.Errorf("Error while fetching plugin from database: %s", err)
		commons.WriteResponse(w, http.StatusInternalServerError, "Error while reading plugins database.")
		return
	}

	if exists {
		log.Errorf("The plugin '%s' already exists.", plugin.EltName)
		commons.WriteResponse(w, http.StatusBadRequest, fmt.Sprintf("A plugin with '%s' element name already exists", plugin.EltName))
		return
	}

	log.Debugf("Moving temporary directory (%s) to plugin's folder (%s)", tempDir, plugin.GetDirectory())
	os.Rename(tempDir, plugin.GetDirectory())

	plugin.URL = body.URL
	for _, tag := range tags {
		plugin.Versions = append(plugin.Versions, tag.Short())
	}
	s.Manager.Add(plugin)

	log.Infof("Plugin successfuly registered : %s (%s)", plugin.EltName, plugin.Name)

	commons.WriteJsonResponse(w, plugin)
}
