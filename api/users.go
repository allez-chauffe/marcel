package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/Zenika/marcel/api/auth"
	"github.com/Zenika/marcel/api/commons"
	"github.com/Zenika/marcel/api/db/users"
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
		unchanged, err := savedUser.CheckPassword(payload.Password)
		if err != nil {
			commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		if !unchanged {
			if err := savedUser.SetPassword(payload.Password); err != nil {
				commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
				return
			}

			savedUser.LastDisconnection = time.Now()
		}
	}

	savedUser.DisplayName = payload.DisplayName
	savedUser.Login = payload.Login

	if auth.CheckPermissions(r, nil, "admin") {
		savedUser.Role = payload.Role
	}

	if err := users.Update(savedUser); err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	commons.WriteJsonResponse(w, savedUser)
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckPermissions(r, nil, "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	users, err := users.List()
	if err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	commons.WriteJsonResponse(w, users)
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
	user := new(UserPayload)

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		commons.WriteResponse(w, http.StatusBadRequest, fmt.Sprintf("Error while parsing JSON (%s)", err.Error()))
		return nil
	}

	return user
}
