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
)

var logFile string = os.Getenv("MARCEL_LOG_FILE");

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
	log.Println("Application started")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		// AllowedOrigins:   []string{"http://localhost:*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTION", "PUT"},
		AllowCredentials: true,
	})

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/weather/forecast/{nbForecasts:[0-9]+}", weather.GetForecastWeatherHandler)
	r.HandleFunc("/api/v1/agenda/incoming/{nbEvents:[0-9]*}", agenda.GetNextEvents)
	r.HandleFunc("/api/v1/twitter/timeline/{nbTweets:[0-9]*}", twitter.GetTimeline)
	//r.HandleFunc("/api/v1/plugins/", plugins.GetTimeline)
	r.HandleFunc("/api/v1/media/{idMedia:[0-9]*}", media.GetMedia)


	handler := c.Handler(r)

	http.ListenAndServe(":8090", handler)

}
