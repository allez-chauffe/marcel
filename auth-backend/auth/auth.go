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

func createTokenCookie(claims jwt.Claims, name string, path string, expiration time.Time) (*http.Cookie, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)

	if err != nil {
		return nil, err
	}

	cookie := &http.Cookie{
		Name:     name,
		Value:    token,
		Expires:  expiration,
		Domain:   config.Domain,
		Secure:   config.SecuredCookies,
		HttpOnly: true,
		Path:     path,
	}

	return cookie, nil
}

func deleteCookie(name, path string) *http.Cookie {
	cookie := &http.Cookie{
		Name:     name,
		Expires:  time.Now(),
		Domain:   config.Domain,
		Path:     path,
		Secure:   config.SecuredCookies,
		HttpOnly: true,
	}

	return cookie
}
