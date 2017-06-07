package medias

import (
	"io/ioutil"
	"encoding/json"
	"errors"
	"log"
	"github.com/mitchellh/mapstructure"
)

var Medias []Media

const path_to_config string = "data/medias.config.json"

func LoadMedias() {
	//Medias configurations are loaded from a JSON file on the FS.
	content, err := ioutil.ReadFile(path_to_config)
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

func GetMedia(idMedia int) (*Media, error) {
	for _, media := range Medias {
		if idMedia == media.ID {
			return &media, nil
		}
	}

	return nil, errors.New("NO_MEDIA_FOUND")
}
