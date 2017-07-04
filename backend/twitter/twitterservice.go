package twitter

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Zenika/MARCEL/backend/auth"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/gorilla/mux"
)

func GetTimeline(w http.ResponseWriter, r *http.Request) {
	client := auth.RequireTwitterClient()
	vars := mux.Vars(r)
	e := vars["nbTweets"]
	nbTweets, _ := strconv.Atoi(e)
	userTimelineParams := &twitter.UserTimelineParams{ScreenName: "ZenikaLille", Count: nbTweets}
	tweets, _, _ := client.Timelines.UserTimeline(userTimelineParams)
	j, err := json.Marshal(tweets)
	if err == nil {
		w.Write([]byte(j))
	} else {
		log.Fatal(err)
	}
}
