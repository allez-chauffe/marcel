package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Zenika/MARCEL/auth-backend/users"
	"github.com/Zenika/MARCEL/backend/commons"
)

type Credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	cred := getCredentials(w, r)

	if cred != nil {
		loginWithCredentials(w, cred.Login, cred.Password)
	} else {
		loginWithRefreshToken(w, r)
		return
	}
}

func loginWithCredentials(w http.ResponseWriter, login string, password string) {
	if login == "" || password == "" {
		commons.WriteResponse(w, http.StatusBadRequest, "Missing required fields")
		return
	}

	user := users.GetByLogin(login, password)

	if user == nil {
		w.WriteHeader(403)
		w.Write([]byte("Wrong login or password"))
		return
	}

	generateAuthToken(w, user)
	generateRefreshToken(w, user)
}

func loginWithRefreshToken(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(refreshCookie)
	if err != nil {
		w.WriteHeader(403)
		w.Write([]byte("No refresh token"))
		return
	}

	claims, err := getVerifiedClaims(cookie.Value, &RefreshClaims{})
	if err != nil {
		w.WriteHeader(403)
		w.Write([]byte(err.Error()))
		return
	}

	refreshClaims := claims.(*RefreshClaims)

	user := users.GetByID(refreshClaims.Subject)
	if user == nil {
		commons.WriteResponse(w, http.StatusNotFound, "User not found")
		return
	}

	if user.LastDisconection > refreshClaims.IssuedAt {
		commons.WriteResponse(w, http.StatusForbidden, "Refresh token has been invalidated")
		return
	}

	generateAuthToken(w, user)
}

func getCredentials(w http.ResponseWriter, r *http.Request) *Credentials {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		return nil
	}

	credentials := &Credentials{}
	if err := json.Unmarshal(body, credentials); err != nil {
		commons.WriteResponse(w, http.StatusBadRequest, fmt.Sprintf("Error while parsing JSON (%s)", err.Error()))
		return nil
	}

	return credentials
}
