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
	"github.com/Zenika/MARCEL/api/user"
)

type User struct {
	ID               string    `json:"id"`
	DisplayName      string    `json:"displayName"`
	Login            string    `json:"login"`
	Role             string    `json:"role"`
	CreatedAt        time.Time `json:"createdAt"`
	LastDisconection time.Time `json:"lastDisconnection"`
	Password         string    `json:"password"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckPermissions(r, nil, "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	body := getUserFromRequest(w, r)

	if body.Login == "" || body.DisplayName == "" || body.Password == "" {
		commons.WriteResponse(w, http.StatusBadRequest, "Malformed request, missing required fields")
		return
	}

	u, err := user.New(body.DisplayName, body.Login, body.Role, body.Password)
	if err != nil {
		commons.WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = users.Insert(u)
	if err != nil {
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

	body := getUserFromRequest(w, r)
	savedUser, err := users.Get(userID)
	if err != nil {
		commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if savedUser == nil || savedUser.ID != body.ID {
		commons.WriteResponse(w, http.StatusNotFound, "")
		return
	}

	if body.Password != "" {
		changed, err := savedUser.CheckPassword(body.Password)
		if err != nil {
			commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		if changed {
			if err := savedUser.SetPassword(body.Password); err != nil {
				commons.WriteResponse(w, http.StatusInternalServerError, err.Error())
				return
			}

			savedUser.LastDisconection = time.Now() // FIXME why ?!
		}
	}

	savedUser.DisplayName = body.DisplayName
	savedUser.Login = body.Login

	if auth.CheckPermissions(r, nil, "admin") {
		savedUser.Role = body.Role
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

	result := make([]*User, len(users))
	for i, u := range users {
		result[i] = adaptUser(&u)
	}

	commons.WriteJsonResponse(w, result)
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

	commons.WriteJsonResponse(w, adaptUser(u))
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

func getUserFromRequest(w http.ResponseWriter, r *http.Request) *User {
	user := &User{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		commons.WriteResponse(w, http.StatusBadRequest, fmt.Sprintf("Error while parsing JSON (%s)", err.Error()))
		return nil
	}

	return user
}

func adaptUser(user *user.User) *User {
	return &User{
		ID:               user.ID,
		DisplayName:      user.DisplayName,
		Login:            user.Login,
		CreatedAt:        user.CreatedAt,
		LastDisconection: user.LastDisconection,
		Role:             user.Role,
	}
}
