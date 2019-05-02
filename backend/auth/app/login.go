package app

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Zenika/MARCEL/backend/auth"
	"github.com/Zenika/MARCEL/backend/auth/middleware"
	"github.com/Zenika/MARCEL/backend/commons"
	"github.com/Zenika/MARCEL/backend/users"
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

	user := users.GetByLogin(login)

	if user == nil || !checkHash(password, user.PasswordHash, user.PasswordSalt) {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	auth.GenerateAuthToken(w, user)
	auth.GenerateRefreshToken(w, user)

	commons.WriteJsonResponse(w, adaptUser(user))
}

func loginWithRefreshToken(w http.ResponseWriter, r *http.Request) {

	refreshClaims, err := auth.GetRefreshToken(r)
	if err != nil {
		commons.WriteResponse(w, http.StatusForbidden, err.Error())
		return
	}

	user := users.GetByID(refreshClaims.Subject)
	if user == nil {
		auth.DeleteRefreshToken(w)
		commons.WriteResponse(w, http.StatusNotFound, "User not found")
		return
	}

	if user.LastDisconection > refreshClaims.IssuedAt {
		auth.DeleteRefreshToken(w)
		commons.WriteResponse(w, http.StatusForbidden, "Refresh token has been invalidated")
		return
	}

	auth.GenerateAuthToken(w, user)

	commons.WriteJsonResponse(w, adaptUser(user))
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	if middleware.CheckPermissions(r, nil) {
		// If the user is connected, update it in database
		userID := middleware.GetAuth(r).Subject

		if user := users.GetByID(userID); user != nil {
			user.LastDisconection = time.Now().Unix()
			users.SaveUsersData()
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
