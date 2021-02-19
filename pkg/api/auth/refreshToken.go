package auth

import (
	"errors"
	"net/http"
	"path"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/allez-chauffe/marcel/pkg/api/commons"
	"github.com/allez-chauffe/marcel/pkg/config"
	"github.com/allez-chauffe/marcel/pkg/db/users"
	"github.com/allez-chauffe/marcel/pkg/module"
)

const RefreshCookieName = "RefreshAuthentication"

type RefreshClaims struct {
	jwt.StandardClaims
}

func GenerateRefreshToken(w http.ResponseWriter, user *users.User) {
	cookie, err := createTokenCookie(
		getRefreshClaims(user),
		RefreshCookieName,
		path.Join(module.URI("API"), "auth", "login"),
		time.Now().Add(config.Default().API().Auth().RefreshExpiration()),
	)

	if err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, "Failed to create auth token")
		return
	}

	http.SetCookie(w, cookie)
}

func GetRefreshToken(r *http.Request) (*RefreshClaims, error) {
	cookie, err := r.Cookie(cookieName(RefreshCookieName, path.Join(module.URI("API"), "auth", "login")))
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

func GenerateRefreshJWT(user *users.User) (string, error) {
	return createToken(getRefreshClaims(user))
}

func DeleteRefreshToken(w http.ResponseWriter) {
	cookie := deleteCookie(RefreshCookieName, path.Join(module.URI("API"), "auth", "login"))
	http.SetCookie(w, cookie)
}

func getRefreshClaims(user *users.User) *RefreshClaims {
	return &RefreshClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:  user.ID,
			IssuedAt: time.Now().Unix(),
		},
	}
}
