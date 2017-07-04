package medias

import (
	"encoding/json"
	"errors"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"log"
	"github.com/Zenika/MARCEL/backend/commons"
)

const MEDIAS_CONFIG_FILE string = "data/medias.config.json"

//List of Loaded Medias
var Medias []Media

func LoadMedias() {
	log.Printf("Start Loading Medias from DB")

	//Medias configurations are loaded from a JSON file on the FS.
	content, err := ioutil.ReadFile(MEDIAS_CONFIG_FILE)
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

func GetMedia(idMedia string) (*Media, error) {
	for _, media := range Medias {
		if idMedia == media.ID {
			return &media, nil
		}
	}

	return nil, errors.New("NO_MEDIA_FOUND")
}

/**
Create a new Media, save it into memory.
 */
func CreateMedia() (*Media) {
	newMedia := new(Media)
	newMedia.ID = commons.GetUID()

	SaveMedia(newMedia)
	Commit()

	return newMedia
}

/**
Remove media from memory.
 */
func RemoveMedia(media *Media) {
	i := GetMediaPosition(media)

	if i >= 0 {
		Medias = append(Medias[:i], Medias[i+1:]...)
	}

	Commit()
}

/**
Return position of a media in the list
 */
func GetMediaPosition(media *Media) int {
	for p, m := range Medias {
		if m.ID == media.ID {
			return p
		}
	}
	return -1
}

/**
Save media information in memory.
 */
func SaveMedia(media *Media) {
	RemoveMedia(media)
	Medias = append(Medias, *media)

	Commit()
}

/**
Save all medias in DB. Here DB is a JSON file
 */
func Commit() {
	content, _ :=json.Marshal(Medias)
	ioutil.WriteFile(MEDIAS_CONFIG_FILE, content, 0644)
}
