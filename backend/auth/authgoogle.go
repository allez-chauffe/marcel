package auth

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
	"io/ioutil"
	"os"
	"log"
)

var Google_API_client *http.Client = nil
var defaultNbEvents string = "10"

var scopes = []string {"https://www.googleapis.com/auth/calendar"}


func ReadKey() []byte{
	var key, err = ioutil.ReadFile(os.Getenv("GOOGLE_API_KEY"));
	if err != nil {
		log.Fatal(err);
	}
	return key;
}

func RequestAuthenticatedClient() *http.Client {
	var key = ReadKey();
	var googleOauthConfig, err = google.JWTConfigFromJSON(
		key,
		scopes...)
	if err != nil {
		log.Fatal(err);
	}
	return googleOauthConfig.Client(oauth2.NoContext);
}
