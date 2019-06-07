package api

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/Zenika/MARCEL/api/auth"
	"github.com/Zenika/MARCEL/api/clients"
	"github.com/Zenika/MARCEL/api/commons"
	"github.com/Zenika/MARCEL/api/db"
	"github.com/Zenika/MARCEL/api/medias"
	"github.com/Zenika/MARCEL/api/plugins"
	"github.com/Zenika/MARCEL/config"
)

type App struct {
	Router http.Handler

	mediaService   *medias.Service
	clientsService *clients.Service
}

func (a *App) Initialize() {
	if err := db.Open(); err != nil {
		log.Fatalln(err)
	}
	a.waitSignal()
	a.initializeData()
	a.initializeRouter()
}

func (a *App) Run() {
	log.Infof("Starting backend server on port %d...", config.Config.Port)

	if !config.Config.Auth.Secure {
		log.Warnln("Secure mode is disabled")
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), a.Router))
}

func (a *App) waitSignal() {
	ch := make(chan os.Signal, 1)

	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-ch

		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}

		os.Exit(0)
	}()
}

func (a *App) initializeRouter() {
	r := mux.NewRouter()
	a.Router = r

	r.Use(auth.Middleware)

	s := r.PathPrefix("/api/v" + commons.MarcelAPIVersion).Subrouter()

	medias := s.PathPrefix("/medias").Subrouter()
	medias.HandleFunc("/", a.mediaService.GetAllHandler).Methods("GET")
	medias.HandleFunc("/", a.mediaService.CreateHandler).Methods("POST")
	medias.HandleFunc("/", a.mediaService.SaveHandler).Methods("PUT")

	media := medias.PathPrefix("/{idMedia:[0-9]*}").Subrouter()
	media.HandleFunc("/", a.mediaService.GetHandler).Methods("GET")
	media.HandleFunc("/", a.mediaService.DeleteHandler).Methods("DELETE")
	media.HandleFunc("/activate", a.mediaService.ActivateHandler).Methods("GET")
	media.HandleFunc("/deactivate", a.mediaService.DeactivateHandler).Methods("GET")
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

	pluginsRouter := s.PathPrefix("/plugins").Subrouter()
	pluginsRouter.HandleFunc("/", plugins.GetAllHandler).Methods("GET")
	pluginsRouter.HandleFunc("/", plugins.AddHandler).Methods("POST")
	pluginsRouter.HandleFunc("/{eltName}", plugins.GetHandler).Methods("GET")
	pluginsRouter.HandleFunc("/{eltName}", plugins.UpdateHandler).Methods("PUT")

	auth := s.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/login", loginHandler).Methods("POST")
	auth.HandleFunc("/logout", logoutHandler).Methods("PUT")
	auth.HandleFunc("/validate", validateHandler).Methods("GET")
	auth.HandleFunc("/validate/admin", validateAdminHandler).Methods("GET")

	users := auth.PathPrefix("/users").Subrouter()
	users.HandleFunc("/", createUserHandler).Methods("POST")
	users.HandleFunc("/", getUsersHandler).Methods("GET")

	user := users.PathPrefix("/{userID}").Subrouter()
	user.HandleFunc("", deleteUserHandler).Methods("DELETE")
	user.HandleFunc("", updateUserHandler).Methods("PUT")
}

func (a *App) initializeData() {
	//load clients list from DB
	a.clientsService = clients.NewService()

	//Load Medias configuration from DB
	a.mediaService = medias.NewService(a.clientsService)
}
