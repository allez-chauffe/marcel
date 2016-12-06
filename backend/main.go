package main

import (
	"github.com/Zenika/MARCEL/backend/agenda"
	"github.com/Zenika/MARCEL/backend/auth"
	"github.com/Zenika/MARCEL/backend/weather"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

func main() {

	c := cors.New(cors.Options{
		//AllowedOrigins:   []string{"*"},
		AllowedOrigins:   []string{"http://localhost:*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTION", "PUT"},
		AllowCredentials: true,
	})

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/weather/forecast/{nbForecast:[1-9]+}", weather.GetForecastWeatherHandler)
	r.HandleFunc("/api/v1/agenda/incoming/{nbEvents:[1-9]*}", agenda.GetNextEvents)
	r.HandleFunc("/api/v1/GoogleLogin", auth.HandleGoogleLogin)
	r.HandleFunc("/api/v1/GoogleCallback", auth.HandleGoogleCallback)

	handler := c.Handler(r)
	http.ListenAndServe(":8090", handler)
}
