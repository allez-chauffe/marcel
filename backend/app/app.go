package app

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/Zenika/MARCEL/backend/medias"
	"github.com/rs/cors"
	"log"
	"os"
	"github.com/Zenika/MARCEL/backend/apidoc"
	"github.com/Zenika/MARCEL/backend/plugins"
)

//current version of the API
const MARCEL_API_VERSION = "1"

var logFileName string = os.Getenv("MARCEL_LOG_FILE")
var logFile *os.File

type App struct {
	Router http.Handler

	mediaService  *medias.Service
	pluginService *plugins.Service
}

func (a *App) Initialize() {

	err := a.InitializeLog(logFileName, logFile)
	if err != nil {
		print(err)
	}

	a.initializeData()

	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
	log.Printf("Server is started and listening on port %v", addr)

	defer logFile.Close()

	select {}
}

func (a *App) initializeRoutes() {

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		// AllowedOrigins:   []string{"http://localhost:*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTION", "PUT"},
		AllowCredentials: true,
	})

	r := mux.NewRouter()
	s := r.PathPrefix("/api/v" + MARCEL_API_VERSION).Subrouter()
	s.HandleFunc("/medias", a.mediaService.GetAllHandler).Methods("GET")
	s.HandleFunc("/medias", a.mediaService.PostHandler).Methods("POST")
	s.HandleFunc("/medias/{idMedia:[0-9]*}", a.mediaService.GetHandler).Methods("GET")
	s.HandleFunc("/medias/config", a.mediaService.GetConfigHandler).Methods("GET")
	s.HandleFunc("/medias/create", a.mediaService.CreateHandler).Methods("GET")
	s.HandleFunc("/plugins", a.pluginService.GetAllHandler).Methods("GET")
	s.HandleFunc("/plugins/config", a.pluginService.GetConfigHandler).Methods("GET")
	s.HandleFunc("/plugins/add", a.pluginService.AddHandler).Methods("POST")
	s.HandleFunc("/plugins/{eltName}", a.pluginService.GetHandler).Methods("GET")
	r.HandleFunc("/swagger.json", apidoc.GetConfigHandler).Methods("GET")

	a.Router = c.Handler(r)
}

func (a *App) InitializeLog(filename string, logFile *os.File) error {
	if len(filename) == 0 {
		filename = "marcel.log"
	}
	logFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	log.SetOutput(logFile)

	return nil
}

func (a *App) initializeData() {

	//Load plugins list from DB
	a.pluginService = plugins.NewService()
	a.pluginService.GetManager().Load()

	//Load Medias configuration from DB
	a.mediaService = medias.NewService()
	a.mediaService.GetManager().Load()
}
