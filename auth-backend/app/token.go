package app

import (
	"errors"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func generateRefreshToken(w http.ResponseWriter) {
	addTokenCookie(w,
		&RefreshClaims{
			StandardClaims: jwt.StandardClaims{
				Subject:  "Valentin",
				IssuedAt: time.Now().Unix(),
			},
		},
		refreshCookie, "/login",
		time.Now().Add(3600*time.Hour),
	)
}

func generateAuthToken(w http.ResponseWriter) {
	expiration := time.Now().Add(30 * time.Second)

	addTokenCookie(w,
		&Claims{
			Role: "admin",
			StandardClaims: jwt.StandardClaims{
				Subject:   "Valentin",
				ExpiresAt: expiration.Unix(),
				IssuedAt:  time.Now().Unix(),
			},
		},
		authCookie, "/",
		expiration,
	)
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
