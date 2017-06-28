package medias

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleGetMedia(w http.ResponseWriter, r *http.Request) {
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

func HandleGetAll(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Medias)
	if err != nil {
		writeResponseWithError(w, http.StatusNotFound)
		return
	}

	w.Write([]byte(b))
}

func HandleCreate(w http.ResponseWriter, r *http.Request) {
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
