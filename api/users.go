package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/Zenika/MARCEL/api/auth"
	"github.com/Zenika/MARCEL/api/commons"
	"github.com/Zenika/MARCEL/api/db/users"
)

type UserPayload struct {
	*users.User
	Password string `json:"password"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckPermissions(r, nil, "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	payload := getUserPayload(w, r)

	if payload.Login == "" || payload.DisplayName == "" || payload.Password == "" {
		commons.WriteResponse(w, http.StatusBadRequest, "Malformed request, missing required fields")
		return
	}

	u := payload.User

	if err := u.SetPassword(payload.Password); err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := users.Insert(u); err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	commons.WriteJsonResponse(w, u)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]

	if !auth.CheckPermissions(r, []string{userID}, "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	payload := getUserPayload(w, r)

	savedUser, err := users.Get(userID)
	if err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if savedUser == nil || savedUser.ID != payload.ID {
		commons.WriteResponse(w, http.StatusNotFound, "")
		return
	}

	if payload.Password != "" {
		changed, err := savedUser.CheckPassword(payload.Password)
		if err != nil {
			commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		if changed {
			if err := savedUser.SetPassword(payload.Password); err != nil {
				commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
				return
			}

			savedUser.LastDisconnection = time.Now() // FIXME why ?!
		}
	}

	savedUser.DisplayName = payload.DisplayName
	savedUser.Login = payload.Login

	if auth.CheckPermissions(r, nil, "admin") {
		savedUser.Role = payload.Role
	}

	commons.WriteJsonResponse(w, savedUser)
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	// FIXME check permissions

	users, err := users.List()
	if err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	commons.WriteJsonResponse(w, users)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckPermissions(r, nil) { // FIXME no roles?
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	vars := mux.Vars(r)
	userID := vars["userID"]

	u, err := users.Get(userID)
	if err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if u == nil {
		commons.WriteResponse(w, http.StatusNotFound, "User not found")
		return
	}

	commons.WriteJsonResponse(w, u)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]

	if !auth.CheckPermissions(r, []string{userID}, "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	err := users.Delete(userID)
	if err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	commons.WriteResponse(w, http.StatusNoContent, "")
}

func getUserPayload(w http.ResponseWriter, r *http.Request) *UserPayload {
	user := &UserPayload{
		User: &users.User{},
	}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		commons.WriteResponse(w, http.StatusBadRequest, fmt.Sprintf("Error while parsing JSON (%s)", err.Error()))
		return nil
	}

	return user
}
