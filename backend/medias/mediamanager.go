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

const MEDIAS_CONFIG_PATH string = "data"
const MEDIAS_CONFIG_FILENAME string = "medias.config.json"

type MediaManager struct {
	mediasConfigFullpath string
	MediasConfig         *MediasConfiguration
}

func NewMediaManager() *MediaManager {
	mediaManager := new(MediaManager)

	mediaManager.mediasConfigFullpath = fmt.Sprintf("%s%c%s", MEDIAS_CONFIG_PATH, os.PathSeparator, MEDIAS_CONFIG_FILENAME)
	mediaManager.MediasConfig = new(MediasConfiguration)

	return mediaManager
}

//List of Loaded Medias
//var Medias []Media

// LoadMedias loads medias configuration from DB and stor it in memory
func (m *MediaManager) LoadMedias() {
	log.Printf("Start Loading Medias from DB.")

	m.CreateSaveFileIfNotExist(MEDIAS_CONFIG_PATH, MEDIAS_CONFIG_FILENAME)

	//Medias configurations are loaded from a JSON file on the FS.
	content, err := ioutil.ReadFile(m.mediasConfigFullpath)
	commons.Check(err)

	var obj interface{}
	json.Unmarshal([]byte(content), &obj)
	err = mapstructure.Decode(obj.(map[string]interface{}), m.MediasConfig)
	if err != nil {
		panic(err)
	}

	log.Print("Medias configurations is loaded...")
}

func (m *MediaManager) GetMediasConfiguration() (*MediasConfiguration) {
	log.Println("Getting global medias config")

	return m.MediasConfig
}

func (m *MediaManager) GetMedias() ([]Media) {
	log.Println("Getting all medias")

	return m.MediasConfig.Medias
}

// GetMedia Return the media with this id
func (m *MediaManager) GetMedia(idMedia int) (*Media, error) {

	log.Println("Getting media with id: ", idMedia)
	for _, media := range m.MediasConfig.Medias {
		if idMedia == media.ID {
			return &media, nil
		}
	}

	return nil, errors.New("NO_MEDIA_FOUND")
}

// CreateMedia Create a new Media, save it into memory and commit
func (m *MediaManager) CreateMedia() (*Media) {

	log.Println("Creating media")

	m.MediasConfig.LastID = m.MediasConfig.LastID + 1

	newMedia := new(Media)
	newMedia.ID = m.MediasConfig.LastID //commons.GetUID()

	m.SaveMedia(newMedia)

	return newMedia
}

// RemoveMedia Remove media from memory and commit
func (m *MediaManager) RemoveMedia(media *Media) {
	log.Println("Removing media")
	i := m.GetMediaPosition(media)

	if i >= 0 {
		m.MediasConfig.Medias = append(m.MediasConfig.Medias[:i], m.MediasConfig.Medias[i+1:]...)
	}

	m.Commit()
}

// GetMediaPosition Return position of a media in the list
func (m *MediaManager) GetMediaPosition(media *Media) int {
	for p, m := range m.MediasConfig.Medias {
		if m.ID == media.ID {
			return p
		}
	}
	return -1
}

// SaveMedia Save media information in memory.
func (m *MediaManager) SaveMedia(media *Media) {
	log.Println("Saving media")
	m.RemoveMedia(media)
	m.MediasConfig.Medias = append(m.MediasConfig.Medias, *media)

	m.Commit()
}

// Commit Save all medias in DB.
// Here DB is a JSON file
func (m *MediaManager) Commit() {
	content, _ := json.Marshal(m.MediasConfig)

	err := ioutil.WriteFile(m.mediasConfigFullpath, content, 0644)

	if err != nil {
		log.Println("Cannot save medias configuration:")
		log.Panic(err)
	}
}

// CreateSaveFileIfNotExist check if the save file for medias exists and create it if not.
func (m *MediaManager) CreateSaveFileIfNotExist(filePath string, fileName string) {

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
