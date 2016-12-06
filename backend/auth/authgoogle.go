package auth

import (
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
	"os"
)

var Google_API_client *http.Client = nil

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8090/api/v1/GoogleCallback",
		ClientID:     os.Getenv("GOOGLE_API_KEY"),    // from https://console.developers.google.com/project/<your-project-id>/apiui/credential
		ClientSecret: os.Getenv("GOOGLE_API_SECRET"), // from https://console.developers.google.com/project/<your-project-id>/apiui/credential
		Scopes:       []string{"https://www.googleapis.com/auth/calendar"},
		Endpoint:     google.Endpoint,
	}

	// Some random string, random for each request
	oauthStateString = "random"
)

func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/api/v1/agenda/incoming/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		http.Redirect(w, r, "/api/v1/agenda/incoming/", http.StatusTemporaryRedirect)
		return
	}

	Google_API_client = oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))
	http.Redirect(w, r, "/api/v1/agenda/incoming/", http.StatusPermanentRedirect)
}
