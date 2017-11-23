package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Zenika/MARCEL/auth-backend/auth"
	"github.com/Zenika/MARCEL/auth-backend/auth/middleware"
	"github.com/Zenika/MARCEL/auth-backend/conf"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const (
	authCookie    = "Authentication"
	refreshCookie = "RefreshAuthentication"
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

	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/validate", validateHandler).Methods("GET")

	userRoutes(r.PathPrefix("/users").Subrouter())

	app = middleware.AuthMiddlware(c.Handler(r))
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
