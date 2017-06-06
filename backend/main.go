package main

import (
	"github.com/Zenika/MARCEL/backend/agenda"
	"github.com/Zenika/MARCEL/backend/weather"
	"github.com/Zenika/MARCEL/backend/twitter"
	"github.com/Zenika/MARCEL/backend/media"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"log"
	"os"
	"strings"
)

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
	media.LoadMedias()

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
	s.HandleFunc("/medias/{idMedia:[0-9]*}", media.HandleGetMedia).Methods("GET")
	s.HandleFunc("/medias", media.HandleGetMedias).Methods("GET")

	handler := c.Handler(r)

	http.ListenAndServe(":8090", handler)
	log.Printf("Server is started and listening on port %v", 8090)
}


