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
	"github.com/Zenika/marcel/module"
)

func Module() module.Module {
	var a = New()

	return module.Module{
		Name: "API",
		Start: func(next module.StartNextFunc) (module.StopFunc, error) {
			if err := db.Open(); err != nil {
				return nil, err //FIXME wrap
			}

			var stop = func() error {
				if err := db.Close(); err != nil {
					return err // FIXME wrap
				}
				return nil
			}

			if err := users.EnsureOneUser(); err != nil {
				return stop, err
			}

			a.initServices() // FIXME manage errors ?!

			return stop, next()
		},
		Http: &module.Http{
			BasePath: config.Default().API().BasePath(),
			Setup:    a.configureSubRouter,
		},
	}
}

type API struct {
	mediaService   *medias.Service
	clientsService *clients.Service
}

func New() *API {
	return new(API)
}

func (a *API) Init() error {
	if err := db.Open(); err != nil {
		return err
	}

	a.waitSignal()

	if err := users.EnsureOneUser(); err != nil {
		return err
	}

	a.initServices()

	return nil
}

func (a *API) Start() {
	r := mux.NewRouter()

	a.ConfigureRouter(r)

	var h http.Handler = r

	if config.Default().API().CORS() {
		h = cors.New(cors.Options{
			AllowOriginFunc:  func(origin string) bool { return true },
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
			AllowedHeaders:   []string{"*"},
			AllowCredentials: true,
		}).Handler(h)
		log.Warn("CORS is enabled")
	}

	log.Infof("API server listening on %d...", config.Default().API().Port())

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Default().API().Port()), h))
}

func (a *API) ConfigureRouter(r *mux.Router) error {
	return a.configureSubRouter(r.PathPrefix(config.Default().API().BasePath()).Subrouter())
}

func (a *API) configureSubRouter(r *mux.Router) error {
	r.Use(auth.Middleware)
	if !config.Default().API().Auth().Secure() {
		log.Warnln("Secure mode is disabled")
	}

	medias := r.PathPrefix("/medias").Subrouter()
	medias.HandleFunc("/", a.mediaService.GetAllHandler).Methods("GET")
	medias.HandleFunc("/", a.mediaService.CreateHandler).Methods("POST")
	medias.HandleFunc("/", a.mediaService.SaveHandler).Methods("PUT")

	media := medias.PathPrefix("/{idMedia:[0-9]*}").Subrouter()
	media.HandleFunc("/", a.mediaService.GetHandler).Methods("GET")
	media.HandleFunc("/", a.mediaService.DeleteHandler).Methods("DELETE")
	media.HandleFunc("/activate", a.mediaService.ActivateHandler).Methods("GET")
	media.HandleFunc("/deactivate", a.mediaService.DeactivateHandler).Methods("GET")
	media.HandleFunc("/plugins/{eltName}/{instanceId}/{filePath:.*}", a.mediaService.GetPluginFilesHandler).Methods("GET")

	clients := r.PathPrefix("/clients").Subrouter()
	clients.HandleFunc("/", a.clientsService.GetAllHandler).Methods("GET")
	clients.HandleFunc("/", a.clientsService.CreateHandler).Methods("POST")
	clients.HandleFunc("/", a.clientsService.UpdateHandler).Methods("PUT")
	clients.HandleFunc("/", a.clientsService.DeleteAllHandler).Methods("DELETE")

	client := clients.PathPrefix("/{clientID}").Subrouter()
	client.HandleFunc("/", a.clientsService.GetHandler).Methods("GET")
	client.HandleFunc("/", a.clientsService.DeleteHandler).Methods("DELETE")
	client.HandleFunc("/ws", a.clientsService.WSConnectionHandler)

	pluginsRouter := r.PathPrefix("/plugins").Subrouter()
	pluginsRouter.HandleFunc("/", plugins.GetAllHandler).Methods("GET")
	pluginsRouter.HandleFunc("/", plugins.AddHandler).Methods("POST")
	pluginsRouter.HandleFunc("/{eltName}", plugins.GetHandler).Methods("GET")
	pluginsRouter.HandleFunc("/{eltName}", plugins.UpdateHandler).Methods("PUT")
	pluginsRouter.HandleFunc("/{eltName}", plugins.DeleteHandler).Methods("DELETE")

	auth := r.PathPrefix("/auth").Subrouter()
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

	return nil
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

func (a *API) initServices() {
	//load clients list from DB
	a.clientsService = clients.NewService()

	//Load Medias configuration from DB
	a.mediaService = medias.NewService(a.clientsService)

	plugins.Initialize()
}
