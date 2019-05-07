package auth

import (
	"errors"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/Zenika/MARCEL/backend/config"
)

var (
	key = []byte("ThisIsTheSecret")
)

func getVerifiedClaims(tokenString string, sampleClaims jwt.Claims) (jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		sampleClaims,
		func(token *jwt.Token) (interface{}, error) { return key, nil },
	)

	if err != nil || !token.Valid {
		return nil, errors.New("Invalid token")
	}

	return token.Claims, nil
}

func createTokenCookie(claims jwt.Claims, name string, path string, expiration time.Time) (*http.Cookie, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)

	if err != nil {
		return nil, err
	}

	cookie := &http.Cookie{
		Name:     name,
		Value:    token,
		Expires:  expiration,
		Domain:   config.Config.Auth.Domain,
		Secure:   config.Config.Auth.Secured,
		HttpOnly: true,
		Path:     path,
	}

	return cookie, nil
}

func deleteCookie(name, path string) *http.Cookie {
	cookie := &http.Cookie{
		Name:     name,
		Expires:  time.Now(),
		Domain:   config.Config.Auth.Domain,
		Path:     path,
		Secure:   config.Config.Auth.Secured,
		HttpOnly: true,
	}

	return cookie
}
