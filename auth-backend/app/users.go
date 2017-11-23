package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	auth "github.com/Zenika/MARCEL/auth-backend/auth/middleware"
	"github.com/Zenika/MARCEL/auth-backend/users"
	"github.com/Zenika/MARCEL/backend/commons"
	"github.com/gorilla/mux"
)

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	if auth.CheckPermissions(r, nil, "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	params := getUserFromRequest(w, r)

	if params.Login == "" || params.DisplayName == "" || params.Password == "" {
		commons.WriteResponse(w, http.StatusBadRequest, "Malformed request, missing required fields")
		return
	}

	user := users.New(params.DisplayName, params.Login, params.Password)
	users.SaveUsersData()

	commons.WriteJsonResponse(w, user)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]

	if !auth.CheckPermissions(r, []string{userID}, "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	user := getUserFromRequest(w, r)
	savedUser := users.GetByID(userID)

	if savedUser == nil || savedUser.ID != user.ID {
		commons.WriteResponse(w, http.StatusNotFound, "")
		return
	}

	if savedUser.Password != user.Password {
		savedUser.LastDisconection = time.Now().Unix()
	}
	savedUser.DisplayName = user.DisplayName
	savedUser.Login = user.Login
	savedUser.Password = user.Password

	if auth.CheckPermissions(r, nil, "admin") {
		savedUser.Role = user.Role
	}

	users.SaveUsersData()
	commons.WriteJsonResponse(w, savedUser)
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	commons.WriteJsonResponse(w, users.GetAll())
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]

	user := users.GetByID(userID)

	if user == nil {
		commons.WriteResponse(w, http.StatusNotFound, "User not found")
		return
	}

	commons.WriteJsonResponse(w, user)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]

	if !auth.CheckPermissions(r, []string{userID}, "admin") {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	ok := users.Delete(userID)

	if !ok {
		commons.WriteResponse(w, http.StatusNotFound, "User not found")
		return
	}

	users.SaveUsersData()
	commons.WriteResponse(w, http.StatusNoContent, "")
}

func getUserFromRequest(w http.ResponseWriter, r *http.Request) *users.User {
	user := &users.User{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		commons.WriteResponse(w, http.StatusBadRequest, fmt.Sprintf("Error while parsing JSON (%s)", err.Error()))
		return nil
	}

	return user
}
