package api

import (
	"github.com/gorilla/mux"
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
	var clientsService *clients.Service
	var mediasService *medias.Service

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

			//load clients list from DB
			clientsService = clients.NewService()

			//Load Medias configuration from DB
			mediasService = medias.NewService(clientsService)

			plugins.Initialize()

			return stop, next()
		},
		Http: &module.Http{
			BasePath: config.Default().API().BasePath(),
			Setup: func(r *mux.Router) {
				r.Use(auth.Middleware)
				if !config.Default().API().Auth().Secure() {
					log.Warnln("Secure mode is disabled")
				}

				medias := r.PathPrefix("/medias").Subrouter()
				medias.HandleFunc("/", mediasService.GetAllHandler).Methods("GET")
				medias.HandleFunc("/", mediasService.CreateHandler).Methods("POST")
				medias.HandleFunc("/", mediasService.SaveHandler).Methods("PUT")

				media := medias.PathPrefix("/{idMedia:[0-9]*}").Subrouter()
				media.HandleFunc("/", mediasService.GetHandler).Methods("GET")
				media.HandleFunc("/", mediasService.DeleteHandler).Methods("DELETE")
				media.HandleFunc("/activate", mediasService.ActivateHandler).Methods("GET")
				media.HandleFunc("/deactivate", mediasService.DeactivateHandler).Methods("GET")
				media.HandleFunc("/plugins/{eltName}/{instanceId}/{filePath:.*}", mediasService.GetPluginFilesHandler).Methods("GET")

				clients := r.PathPrefix("/clients").Subrouter()
				clients.HandleFunc("/", clientsService.GetAllHandler).Methods("GET")
				clients.HandleFunc("/", clientsService.CreateHandler).Methods("POST")
				clients.HandleFunc("/", clientsService.UpdateHandler).Methods("PUT")
				clients.HandleFunc("/", clientsService.DeleteAllHandler).Methods("DELETE")

				client := clients.PathPrefix("/{clientID}").Subrouter()
				client.HandleFunc("/", clientsService.GetHandler).Methods("GET")
				client.HandleFunc("/", clientsService.DeleteHandler).Methods("DELETE")
				client.HandleFunc("/ws", clientsService.WSConnectionHandler)

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
			},
		},
	}
}
