package medias

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

//List of Loaded Medias
//var Medias []Media

// LoadMedias loads medias configuration from DB and stor it in memory
func (m *Manager) Load() {
	log.Printf("Start Loading Medias from DB.")

	m.CreateSaveFileIfNotExist(m.configPath, m.configFileName)

	//Medias configurations are loaded from a JSON file on the FS.
	content, err := ioutil.ReadFile(m.configFullpath)
	commons.Check(err)

	var obj interface{}
	json.Unmarshal([]byte(content), &obj)
	err = mapstructure.Decode(obj.(map[string]interface{}), m.Config)
	if err != nil {
		panic(err)
	}

	log.Print("Medias configurations is loaded...")
}

func (m *Manager) GetConfiguration() (*Configuration) {
	log.Println("Getting global medias config")

	return m.Config
}

func (m *Manager) GetAll() ([]Media) {
	log.Println("Getting all medias")

	return m.Config.Medias
}

// GetMedia Return the media with this id
func (m *Manager) Get(idMedia int) (*Media, error) {

	log.Println("Getting media with id: ", idMedia)
	for _, media := range m.Config.Medias {
		if idMedia == media.ID {
			return &media, nil
		}
	}

	return nil, errors.New("NO_MEDIA_FOUND")
}

// CreateMedia Create a new Media, save it into memory and commit
func (m *Manager) Create() (*Media) {

	log.Println("Creating media")

	m.Config.LastID = m.Config.LastID + 1

	newMedia := NewMedia()
	newMedia.ID = m.Config.LastID //commons.GetUID()

	//save it into the MediasConfiguration
	m.Save(newMedia)

	return newMedia
}

// RemoveMedia Remove media from memory and commit
func (m *Manager) Remove(media *Media) {
	log.Println("Removing media")
	i := m.GetPosition(media)

	if i >= 0 {
		m.Config.Medias = append(m.Config.Medias[:i], m.Config.Medias[i+1:]...)
	}

	m.Commit()
}

// GetMediaPosition Return position of a media in the list
func (m *Manager) GetPosition(media *Media) int {
	for p, m := range m.Config.Medias {
		if m.ID == media.ID {
			return p
		}
	}
	return -1
}

// SaveMedia Save media information in memory.
func (m *Manager) Save(media *Media) {
	log.Println("Saving media")
	m.Remove(media)
	m.Config.Medias = append(m.Config.Medias, *media)

	m.Commit()
}

// Commit Save all medias in DB.
// Here DB is a JSON file
func (m *Manager) Commit() {
	content, _ := json.Marshal(m.Config)

	err := ioutil.WriteFile(m.configFullpath, content, 0644)

	if err != nil {
		log.Println("Cannot save medias configuration:")
		log.Panic(err)
	}
}

// CreateSaveFileIfNotExist check if the save file for medias exists and create it if not.
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

		log.Println("Medias configuration file created at %v", fullPath)

		f.Close()

		//commit a first time to ensure the configuration has been saved
		m.Commit()
	}
}
