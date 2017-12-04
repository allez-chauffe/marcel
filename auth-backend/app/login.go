package app

import (
	"encoding/json"
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
		return
	}

	loginWithRefreshToken(w, r)
}

func loginWithCredentials(w http.ResponseWriter, login string, password string) {
	if login == "" || password == "" {
		commons.WriteResponse(w, http.StatusBadRequest, "Missing required fields")
		return
	}

	user := users.GetByLoginAndPassword(login, password)

	if user == nil {
		commons.WriteResponse(w, http.StatusForbidden, "Wrong login or password")
		return
	}

	generateAuthToken(w, user)
	generateRefreshToken(w, user)
}

func loginWithRefreshToken(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(refreshCookie)
	if err != nil {
		commons.WriteResponse(w, http.StatusForbidden, "No refresh token")
		return
	}

	claims, err := getVerifiedClaims(cookie.Value, &RefreshClaims{})
	if err != nil {
		commons.WriteResponse(w, http.StatusForbidden, err.Error())
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
	credentials := &Credentials{}

	if err := json.NewDecoder(r.Body).Decode(credentials); err != nil {
		return nil
	}

	return credentials
}
