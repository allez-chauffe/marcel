package weather

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	owm "github.com/gwennaelbuchet/openweathermap"
)

// URL is a constant that contains where to find the IP locale info
const URL = "http://ip-api.com/json"

// LocationData will hold the result of the query to get the IP
// address of the caller.
type LocationData struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	ISP         string  `json:"isp"`
	ORG         string  `json:"org"`
	AS          string  `json:"as"`
	Message     string  `json:"message"`
	Query       string  `json:"query"`
}

// getLocation will get the location details for where this
// application has been run from.
func getLocation() (*LocationData, error) {
	response, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	r := &LocationData{}
	if err = json.NewDecoder(response.Body).Decode(&r); err != nil {
		return nil, err
	}
	return r, nil
}

func getForecast(l *LocationData, u string, lang string, nbForecast int) *owm.HourlyForecastData {
	w, err := owm.NewHourlyForecast(u, lang) // Create the instance with the given unit
	if err != nil {
		log.Fatal(err)
	}

	w.HourlyForecastByName(l.City, nbForecast)

	return w
}

// GetForecastWeatherHandler take care of request for this current day weather data
func GetForecastWeatherHandler(w http.ResponseWriter, r *http.Request) {

	location, err := getLocation()
	if err != nil {
		log.Fatal(err)
	}

	vars := mux.Vars(r)
	f := vars["nbForecast"]
	nbForecast, _ := strconv.Atoi(f)

	wd := getForecast(location, "C", "FR", nbForecast)

	j, err := json.Marshal(wd)
	if err == nil {
		w.Write([]byte(j))
	}
}
