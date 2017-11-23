package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/Zenika/MARCEL/auth-backend/users"
	jwt "github.com/dgrijalva/jwt-go"
)

const refreshCookie = "RefreshAuthentication"

type RefreshClaims struct {
	jwt.StandardClaims
}

func GenerateRefreshToken(w http.ResponseWriter, user *users.User) {
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

func GetRefreshToken(r *http.Request) (*RefreshClaims, error) {
	cookie, err := r.Cookie(refreshCookie)
	if err != nil {
		return nil, errors.New("No Refresh Token")
	}

	claims, err := getVerifiedClaims(cookie.Value, &RefreshClaims{})
	if err != nil {
		return nil, err
	}

	refreshClaims, ok := claims.(*RefreshClaims)
	if !ok {
		return nil, errors.New("Invlaid Refresh Token")
	}

	return refreshClaims, nil
}
