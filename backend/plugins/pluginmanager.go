package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gopkg.in/src-d/go-billy.v4/osfs"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/storage/memory"

	"github.com/Zenika/MARCEL/backend/commons"
	"github.com/Zenika/MARCEL/backend/config"
)

const (
	ErrPluginNotFound errPluginNotFound = "NO_PLUGIN_FOUND"
)

type errPluginNotFound string

func (err errPluginNotFound) Error() string {
	return string(err)
}

type Manager struct {
	ConfigPath     string
	ConfigFileName string
	ConfigFullpath string
	Config         *Configuration
}

func NewManager(configPath, configFilename string) *Manager {
	manager := new(Manager)

	manager.ConfigPath = configPath
	manager.ConfigFileName = configFilename

	manager.ConfigFullpath = filepath.Join(configPath, configFilename)
	manager.Config = NewConfiguration()

	return manager
}

// LoadPlugins loads plugins configuration from DB and store it in memory
func (m *Manager) LoadFromDB() {
	log.Debugln("Start Loading Plugins from DB.")

	commons.LoadFromDB(m)

	log.Debugln("Plugins configurations is loaded...")
}

func (m *Manager) GetConfiguration() *Configuration {
	log.Debugln("Getting global plugins config")

	return m.Config
}

func (m *Manager) GetConfig() interface{} {
	return m.Config
}

func (m *Manager) GetAll() []Plugin {
	log.Debugln("Getting all plugins")

	return m.Config.Plugins
}

// GetPlugin Return the plugin with its eltName
func (m *Manager) Get(eltName string) (*Plugin, error) {

	log.Debugln("Getting plugin with eltName: ", eltName)
	for _, plugin := range m.Config.Plugins {
		if eltName == plugin.EltName {
			return &plugin, nil
		}
	}

	return nil, ErrPluginNotFound
}

// RemovePlugin Remove plugin from memory and commit
func (m *Manager) RemoveFromDB(plugin *Plugin) {
	log.Debugln("Removing plugin")
	i := m.GetPosition(plugin)

	if i >= 0 {
		m.Config.Plugins = append(m.Config.Plugins[:i], m.Config.Plugins[i+1:]...)
	}

	m.Commit()
}

// Save plugin information.
func (m *Manager) Add(plugin *Plugin) {
	log.Debugln("Saving plugin")
	m.RemoveFromDB(plugin)
	m.Config.Plugins = append(m.Config.Plugins, *plugin)

	m.Commit()
}

// Commit Save all plugins in DB.
// Here DB is a JSON file
func (m *Manager) Commit() error {
	return commons.Commit(m)
}

// GetPluginPosition Return position of a plugin in the list
func (m *Manager) GetPosition(plugin *Plugin) int {
	for p, m := range m.Config.Plugins {
		if m.EltName == plugin.EltName {
			return p
		}
	}
	return -1
}

func (m *Manager) GetSaveFilePath() (string, string, string) {
	return m.ConfigFullpath, m.ConfigPath, m.ConfigFileName
}

func (m *Manager) Exists(eltName string) (bool, error) {
	switch _, err := m.Get(eltName); err {
	case nil:
		return true, nil
	case ErrPluginNotFound:
		return false, nil
	default:
		return false, err
	}
}

func (p *Plugin) GetDirectory() string {
	return filepath.Join(config.Config.PluginsPath, p.EltName)
}

func (m *Manager) FetchFromGit(url string) (plugin *Plugin, tempDir string, err error) {
	log.Debugf("Cloning %s...", url)

	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL:          url,
		SingleBranch: true,
		NoCheckout:   true,
		Depth:        1,
		Tags:         git.NoTags,
	})
	if err != nil {
		return nil, "", fmt.Errorf("Error while cloning %s: %s", url, err)
	}

	remote, err := repo.Remote("origin")
	if err != nil {
		return nil, "", fmt.Errorf("Error retrieving origin remote from %s: %s", url, err)
	}

	log.Debug("Fetching tags...")
	refs, err := remote.List(&git.ListOptions{})
	if err != nil {
		return nil, "", fmt.Errorf("Error fetching tags from %s: %s", url, err)
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
		log.Warnf("No tags were found on %s. Using default reference (%s)", url, ref.Short())
	} else {
		return nil, "", fmt.Errorf("The repository %s has no tags and no master branch.")
	}

	tempDir, err = ioutil.TempDir(config.Plugins.Path, "new_plugin")
	if err != nil {
		return nil, "", fmt.Errorf("Error while trying to create temporary directory : %s", err)
	}

	log.Debugf("Cloning %s into %s ...", ref.Short(), tempDir)
	repo, err = git.Clone(memory.NewStorage(), osfs.New(tempDir), &git.CloneOptions{
		URL:           url,
		SingleBranch:  true,
		NoCheckout:    true,
		Depth:         1,
		Tags:          git.NoTags,
		ReferenceName: ref,
	})
	if err != nil {
		return nil, tempDir, fmt.Errorf("Error while cloning %s into %s : %s", ref.Short(), tempDir, err)
	}

	log.Debug("Checking out manifest...")

	wt, err := repo.Worktree()
	if err != nil {
		return nil, tempDir, fmt.Errorf("Error while getting WorkTree of %s: %s", url, err)
	}

	if err = wt.Checkout(&git.CheckoutOptions{Branch: ref}); err != nil {
		return nil, tempDir, fmt.Errorf("Error while checking out manifest of %s: %s", url, err)
	}

	manifest, err := wt.Filesystem.Open("marcel.json")
	if err != nil {
		return nil, tempDir, fmt.Errorf("Error while opening manifest of %s: %s", url, err)
	}
	defer manifest.Close()

	plugin = &Plugin{}
	if err := json.NewDecoder(manifest).Decode(plugin); err != nil {
		return nil, tempDir, fmt.Errorf("Error while reading manifest of %s: %s", url, err)
	}

	plugin.URL = url
	for _, tag := range tags {
		plugin.Versions = append(plugin.Versions, tag.Short())
	}

	return plugin, tempDir, nil
}
