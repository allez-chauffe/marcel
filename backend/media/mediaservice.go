package media

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"io/ioutil"
)

func GetMedia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	f := vars["idMedia"]
	idMedia, _ := strconv.Atoi(f)

	if idMedia > 0 {

		config, err := ioutil.ReadFile("data/media.config.json")
		check(err)

		if err == nil {
			w.Write([]byte(config))
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
