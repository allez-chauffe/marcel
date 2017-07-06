package medias

import (
	"encoding/json"
	"errors"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"log"
	"github.com/Zenika/MARCEL/backend/commons"
	"os"
	"fmt"
)

const MEDIAS_CONFIG_PATH string = "data"
const MEDIAS_CONFIG_FILENAME string = "medias.config.json"

var mediasConfigFullpath string = fmt.Sprintf("%s%c%s", MEDIAS_CONFIG_PATH, os.PathSeparator, MEDIAS_CONFIG_FILENAME)

//List of Loaded Medias
var Medias []Media

// LoadMedias loads medias configuration from DB and stor it in memory
func LoadMedias() {
	log.Printf("Start Loading Medias from DB.")

	CreateSaveFileIfNotExist(MEDIAS_CONFIG_PATH, MEDIAS_CONFIG_FILENAME)

	//Medias configurations are loaded from a JSON file on the FS.
	content, err := ioutil.ReadFile(mediasConfigFullpath)
	check(err)

	var obj []interface{}
	json.Unmarshal([]byte(content), &obj)

	//Map the json to the Media structure
	for _, b := range obj {
		var media Media
		err = mapstructure.Decode(b.(map[string]interface{}), &media)
		if err != nil {
			panic(err)
		}
		Medias = append(Medias, media)
	}

	log.Print("Medias configurations is loaded...")
}

// GetMedia Return the media with this id
func GetMedia(idMedia string) (*Media, error) {

	log.Println("Getting media with id: ", idMedia)
	for _, media := range Medias {
		if idMedia == media.ID {
			return &media, nil
		}
	}

	return nil, errors.New("NO_MEDIA_FOUND")
}

// CreateMedia Create a new Media, save it into memory and commit
func CreateMedia() (*Media) {

	log.Println("Creating media")
	newMedia := new(Media)
	newMedia.ID = commons.GetUID()

	SaveMedia(newMedia)
	Commit()

	return newMedia
}

// RemoveMedia Remove media from memory and commit
func RemoveMedia(media *Media) {
	log.Println("Removing media")
	i := GetMediaPosition(media)

	if i >= 0 {
		Medias = append(Medias[:i], Medias[i+1:]...)
	}

	Commit()
}

// GetMediaPosition Return position of a media in the list
func GetMediaPosition(media *Media) int {
	for p, m := range Medias {
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
	Medias = append(Medias, *media)

	Commit()
}

// Commit Save all medias in DB.
// Here DB is a JSON file
func Commit() {
	content, _ := json.Marshal(Medias)

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
