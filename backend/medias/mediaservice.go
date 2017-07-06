package medias

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)


// swagger:route GET /medias/config GetConfigHandler
//
// Gets information of all medias
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
func GetConfigHandler(w http.ResponseWriter, r *http.Request) {

	c:= GetMediasConfiguration()
	b, err := json.Marshal(c)
	if err != nil {
		writeResponseWithError(w, http.StatusNotFound)
		return
	}

	w.Write([]byte(b))
}

// swagger:route GET /medias GetAllHandler
//
// Gets information of all medias
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
func GetAllHandler(w http.ResponseWriter, r *http.Request) {

	m := GetMedias()
	b, err := json.Marshal(m)
	if err != nil {
		writeResponseWithError(w, http.StatusNotFound)
		return
	}

	w.Write([]byte(b))
}


// swagger:route GET /medias/{idMedia} GetHandler
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
	attr := vars["idMedia"]

	idMedia, err := strconv.Atoi(attr)
	if err != nil {
		writeResponseWithError(w, http.StatusBadRequest)
		return
	}

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

// swagger:route POST /medias/{idMedia} PostHandler
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
	attr := vars["idMedia"]

	idMedia, err := strconv.Atoi(attr)
	if err != nil {
		writeResponseWithError(w, http.StatusBadRequest)
		return
	}

	_, err = GetMedia(idMedia)
	if err != nil {
		writeResponseWithError(w, http.StatusNotFound)
		return
	}

	//todo : save
}

// swagger:route GET /medias CreateHandler
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
