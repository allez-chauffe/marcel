package app

import (
	"errors"
	"net/http"
	"time"

	"github.com/Zenika/MARCEL/auth-backend/users"
	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	DisplayName string `json:"display"`
	Role        string `json:"role"`
	jwt.StandardClaims
}

type RefreshClaims struct {
	jwt.StandardClaims
}

func generateRefreshToken(w http.ResponseWriter, user *users.User) {
	addTokenCookie(w,
		&RefreshClaims{
			StandardClaims: jwt.StandardClaims{
				Subject:  user.ID,
				IssuedAt: time.Now().Unix(),
			},
		},
		refreshCookie, "/login",
		time.Now().Add(time.Duration(config.RefreshExpiration)*time.Second),
	)
}

func generateAuthToken(w http.ResponseWriter, user *users.User) {
	expiration := time.Now().Add(time.Duration(config.AuthExpiration) * time.Second)

	addTokenCookie(w,
		&Claims{
			DisplayName: user.DisplayName,
			Role:        user.Role,
			StandardClaims: jwt.StandardClaims{
				Subject:   user.ID,
				ExpiresAt: expiration.Unix(),
				IssuedAt:  time.Now().Unix(),
			},
		},
		authCookie, "/",
		expiration,
	)
}

func addTokenCookie(w http.ResponseWriter, claims jwt.Claims, name string, path string, expiration time.Time) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secretKey)

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
		func(token *jwt.Token) (interface{}, error) { return secretKey, nil },
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Invlaid token")
	}

	return token.Claims, nil
}
