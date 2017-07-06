package medias

import (
	"encoding/json"
	"errors"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"log"
	"os"
	"fmt"
)

const MEDIAS_CONFIG_PATH string = "data"
const MEDIAS_CONFIG_FILENAME string = "medias.config.json"

var mediasConfigFullpath string = fmt.Sprintf("%s%c%s", MEDIAS_CONFIG_PATH, os.PathSeparator, MEDIAS_CONFIG_FILENAME)

//List of Loaded Medias
//var Medias []Media
var MediasConfig MediasConfiguration

// LoadMedias loads medias configuration from DB and stor it in memory
func LoadMedias() {
	log.Printf("Start Loading Medias from DB.")

	MediasConfig = MediasConfiguration{}
	CreateSaveFileIfNotExist(MEDIAS_CONFIG_PATH, MEDIAS_CONFIG_FILENAME)

	//Medias configurations are loaded from a JSON file on the FS.
	content, err := ioutil.ReadFile(mediasConfigFullpath)
	check(err)

	var obj interface{}
	json.Unmarshal([]byte(content), &obj)
	err = mapstructure.Decode(obj.(map[string]interface{}), &MediasConfig)
	if err != nil {
		panic(err)
	}

	//MediasConfig.LastID = len(obj)
	//Map the json to the MediasConfig structure
	/*for _, b := range obj {
		var media Media
		err = mapstructure.Decode(b.(map[string]interface{}), &media)
		if err != nil {
			panic(err)
		}
		MediasConfig.Medias = append(MediasConfig.Medias, media)
	}*/

	log.Print("Medias configurations is loaded...")
}

func GetMediasConfiguration() (*MediasConfiguration) {
	log.Println("Getting global medias config")

	return &MediasConfig
}

func GetMedias() ([]Media) {
	log.Println("Getting all medias")

	return MediasConfig.Medias
}

// GetMedia Return the media with this id
func GetMedia(idMedia int) (*Media, error) {

	log.Println("Getting media with id: ", idMedia)
	for _, media := range MediasConfig.Medias {
		if idMedia == media.ID {
			return &media, nil
		}
	}

	return nil, errors.New("NO_MEDIA_FOUND")
}

// CreateMedia Create a new Media, save it into memory and commit
func CreateMedia() (*Media) {

	log.Println("Creating media")


	MediasConfig.LastID = MediasConfig.LastID +1

	newMedia := new(Media)
	newMedia.ID = MediasConfig.LastID //commons.GetUID()

	SaveMedia(newMedia)

	return newMedia
}

// RemoveMedia Remove media from memory and commit
func RemoveMedia(media *Media) {
	log.Println("Removing media")
	i := GetMediaPosition(media)

	if i >= 0 {
		MediasConfig.Medias = append(MediasConfig.Medias[:i], MediasConfig.Medias[i+1:]...)
	}

	Commit()
}

// GetMediaPosition Return position of a media in the list
func GetMediaPosition(media *Media) int {
	for p, m := range MediasConfig.Medias {
		if m.ID == media.ID {
			return p
		}
	}
	return -1
}

// SaveMedia Save media information in memory.
func SaveMedia(media *Media) {
	log.Println("Saving media")
	RemoveMedia(media)
	MediasConfig.Medias = append(MediasConfig.Medias, *media)

	Commit()
}

// Commit Save all medias in DB.
// Here DB is a JSON file
func Commit() {
	content, _ := json.Marshal(MediasConfig)

	err := ioutil.WriteFile(mediasConfigFullpath, content, 0644)

	if err != nil {
		log.Println("Cannot save medias configuration:")
		log.Panic(err)
	}
}

// CreateSaveFileIfNotExist check if the save file for medias exists and create it if not.
func CreateSaveFileIfNotExist(filePath string, fileName string) {

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Println("Data directory did not exist. Create it.")
		os.Mkdir(filePath, 0755)
	}

	var fullPath string = fmt.Sprintf("%s%c%s", filePath, os.PathSeparator, fileName)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		content := "[\n]"

		f, err := os.Create(fullPath)
		check(err)

		f.WriteString(content)

		log.Println("Medias configuration file created at %v", fullPath)

		f.Close()
	}
}
