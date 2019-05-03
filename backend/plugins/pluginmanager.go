package plugins

import (
	"path/filepath"

	log "github.com/sirupsen/logrus"

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
	return filepath.Join(config.Plugins.Path, p.EltName)
}
