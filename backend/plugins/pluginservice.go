package plugins

import (
	"net/http"
	"encoding/json"
	"github.com/Zenika/MARCEL/backend/commons"
	"github.com/gorilla/mux"
	"fmt"
	"io"
	"os"
	"log"
	"path"
	"errors"
	"strings"
)

const PLUGINS_CONFIG_PATH string = "data"
const PLUGINS_CONFIG_FILENAME string = "plugins.json"
const PLUGINS_TEMPORARY_FOLDER string = "uploadedfiles"

type Service struct {
	Manager *Manager
}

func NewService() *Service {
	var p = new(Service)

	p.Manager = NewManager(PLUGINS_CONFIG_PATH, PLUGINS_CONFIG_FILENAME)

	return p
}

func (s *Service) GetManager() (*Manager) {
	return s.Manager
}

// swagger:route GET /plugins/config GetConfigHandler
//
// Gets information of all plugins
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
func (s *Service) GetConfigHandler(w http.ResponseWriter, r *http.Request) {

	c := s.Manager.GetConfiguration()
	b, err := json.Marshal(c)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotFound, "Impossible to get configuration of the plugins")
		return
	}

	commons.WriteResponse(w, http.StatusOK, (string)(b))
}

// swagger:route GET /plugins GetAllHandler
//
// Gets information of all plugins
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
func (m *Service) GetAllHandler(w http.ResponseWriter, r *http.Request) {

	media := m.Manager.GetAll()
	b, err := json.Marshal(media)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotFound, "Impossible to get all plugins")
		return
	}

	commons.WriteResponse(w, http.StatusOK, (string)(b))
}

// swagger:route GET /plugins/{idMedia} GetHandler
//
// Gets information of a plugin
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
// swagger:parameters idPlugin
func (s *Service) GetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eltName := vars["eltName"]

	plugin, err := s.Manager.Get(eltName)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotFound, err.Error())
		return
	}

	b, err := json.Marshal(*plugin)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotFound, err.Error())
		return
	}

	commons.WriteResponse(w, http.StatusOK, (string)(b))
}

func (s *Service) AddHandler(w http.ResponseWriter, r *http.Request) {
	// 0 : Get files content and copy it into a temporary folder
	filename, err := UploadFile(r)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotFound, err.Error())
	}

	// 1 : Check extension
	ext, err := CheckExtension(filename)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotAcceptable, err.Error())
	}
	fmt.Println(ext)

	// 2 : open zip file
	// 3 : parse description file
	// 4 : check there's no plugin already installed with same name or reject
	// 5 : unzip into /plugins folder
	// 6 : save into plugins.json file
	// 7 : delete temporary file

	commons.WriteResponse(w, http.StatusOK, "Plugin correctly added in the catalog")
}

func UploadFile(r *http.Request) (string, error) {
	file, header, err := r.FormFile("uploadfile")

	if err != nil {
		log.Println(err)
		return "", err
	}

	defer file.Close()

	out, err := os.Create(PLUGINS_TEMPORARY_FOLDER + string(os.PathSeparator) + header.Filename)
	if err != nil {
		log.Println("Unable to create the file for writing. Check your write access privilege")
		return "", err
	}

	defer out.Close()

	// write the content from POST to the file
	_, err = io.Copy(out, file)
	if err != nil {
		log.Println(err)
		return "", err
	}

	log.Println("File uploaded successfully : ")

	return header.Filename, nil
}

// Return extension of the file or an error if the extension is not supported by this program
func CheckExtension(filename string) (string, error) {
	acceptedExtensions := []string{".zip", ".gzip", ".tar"}

	ext := path.Ext(filename)

	if accepted, _ := commons.IsInArray(ext, acceptedExtensions); accepted == false {
		v := strings.Join(acceptedExtensions, ",")
		return "", errors.New("File extension (" + ext + ") is not supported. Accepted extensions are: " + v)
	}

	return ext, nil
}
