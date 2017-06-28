package main

import (
	"github.com/Zenika/MARCEL/backend/agenda"
	"github.com/Zenika/MARCEL/backend/weather"
	"github.com/Zenika/MARCEL/backend/twitter"
	"github.com/Zenika/MARCEL/backend/medias"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"log"
	"os"
	"strings"
)

//todo : service to create a new media
//todo : service to log with jwt (or at least a token based system)
//todo : service to delete a media
//todo : service to save an existing media

var logFile string = os.Getenv("MARCEL_LOG_FILE")
//current version of the API
const MARCEL_API_VERSION = "1"

func main() {

	if len(logFile) == 0 {
		logFile = "marcel.log"
	}
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)

	//Load plugins list from DB
	//plugins.LoadPluginsCatalog()

	//Load Medias configuration from DB
	medias.LoadMedias()

	//Set API services
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		// AllowedOrigins:   []string{"http://localhost:*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTION", "PUT"},
		AllowCredentials: true,
	})

	r := mux.NewRouter()
	s := r.PathPrefix("/api/v" + MARCEL_API_VERSION).Subrouter()
	s.HandleFunc("/weather/forecast/{nbForecasts:[0-9]+}", weather.GetForecastWeatherHandler).Methods("GET")
	s.HandleFunc("/agenda/incoming/{nbEvents:[0-9]*}", agenda.GetNextEvents).Methods("GET")
	s.HandleFunc("/twitter/timeline/{nbTweets:[0-9]*}", twitter.GetTimeline).Methods("GET")
	s.HandleFunc("/medias/{idMedia:[0-9]*}", medias.HandleGetMedia).Methods("GET")
	s.HandleFunc("/medias", medias.HandleGetAll).Methods("GET")
	s.HandleFunc("/medias/create", medias.HandleCreate).Methods("GET")

	handler := c.Handler(r)

	ExposeAPI(r)

	http.ListenAndServe(":8090", handler)
	log.Printf("Server is started and listening on port %v", 8090)
}

func ExposeAPI(r *mux.Router) {
	//export API to the log
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		// p will contain regular expression is compatible with regular expression in Perl, Python, and other languages.
		// for instance the regular expression for path '/articles/{id}' will be '^/articles/(?P<v0>[^/]+)$'
		p, err := route.GetPathRegexp()
		if err != nil {
			return err
		}
		m, err := route.GetMethods()
		if err != nil {
			return err
		}
		log.Println(strings.Join(m, ","), t, p)
		return nil
	})
}
