package medias

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// swagger:route GET /medias/{idMedia} getMediaByID
//
// Gets information of a media
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
// swagger:parameters idMedia
func GetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idMedia := vars["idMedia"]

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

// swagger:route POST /medias/{idMedia} setMedias
//
// Posts information for a media
//
//     Consumes:
//     - application/json
//
//     Schemes: http, https
// swagger:parameters idMedia
func PostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idMedia := vars["idMedia"]

	_, err := GetMedia(idMedia)
	if err != nil {
		writeResponseWithError(w, http.StatusNotFound)
		return
	}

	//todo : save
}

// swagger:route GET /medias getMedias
//
// Gets information of all medias
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Medias)
	if err != nil {
		writeResponseWithError(w, http.StatusNotFound)
		return
	}

	w.Write([]byte(b))
}

// swagger:route GET /medias createMedia
//
// Gets information of all medias
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	//get a new media
	newMedia := CreateMedia()

	//return it to the client
	b, err := json.Marshal(*newMedia)
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
