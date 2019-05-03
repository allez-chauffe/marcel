package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"

	"github.com/Zenika/MARCEL/backend/auth/middleware"
	"github.com/Zenika/MARCEL/backend/config"
)

var (
	secretKey = []byte("ThisIsTheSecret")
	app       http.Handler
)

func init() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTION", "PUT"},
		AllowCredentials: true,
	})

	r := mux.NewRouter()
	base := r.PathPrefix("").Subrouter()

	base.HandleFunc("/login", loginHandler).Methods("POST")
	base.HandleFunc("/logout", logoutHandler).Methods("PUT")
	base.HandleFunc("/validate", validateHandler).Methods("GET")
	base.HandleFunc("/validate/admin", validateAdminHandler).Methods("GET")

	userRoutes(base.PathPrefix("/users").Subrouter())

	app = handlers.LoggingHandler(os.Stdout, middleware.AuthMiddlware(c.Handler(r)))
}

func userRoutes(r *mux.Router) {
	r.HandleFunc("/", createUserHandler).Methods("POST")
	r.HandleFunc("/", getUsersHandler).Methods("GET")
	r.HandleFunc("/{userID}", getUserHandler).Methods("GET")
	r.HandleFunc("/{userID}", deleteUserHandler).Methods("DELETE")
	r.HandleFunc("/{userID}", updateUserHandler).Methods("PUT")
}

func Run() {
	addr := fmt.Sprintf(":%d", config.Config.Auth.Port)

	secureMode := ""
	if config.Config.Auth.Secured {
		secureMode = " with secure mode enabled"
	}

	log.Infof("Starting auth server on %s%s", addr, secureMode)
	log.Fatal(http.ListenAndServe(addr, app))
}
