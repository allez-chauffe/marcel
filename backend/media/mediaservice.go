package media

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
	"log"
	"github.com/gorilla/mux"
	"strconv"
	"github.com/mitchellh/mapstructure"
)

var Medias []Media
var y []interface{}

func LoadMedias() {
	//Medias configurations are loaded from a JSON file on the FS.
	content, err := ioutil.ReadFile("data/media.config.json")
	check(err)

	json.Unmarshal([]byte(content), &y)

	//Map the json to the Media structure
	var m Media
	err = mapstructure.Decode(y[0].(map[string]interface{}), &m)
	if err != nil {
		panic(err)
	}

	Medias = append(Medias, m)

	log.Print("Medias configurations is loaded...")
}

func GetMedia(idMedia int) (*Media, error) {
	for _, m := range Medias {
		if idMedia == m.ID {
			return &m, nil
		}
	}

	return nil, errors.New("NO_MEDIA_FOUND")
}

func HandleGetMedia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	f := vars["idMedia"]
	idMedia, _ := strconv.Atoi(f)

	m, err := GetMedia(idMedia)
	if err != nil {
		writeResponseWithError(w, http.StatusNotFound)
		return
	}

	b, err := json.Marshal(*m)
	if err != nil {
		writeResponseWithError(w, http.StatusNotFound)
		return
	}

	w.Write([]byte(b))
}

func writeResponseWithError(w http.ResponseWriter, errorCode int) {
	w.WriteHeader(errorCode)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}
