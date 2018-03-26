package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/Zenika/MARCEL/auth-backend/users"
	"github.com/Zenika/MARCEL/backend/commons"
	jwt "github.com/dgrijalva/jwt-go"
)

type RefreshClaims struct {
	jwt.StandardClaims
}

func GenerateRefreshToken(w http.ResponseWriter, user *users.User) {
	cookie, err := createTokenCookie(
		&RefreshClaims{
			StandardClaims: jwt.StandardClaims{
				Subject:  user.ID,
				IssuedAt: time.Now().Unix(),
			},
		},
		refreshCookieName(), config.BaseURL+"/login",
		time.Now().Add(time.Duration(config.RefreshExpiration)*time.Second),
	)

	if err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, "Failed to create auth token")
		return
	}

	http.SetCookie(w, cookie)
}

func GetRefreshToken(r *http.Request) (*RefreshClaims, error) {
	cookie, err := r.Cookie(refreshCookieName())
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

func DeleteRefreshToken(w http.ResponseWriter) {
	cookie := deleteCookie(refreshCookieName(), "/auth/login")
	http.SetCookie(w, cookie)
}

func refreshCookieName() string {
	const refreshCookieName = "RefreshAuthentication"
	return config.CookiesNamePrefix + refreshCookieName
}
