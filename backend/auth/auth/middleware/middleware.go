package middleware

import (
	"context"
	"net/http"

	"github.com/Zenika/MARCEL/backend/auth/auth"
	"github.com/Zenika/MARCEL/backend/commons"
)

type authContextKeyType string

const authContextKey = authContextKeyType("MARCEL/AUT_BACKEND/AUTH")

func AuthMiddlware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := auth.GetAuthToken(r)

		if err != nil {
			h.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), authContextKey, token)

		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetAuth(r *http.Request) *auth.Claims {
	auth, ok := r.Context().Value(authContextKey).(*auth.Claims)

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
