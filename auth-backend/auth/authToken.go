package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/Zenika/MARCEL/backend/commons"

	"github.com/Zenika/MARCEL/auth-backend/users"
	jwt "github.com/dgrijalva/jwt-go"
)

const authCookie = "Authentication"

type Claims struct {
	DisplayName string `json:"display"`
	Role        string `json:"role"`
	jwt.StandardClaims
}

func GetAuthToken(r *http.Request) (*Claims, error) {
	cookie, err := r.Cookie(authCookie)

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
	expiration := time.Now().Add(time.Duration(config.AuthExpiration) * time.Second)

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
		authCookie, "/",
		expiration,
	)

	if err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, "Failed to create auth token")
		return
	}

	http.SetCookie(w, cookie)
}
