package plugins

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"github.com/Zenika/MARCEL/backend/auth/auth/middleware"
	"github.com/Zenika/MARCEL/backend/commons"
	"github.com/Zenika/MARCEL/backend/config"
)

var (
	pluginsTempDir     string
	initPluginsTempDir sync.Once
)

type Service struct {
	Manager *Manager
}

func NewService() *Service {
	var p = new(Service)

	p.Manager = NewManager(config.Global.DataPath, config.Global.PluginsFile)

	return p
}

func (s *Service) GetManager() *Manager {
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
	if !middleware.CheckPermissions(r, nil) {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	commons.WriteJsonResponse(w, s.Manager.GetConfiguration())
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
	if !middleware.CheckPermissions(r, nil) {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	commons.WriteJsonResponse(w, m.Manager.GetAll())
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
	if !middleware.CheckPermissions(r, nil) {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	vars := mux.Vars(r)
	eltName := vars["eltName"]

	plugin, err := s.Manager.Get(eltName)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotFound, err.Error())
		return
	}

	commons.WriteJsonResponse(w, plugin)
}

func (s *Service) AddHandler(w http.ResponseWriter, r *http.Request) {
	if !middleware.CheckPermissions(r, nil, "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	// -1 : Create temp dir once
	initPluginsTempDir.Do(func() {
		var err error
		if pluginsTempDir, err = ioutil.TempDir(config.Plugins.Path, "upload"); err != nil {
			log.Panic(err)
		}
	})

	// 0 : Get files content and copy it into a temporary folder
	foldername, filename, err := UploadFile(r)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotFound, err.Error())
		return
	}

	// 1 : Check extension
	_, err = CheckExtension(filename)
	if err != nil {
		os.Remove(filepath.Join(pluginsTempDir, foldername))
		commons.WriteResponse(w, http.StatusNotAcceptable, err.Error())
		return
	}

	// 2 : unzip into /plugins folder
	pluginFolder := filepath.Join(config.Plugins.Path, commons.FileBasename(foldername))

	err = UncompressFile(filepath.Join(pluginsTempDir, foldername), pluginFolder)
	if err != nil {
		commons.WriteResponse(w, http.StatusNotAcceptable, err.Error())
		return
	}

	// 3 : check structure of the plugin
	if exists := commons.FileOrFolderExists(filepath.Join(pluginFolder, "description.json")); exists == false {
		commons.WriteResponse(w, http.StatusNotAcceptable, "'description.json' file not found at the root of the plugin folder")
		return
	}

	if exists := commons.FileOrFolderExists(filepath.Join(pluginFolder, "front")); exists == false {
		commons.WriteResponse(w, http.StatusNotAcceptable, "'front' folder not found at the root of the plugin folder")
		return
	}

	f, err := os.Open(filepath.Join(pluginFolder, "description.json"))
	if err != nil {
		commons.WriteResponse(w, http.StatusNotAcceptable, "Impossible to read 'description.json' file")
		return
	}

	// 4 : Parse description file and add
	plugin := NewPlugin()
	if err = json.NewDecoder(f).Decode(plugin); err != nil {
		commons.WriteResponse(w, http.StatusNotAcceptable, "Impossible to parse 'description.json' file : "+err.Error())
		return
	}

	// todo : if plugin already exists and at least 1 instance of the backend is running, so stop them before replacing the files and relaunch them again after

	// 5 : rename plugin folder with it's EltName (should be unique)
	os.Rename(pluginFolder, filepath.Join(config.Plugins.Path, plugin.EltName))

	// 6 : check there's no plugin already installed with same name or remove&replace
	s.Manager.Add(plugin)

	// 7 : delete temporary file
	os.Remove(filepath.Join(pluginsTempDir, foldername))

	commons.WriteResponse(w, http.StatusOK, "Plugin correctly added to the catalog")
}

func UploadFile(r *http.Request) (string, string, error) {
	file, header, err := r.FormFile("uploadfile")

	if err != nil {
		log.Println(err)
		return "", "", err
	}

	defer file.Close()

	foldername := commons.GetUID()
	out, err := os.Create(filepath.Join(pluginsTempDir, foldername))
	if err != nil {
		log.Println("Unable to create the file for writing. Check your write access privilege")
		return "", "", err
	}

	defer out.Close()

	// write the content from POST to the file
	_, err = io.Copy(out, file)
	if err != nil {
		log.Println(err)
		return "", "", err
	}

	log.Println("File uploaded successfully : ")

	return foldername, header.Filename, nil
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
