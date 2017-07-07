package medias

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/Zenika/MARCEL/backend/commons"
	"io/ioutil"
)

type MediaService struct {
	mediaManager *MediaManager
}

func NewMediaService() MediaService {
	mediaService := MediaService{}

	mediaService.mediaManager = NewMediaManager()

	return mediaService
}

func (m *MediaService) GetMediaManager() (*MediaManager) {
	return m.mediaManager
}

// swagger:route GET /medias/config GetConfigHandler
//
// Gets information of all medias
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
func (m *MediaService) GetConfigHandler(w http.ResponseWriter, r *http.Request) {

	c := m.mediaManager.GetMediasConfiguration()
	b, err := json.Marshal(c)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotFound)
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
func (m *MediaService) GetAllHandler(w http.ResponseWriter, r *http.Request) {

	media := m.mediaManager.GetMedias()
	b, err := json.Marshal(media)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotFound)
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
func (m *MediaService) GetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	attr := vars["idMedia"]

	idMedia, err := strconv.Atoi(attr)
	if err != nil {
		commons.WriteResponse(w, http.StatusBadRequest)
		return
	}

	media, err := m.mediaManager.GetMedia(idMedia)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotFound)
		return
	}

	b, err := json.Marshal(*media)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotFound)
		return
	}

	w.Write([]byte(b))
}

// swagger:route POST /medias PostHandler
//
// Posts information for a media
//
//     Consumes:
//     - application/json
//
//     Schemes: http, https
func (m *MediaService) PostHandler(w http.ResponseWriter, r *http.Request) {
	//to be tested : decoder := json.NewDecoder(r.Body)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		commons.WriteResponse(w, http.StatusBadRequest)
	}

	var media *Media = NewMedia()
	err = json.Unmarshal(body, &media)

	m.mediaManager.SaveMedia(media)
}

// swagger:route GET /medias CreateHandler
//
// Gets information of all medias
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
func (m *MediaService) CreateHandler(w http.ResponseWriter, r *http.Request) {
	//get a new media
	newMedia := m.mediaManager.CreateMedia()

	//return it to the client
	b, err := json.Marshal(*newMedia)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotFound)
		return
	}

	w.Write([]byte(b))
}

