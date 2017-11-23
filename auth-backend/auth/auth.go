package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/Zenika/MARCEL/auth-backend/conf"
	jwt "github.com/dgrijalva/jwt-go"
)

var (
	key    = []byte("ThisIsTheSecret")
	config *conf.Config
)

func SetConfig(c *conf.Config) {
	config = c
}

func getVerifiedClaims(tokenString string, sampleClaims jwt.Claims) (jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		sampleClaims,
		func(token *jwt.Token) (interface{}, error) { return key, nil },
	)

	if err != nil || !token.Valid {
		return nil, errors.New("Invlaid token")
	}

	return token.Claims, nil
}

func addTokenCookie(w http.ResponseWriter, claims jwt.Claims, name string, path string, expiration time.Time) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)

	if err != nil {
		w.WriteHeader(403)
		w.Write([]byte(err.Error()))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    token,
		Expires:  expiration,
		Secure:   config.SecuredCookies,
		HttpOnly: true,
		Path:     path,
	})
}
