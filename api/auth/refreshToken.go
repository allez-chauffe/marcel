package auth

import (
	"errors"
	"net/http"
	"path"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/Zenika/marcel/api/commons"
	"github.com/Zenika/marcel/api/db/users"
	"github.com/Zenika/marcel/config"
)

const RefreshCookieName = "RefreshAuthentication"

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
		RefreshCookieName,
		path.Join(config.Config.API.BasePath, "auth", "login"),
		time.Now().Add(config.Config.API.Auth.RefreshExpiration),
	)

	if err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, "Failed to create auth token")
		return
	}

	http.SetCookie(w, cookie)
}

func GetRefreshToken(r *http.Request) (*RefreshClaims, error) {
	cookie, err := r.Cookie(cookieName(RefreshCookieName, path.Join(config.Config.API.BasePath, "auth", "login")))
	if err == http.ErrNoCookie {
		return nil, nil
	}
	if err != nil { // Should not happen
		return nil, err
	}

	claims, err := getVerifiedClaims(cookie.Value, &RefreshClaims{})
	if err != nil {
		return nil, err
	}

	refreshClaims, ok := claims.(*RefreshClaims)
	if !ok {
		return nil, errors.New("Invalid Refresh Token")
	}

	return refreshClaims, nil
}

func DeleteRefreshToken(w http.ResponseWriter) {
	cookie := deleteCookie(RefreshCookieName, path.Join(config.Config.API.BasePath, "auth", "login"))
	http.SetCookie(w, cookie)
}
