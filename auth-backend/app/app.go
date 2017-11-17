package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Zenika/MARCEL/auth-backend/conf"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const (
	authCookie    = "Authentication"
	refreshCookie = "RefreshAuthentication"
)

var (
	key    = []byte("ThisIsTheSecret")
	router http.Handler
	config *conf.Config
)

type Claims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}

type RefreshClaims struct {
	jwt.StandardClaims
}

func init() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTION", "PUT"},
		AllowCredentials: true,
	})

	r := mux.NewRouter()

	r.HandleFunc("/login", loginHandler).Methods("GET")
	r.HandleFunc("/validate", validateHandler).Methods("GET")

	router = c.Handler(r)
}

func Run(c *conf.Config) {
	config = c
	addr := fmt.Sprintf(":%d", config.Port)

	secureMode := ""
	if config.SecuredCookies {
		secureMode = "with secure mode enabled"
	}

	log.Printf("Starting auth-backend on %s %s", addr, secureMode)

	log.Fatal(http.ListenAndServe(addr, router))
}
