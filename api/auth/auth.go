package auth

import (
	"errors"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/Zenika/marcel/config"
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
		Name:     cookieName(name, path),
		Value:    token,
		Expires:  expiration,
		Secure:   config.Config().API().Auth().Secure(),
		HttpOnly: true,
		Path:     path,
		SameSite: http.SameSiteStrictMode,
	}

	return cookie, nil
}

func deleteCookie(name, path string) *http.Cookie {
	cookie := &http.Cookie{
		Name:     cookieName(name, path),
		Expires:  time.Now(),
		Path:     path,
		Secure:   config.Config().API().Auth().Secure(),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	return cookie
}

func cookieName(name, path string) string {
	if !config.Config().API().Auth().Secure() {
		return name
	}
	if path == "/" {
		return "__Host-" + name
	}
	return "__Secure-" + name
}
