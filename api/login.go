package api

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/Zenika/MARCEL/api/auth"
	"github.com/Zenika/MARCEL/api/commons"
	"github.com/Zenika/MARCEL/api/db/users"
)

type Credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	cred := getCredentials(w, r)

	if cred != nil {
		loginWithCredentials(w, cred.Login, cred.Password)
		return
	}

	loginWithRefreshToken(w, r)
}

func loginWithCredentials(w http.ResponseWriter, login string, password string) {
	if login == "" || password == "" {
		commons.WriteResponse(w, http.StatusBadRequest, "Missing required fields")
		return
	}

	user, err := users.GetByLogin(login)
	if err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if user == nil {
		commons.WriteResponse(w, http.StatusUnauthorized, "")
		return
	}

	ok, err := user.CheckPassword(password)
	if err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !ok {
		commons.WriteResponse(w, http.StatusUnauthorized, "")
		return
	}

	auth.GenerateAuthToken(w, user)
	auth.GenerateRefreshToken(w, user)

	commons.WriteJsonResponse(w, user)
}

func loginWithRefreshToken(w http.ResponseWriter, r *http.Request) {
	refreshClaims, err := auth.GetRefreshToken(r)
	if err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	if refreshClaims == nil {
		commons.WriteResponse(w, http.StatusUnauthorized, "Refresh token not found")
		return
	}

	user, err := users.Get(refreshClaims.Subject)
	if err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	if user == nil {
		auth.DeleteRefreshToken(w)
		commons.WriteResponse(w, http.StatusUnauthorized, "User not found")
		return
	}

	if user.LastDisconnection.Unix() > refreshClaims.IssuedAt {
		auth.DeleteRefreshToken(w)
		commons.WriteResponse(w, http.StatusUnauthorized, "Refresh token has been invalidated")
		return
	}

	auth.GenerateAuthToken(w, user)

	commons.WriteJsonResponse(w, user)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	if auth.CheckPermissions(r, nil) {
		// If the user is connected, update it in database
		userID := auth.GetAuth(r).Subject

		if err := users.Disconnect(userID); err != nil {
			// Do not return here we want to delete the tokens
			log.Errorf("Error while disconnecting user %s: %s", userID, err)
		}
	}

	auth.DeleteAuthToken(w)
	auth.DeleteRefreshToken(w)
}

func getCredentials(w http.ResponseWriter, r *http.Request) *Credentials {
	credentials := &Credentials{}

	if err := json.NewDecoder(r.Body).Decode(credentials); err != nil {
		return nil
	}

	return credentials
}
