package auth

import (
	"context"
	"net/http"

	"github.com/Zenika/marcel/api/commons"
)

type authContextKeyType string

const authContextKey = authContextKeyType("MARCEL/AUT_BACKEND/AUTH")

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := GetAuthToken(r)

		if err != nil {
			h.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), authContextKey, token)

		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetAuth(r *http.Request) *Claims {
	auth, ok := r.Context().Value(authContextKey).(*Claims)

	if !ok {
		return nil
	}

	return auth
}

func CheckPermissions(r *http.Request, users []string, roles ...string) bool {
	auth := GetAuth(r)

	if auth == nil {
		return false
	}

	if len(roles) == 0 && len(users) == 0 {
		return true
	}

	isAuthorizedRole, _ := commons.IsInArray(auth.Role, roles)
	isAuthorizedUser, _ := commons.IsInArray(auth.Subject, users)

	return isAuthorizedUser || isAuthorizedRole
}
