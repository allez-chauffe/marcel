package api

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"

	"github.com/Zenika/marcel/api/auth"
	"github.com/Zenika/marcel/api/clients"
	"github.com/Zenika/marcel/api/db"
	"github.com/Zenika/marcel/api/db/users"
	"github.com/Zenika/marcel/api/medias"
	"github.com/Zenika/marcel/api/plugins"
	"github.com/Zenika/marcel/config"
)

type API struct {
	mediaService   *medias.Service
	clientsService *clients.Service
}

func New() *API {
	return new(API)
}

func (a *API) Initialize() {
	if err := db.Open(); err != nil {
		log.Fatalln(err)
	}

	a.waitSignal()

	if err := users.EnsureOneUser(); err != nil {
		log.Fatal(err)
	}

	a.initializeServices()
}

func (a *API) Start() {
	r := mux.NewRouter()

	a.ConfigureRouter(r)

	var h http.Handler = r

	if config.Config().API().CORS() {
		h = cors.New(cors.Options{
			AllowOriginFunc:  func(origin string) bool { return true },
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
			AllowedHeaders:   []string{"*"},
			AllowCredentials: true,
		}).Handler(h)
		log.Warn("CORS is enabled")
	}

	log.Infof("Starting API server on port %d...", config.Config().API().Port())

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Config().API().Port()), h))
}

func (a *API) waitSignal() {
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

func (a *API) ConfigureRouter(r *mux.Router) {
	b := r.PathPrefix(config.Config().API().BasePath()).Subrouter()

	b.Use(auth.Middleware)
	if !config.Config().API().Auth().Secure() {
		log.Warnln("Secure mode is disabled")
	}

	medias := b.PathPrefix("/medias").Subrouter()
	medias.HandleFunc("/", a.mediaService.GetAllHandler).Methods("GET")
	medias.HandleFunc("/", a.mediaService.CreateHandler).Methods("POST")
	medias.HandleFunc("/", a.mediaService.SaveHandler).Methods("PUT")

	media := medias.PathPrefix("/{idMedia:[0-9]*}").Subrouter()
	media.HandleFunc("/", a.mediaService.GetHandler).Methods("GET")
	media.HandleFunc("/", a.mediaService.DeleteHandler).Methods("DELETE")
	media.HandleFunc("/activate", a.mediaService.ActivateHandler).Methods("GET")
	media.HandleFunc("/deactivate", a.mediaService.DeactivateHandler).Methods("GET")
	media.HandleFunc("/plugins/{eltName}/{instanceId}/{filePath:.*}", a.mediaService.GetPluginFilesHandler).Methods("GET")

	clients := b.PathPrefix("/clients").Subrouter()
	clients.HandleFunc("/", a.clientsService.GetAllHandler).Methods("GET")
	clients.HandleFunc("/", a.clientsService.CreateHandler).Methods("POST")
	clients.HandleFunc("/", a.clientsService.UpdateHandler).Methods("PUT")
	clients.HandleFunc("/", a.clientsService.DeleteAllHandler).Methods("DELETE")

	client := clients.PathPrefix("/{clientID}").Subrouter()
	client.HandleFunc("/", a.clientsService.GetHandler).Methods("GET")
	client.HandleFunc("/", a.clientsService.DeleteHandler).Methods("DELETE")
	client.HandleFunc("/ws", a.clientsService.WSConnectionHandler)

	pluginsRouter := b.PathPrefix("/plugins").Subrouter()
	pluginsRouter.HandleFunc("/", plugins.GetAllHandler).Methods("GET")
	pluginsRouter.HandleFunc("/", plugins.AddHandler).Methods("POST")
	pluginsRouter.HandleFunc("/{eltName}", plugins.GetHandler).Methods("GET")
	pluginsRouter.HandleFunc("/{eltName}", plugins.UpdateHandler).Methods("PUT")
	pluginsRouter.HandleFunc("/{eltName}", plugins.DeleteHandler).Methods("DELETE")

	auth := b.PathPrefix("/auth").Subrouter()
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

func (a *API) initializeServices() {
	//load clients list from DB
	a.clientsService = clients.NewService()

	//Load Medias configuration from DB
	a.mediaService = medias.NewService(a.clientsService)

	plugins.Initialize()
}
