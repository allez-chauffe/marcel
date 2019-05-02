package app

import (
	"net/http"

	auth "github.com/Zenika/MARCEL/backend/auth/middleware"
	"github.com/Zenika/MARCEL/backend/commons"
)

func validateHandler(w http.ResponseWriter, r *http.Request) {
	if auth := auth.GetAuth(r); auth == nil {
		commons.WriteResponse(w, http.StatusForbidden, "")
	}
}

func validateAdminHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckPermissions(r, nil, "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}
}
