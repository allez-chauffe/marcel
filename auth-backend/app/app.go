package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"

	"github.com/Zenika/MARCEL/auth-backend/auth"
	"github.com/Zenika/MARCEL/auth-backend/auth/middleware"
	"github.com/Zenika/MARCEL/auth-backend/conf"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var (
	secretKey = []byte("ThisIsTheSecret")
	app       http.Handler
	config    *conf.Config
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

func Run(c *conf.Config) {
	config = c
	auth.SetConfig(c)
	addr := fmt.Sprintf(":%d", config.Port)

	secureMode := ""
	if config.SecuredCookies {
		secureMode = "with secure mode enabled"
	}

	log.Printf("Starting auth-backend on %s %s", addr, secureMode)

	log.Fatal(http.ListenAndServe(addr, app))
}
