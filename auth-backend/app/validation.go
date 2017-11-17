package app

import (
	"log"
	"net/http"
	"time"
)

func validateHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(authCookie)

	if err != nil {
		w.WriteHeader(403)
		w.Write([]byte("Missing Authentication cookie"))
		return
	}

	token, err := getVerifiedClaims(cookie.Value, &Claims{})

	claims, ok := token.(*Claims)
	if !ok || err != nil {
		w.WriteHeader(403)
		w.Write([]byte("Invalid token"))
		return
	}

	log.Printf("User : %s (%s, %d, %d)", claims.Subject, claims.Role, time.Now().Unix()-claims.ExpiresAt, claims.IssuedAt)
}
