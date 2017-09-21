package plugins

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/Zenika/MARCEL/backend/commons"
)

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

	manager.ConfigFullpath = fmt.Sprintf("%s%c%s", configPath, os.PathSeparator, configFilename)
	manager.Config = NewConfiguration()

	return manager
}

// LoadPlugins loads plugins configuration from DB and store it in memory
func (m *Manager) LoadFromDB() {
	log.Printf("Start Loading Plugins from DB.")

	commons.LoadFromDB(m)

	log.Print("Plugins configurations is loaded...")
}

func (m *Manager) GetConfiguration() *Configuration {
	log.Println("Getting global plugins config")

	return m.Config
}

func (m *Manager) GetConfig() interface{} {
	return m.Config
}

func (m *Manager) GetAll() []Plugin {
	log.Println("Getting all plugins")

	return m.Config.Plugins
}

// GetPlugin Return the plugin with its eltName
func (m *Manager) Get(eltName string) (*Plugin, error) {

	log.Println("Getting plugin with eltName: ", eltName)
	for _, plugin := range m.Config.Plugins {
		if eltName == plugin.EltName {
			return &plugin, nil
		}
	}

	return nil, errors.New("NO_MEDIA_FOUND")
}

// RemovePlugin Remove plugin from memory and commit
func (m *Manager) RemoveFromDB(plugin *Plugin) {
	log.Println("Removing plugin")
	i := m.GetPosition(plugin)

	if i >= 0 {
		m.Config.Plugins = append(m.Config.Plugins[:i], m.Config.Plugins[i+1:]...)
	}

	m.Commit()
}

// Save plugin information.
func (m *Manager) Add(plugin *Plugin) {
	log.Println("Saving plugin")
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
