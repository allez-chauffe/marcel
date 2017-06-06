package media

import (
	"net/http"
	//"github.com/gorilla/mux"
	//"strconv"
	"io/ioutil"
	"encoding/json"
	"errors"
	"log"
	"fmt"
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

	/*err = json.Unmarshal(content, &Medias)
	check(err)
	}*/

	json.Unmarshal([]byte(content), &y)
	fmt.Println("1. ---------------------------------")
	fmt.Printf("Type: %T \n", y[0])
	fmt.Println("2. ---------------------------------")

	//for now: only the first media
	/*mi := y[0].(map[string]interface{})
	log.Println("mi = ")
	log.Println(mi)

	m := &Media{}
	err = m.FillStruct(mi)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)*/


	/*for k, v := range mi {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}*/

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
	//b, err := json.Marshal(y)
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
