package plugins

import (
	"net/http"
	"encoding/json"
	"github.com/Zenika/MARCEL/backend/commons"
	"github.com/gorilla/mux"
	"io"
	"os"
	"log"
	"path"
	"errors"
	"strings"
	"archive/zip"
	"path/filepath"
)

const PLUGINS_CONFIG_PATH string = "data"
const PLUGINS_CONFIG_FILENAME string = "plugins.json"
const PLUGINS_TEMPORARY_FOLDER string = "uploadedfiles"
const PLUGINS_FOLDER string = "plugins"

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
		return
	}

	// 1 : Check extension
	_, err = CheckExtension(filename)
	if err != nil {
		os.Remove(PLUGINS_TEMPORARY_FOLDER + string(os.PathSeparator) + filename)
		commons.WriteResponse(w, http.StatusNotAcceptable, err.Error())
		return
	}
	// 2 : unzip into /plugins folder
	//todo : créer un dossieravec un nom temporaire pour déposer le plugin et le renomer plus tard avec le nom du plugin
	err = UncompressFile(
		PLUGINS_TEMPORARY_FOLDER+string(os.PathSeparator)+filename,
		PLUGINS_FOLDER+string(os.PathSeparator)+commons.Basename(filename)+string(os.PathSeparator))
	if err != nil {
		commons.WriteResponse(w, http.StatusNotAcceptable, err.Error())
		return
	}
	// 3 : parse description file
	// 4 : check there's no plugin already installed with same name or reject
	// 5 : save into plugins.json file
	// 6 : delete temporary file

	commons.WriteResponse(w, http.StatusOK, "Plugin correctly added to the catalog")
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
	acceptedExtensions := []string{".zip"}

	ext := path.Ext(filename)

	if accepted, _ := commons.IsInArray(ext, acceptedExtensions); accepted == false {
		v := strings.Join(acceptedExtensions, ", ")
		return "", errors.New("File extension (" + ext + ") is not supported. Accepted extensions are: " + v)
	}

	return ext, nil
}

func UncompressFile(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}