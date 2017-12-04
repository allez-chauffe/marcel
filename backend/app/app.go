package app

import (
	"log"
	"net/http"
	"os"

	"github.com/Zenika/MARCEL/auth-backend/auth/middleware"
	"github.com/Zenika/MARCEL/backend/apidoc"
	"github.com/Zenika/MARCEL/backend/clients"
	"github.com/Zenika/MARCEL/backend/medias"
	"github.com/Zenika/MARCEL/backend/plugins"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//current version of the API
const MARCEL_API_VERSION = "1"

var logFileName string = os.Getenv("MARCEL_LOG_FILE")
var logFile *os.File

type App struct {
	Router http.Handler

	mediaService   *medias.Service
	pluginService  *plugins.Service
	clientsService *clients.Service
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
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTION", "PUT"},
		AllowCredentials: true,
	})

	r := mux.NewRouter()

	s := r.PathPrefix("/api/v" + MARCEL_API_VERSION).Subrouter()
	r.HandleFunc("/swagger.json", apidoc.GetConfigHandler).Methods("GET")

	medias := s.PathPrefix("/medias").Subrouter()
	medias.HandleFunc("/", a.mediaService.GetAllHandler).Methods("GET")
	medias.HandleFunc("/", a.mediaService.CreateHandler).Methods("POST")
	medias.HandleFunc("/", a.mediaService.SaveHandler).Methods("PUT")
	medias.HandleFunc("/", a.mediaService.DeleteAllHandler).Methods("DELETE")
	medias.HandleFunc("/config", a.mediaService.GetConfigHandler).Methods("GET")

	media := medias.PathPrefix("/{idMedia:[0-9]*}").Subrouter()
	media.HandleFunc("/", a.mediaService.GetHandler).Methods("GET")
	media.HandleFunc("/", a.mediaService.DeleteHandler).Methods("DELETE")
	media.HandleFunc("/activate", a.mediaService.ActivateHandler).Methods("GET")
	media.HandleFunc("/deactivate", a.mediaService.DeactivateHandler).Methods("GET")
	media.HandleFunc("/restart", a.mediaService.RestartHandler).Methods("GET")
	media.HandleFunc("/plugins/{eltName}/{instanceId}/{filePath:.*}", a.mediaService.GetPluginFilesHandler).Methods("GET")

	clients := s.PathPrefix("/clients").Subrouter()
	clients.HandleFunc("/", a.clientsService.GetAllHandler).Methods("GET")
	clients.HandleFunc("/", a.clientsService.CreateHandler).Methods("POST")
	clients.HandleFunc("/", a.clientsService.UpdateHandler).Methods("PUT")
	clients.HandleFunc("/", a.clientsService.DeleteAllHandler).Methods("DELETE")

	client := clients.PathPrefix("/{clientID}").Subrouter()
	client.HandleFunc("/", a.clientsService.GetHandler).Methods("GET")
	client.HandleFunc("/", a.clientsService.DeleteHandler).Methods("DELETE")
	client.HandleFunc("/ws", a.clientsService.WSConnectionHandler)

	plugins := s.PathPrefix("/plugins").Subrouter()
	plugins.HandleFunc("/", a.pluginService.GetAllHandler).Methods("GET")
	plugins.HandleFunc("/config", a.pluginService.GetConfigHandler).Methods("GET")
	plugins.HandleFunc("/add", a.pluginService.AddHandler).Methods("POST")
	plugins.HandleFunc("/{eltName}", a.pluginService.GetHandler).Methods("GET")

	a.Router = middleware.AuthMiddlware(c.Handler(r))
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

	//load clients list from DB
	a.clientsService = clients.NewService()
	a.clientsService.GetManager().LoadFromDB()

	//Load plugins list from DB
	a.pluginService = plugins.NewService()
	a.pluginService.GetManager().LoadFromDB()

	//Load Medias configuration from DB
	a.mediaService = medias.NewService(a.pluginService.GetManager(), a.clientsService)
	a.mediaService.GetManager().LoadFromDB()
}
