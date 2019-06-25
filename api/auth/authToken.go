package auth

import (
	"errors"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/Zenika/marcel/api/commons"
	"github.com/Zenika/marcel/api/db/users"
	"github.com/Zenika/marcel/config"
)

const AuthCookieName = "Authentication"

type Claims struct {
	DisplayName string `json:"display"`
	Role        string `json:"role"`
	jwt.StandardClaims
}

func GetAuthToken(r *http.Request) (*Claims, error) {
	cookie, err := r.Cookie(cookieName(AuthCookieName, "/"))

	if err != nil {
		return nil, err
	}

	token, err := getVerifiedClaims(cookie.Value, &Claims{})

	claims, ok := token.(*Claims)
	if !ok || err != nil {
		return nil, errors.New("Invalid Auth Token")
	}

	return claims, nil
}

func GenerateAuthToken(w http.ResponseWriter, user *users.User) {
	expiration := time.Now().Add(config.Config.API.Auth.Expiration)

	cookie, err := createTokenCookie(
		&Claims{
			DisplayName: user.DisplayName,
			Role:        user.Role,
			StandardClaims: jwt.StandardClaims{
				Subject:   user.ID,
				ExpiresAt: expiration.Unix(),
				IssuedAt:  time.Now().Unix(),
			},
		},
		AuthCookieName,
		"/",
		expiration,
	)

	if err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, "Failed to create auth token")
		return
	}

	http.SetCookie(w, cookie)
}

func DeleteAuthToken(w http.ResponseWriter) {
	cookie := deleteCookie(AuthCookieName, "/")
	http.SetCookie(w, cookie)
}
