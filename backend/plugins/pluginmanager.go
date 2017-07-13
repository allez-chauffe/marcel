package plugins

import (
	"encoding/json"
	"errors"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"log"
	"os"
	"fmt"
	"github.com/Zenika/MARCEL/backend/commons"
)

type Manager struct {
	configPath     string
	configFileName string
	configFullpath string
	Config         *Configuration
}

func NewManager(configPath, configFilename string, configuration *Configuration) *Manager {
	manager := new(Manager)

	manager.Config = configuration
	manager.configPath = configPath
	manager.configFileName = configFilename

	manager.configFullpath = fmt.Sprintf("%s%c%s", configPath, os.PathSeparator, configFilename)
	manager.Config = NewConfiguration()

	return manager
}

// LoadPlugins loads plugins configuration from DB and store it in memory
func (m *Manager) Load() {
	log.Printf("Start Loading Plugins from DB.")

	m.CreateSaveFileIfNotExist(m.configPath, m.configFileName)

	//Plugins configurations are loaded from a JSON file on the FS.
	content, err := ioutil.ReadFile(m.configFullpath)
	commons.Check(err)

	var obj interface{}
	json.Unmarshal([]byte(content), &obj)
	err = mapstructure.Decode(obj.(map[string]interface{}), m.Config)
	if err != nil {
		panic(err)
	}

	log.Print("Plugins configurations is loaded...")
}

func (m *Manager) GetConfiguration() (*Configuration) {
	log.Println("Getting global plugins config")

	return m.Config
}

func (m *Manager) GetAll() ([]Plugin) {
	log.Println("Getting all plugins")

	return m.Config.Plugins
}

// GetPlugin Return the plugin with this id
func (m *Manager) Get(idPlugin int) (*Plugin, error) {

	log.Println("Getting plugin with id: ", idPlugin)
	for _, plugin := range m.Config.Plugins {
		if idPlugin == plugin.ID {
			return &plugin, nil
		}
	}

	return nil, errors.New("NO_MEDIA_FOUND")
}

// CreatePlugin Create a new Plugin, save it into memory and commit
func (m *Manager) Create() (*Plugin) {

	log.Println("Creating plugin")

	m.Config.LastID = m.Config.LastID + 1

	newPlugin := NewPlugin()
	newPlugin.ID = m.Config.LastID //commons.GetUID()

	//save it into the Configuration
	m.Save(newPlugin)

	return newPlugin
}

// RemovePlugin Remove plugin from memory and commit
func (m *Manager) Remove(plugin *Plugin) {
	log.Println("Removing plugin")
	i := m.GetPosition(plugin)

	if i >= 0 {
		m.Config.Plugins = append(m.Config.Plugins[:i], m.Config.Plugins[i+1:]...)
	}

	m.Commit()
}

// GetPluginPosition Return position of a plugin in the list
func (m *Manager) GetPosition(plugin *Plugin) int {
	for p, m := range m.Config.Plugins {
		if m.ID == plugin.ID {
			return p
		}
	}
	return -1
}

// SavePlugin Save plugin information in memory.
func (m *Manager) Save(plugin *Plugin) {
	log.Println("Saving plugin")
	m.Remove(plugin)
	m.Config.Plugins = append(m.Config.Plugins, *plugin)

	m.Commit()
}

// Commit Save all plugins in DB.
// Here DB is a JSON file
func (m *Manager) Commit() {
	content, _ := json.Marshal(m.Config)

	err := ioutil.WriteFile(m.configFullpath, content, 0644)

	if err != nil {
		log.Println("Cannot save plugins configuration:")
		log.Panic(err)
	}
}

// CreateSaveFileIfNotExist check if the save file for plugins exists and create it if not.
func (m *Manager) CreateSaveFileIfNotExist(filePath string, fileName string) {

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Println("Data directory did not exist. Create it.")
		os.Mkdir(filePath, 0755)
	}

	var fullPath string = fmt.Sprintf("%s%c%s", filePath, os.PathSeparator, fileName)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {

		f, err := os.Create(fullPath)
		commons.Check(err)

		//content := "[\n]"
		//f.WriteString(content)

		log.Println("Plugins configuration file created at %v", fullPath)

		f.Close()

		//commit a first time to ensure the configuration has been saved
		m.Commit()
	}
}
