package app

import (
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")

	if login == "" {
		loginWithRefreshToken(w, r)
		return
	}

	if password == "" {
		w.WriteHeader(400)
		w.Write([]byte("Missing password"))
		return
	}

	loginWithCredentials(w, login, password)
}

func loginWithCredentials(w http.ResponseWriter, login string, password string) {
	if login != "admin" || password != "admin" {
		w.WriteHeader(403)
		w.Write([]byte("Wrong login or password"))
	}

	generateAuthToken(w)
	generateRefreshToken(w)
}

func loginWithRefreshToken(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(refreshCookie)
	if err != nil {
		w.WriteHeader(403)
		w.Write([]byte("No refresh token"))
		return
	}

	_, err = getVerifiedClaims(cookie.Value, &RefreshClaims{})
	if err != nil {
		w.WriteHeader(403)
		w.Write([]byte(err.Error()))
		return
	}

	generateAuthToken(w)
}
