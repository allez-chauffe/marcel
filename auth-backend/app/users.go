package app

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	auth "github.com/Zenika/MARCEL/auth-backend/auth/middleware"
	"github.com/Zenika/MARCEL/auth-backend/users"
	"github.com/Zenika/MARCEL/backend/commons"
	"github.com/gorilla/mux"
)

var passwordSecretKey = []byte("This is the password secret key !")

type User struct {
	ID               string `json:"id"`
	DisplayName      string `json:"displayName"`
	Login            string `json:"login"`
	Role             string `json:"role"`
	CreatedAt        int64  `json:"createdAt"`
	LastDisconection int64  `json:"lastDisconnection"`
	Password         string `json:"password"`
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

	hash, salt := generateHash(body.Password)

	user := users.New(body.DisplayName, body.Login, hash, salt)
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

	body := getUserFromRequest(w, r)
	savedUser := users.GetByID(userID)

	if savedUser == nil || savedUser.ID != body.ID {
		commons.WriteResponse(w, http.StatusNotFound, "")
		return
	}

	if body.Password != "" && !checkHash(body.Password, savedUser.PasswordHash, savedUser.PasswordSalt) {
		savedUser.LastDisconection = time.Now().Unix()
		hash, salt := generateHash(body.Password)
		savedUser.PasswordHash = hash
		savedUser.PasswordSalt = salt
	}
	savedUser.DisplayName = body.DisplayName
	savedUser.Login = body.Login

	if auth.CheckPermissions(r, nil, "admin") {
		savedUser.Role = body.Role
	}

	users.SaveUsersData()
	commons.WriteJsonResponse(w, savedUser)
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	result := []*User{}
	for _, user := range users.GetAll() {
		result = append(result, adaptUser(user))
	}
	commons.WriteJsonResponse(w, result)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckPermissions(r, nil) {
		commons.WriteResponse(w, http.StatusForbidden, "")
		return
	}

	vars := mux.Vars(r)
	userID := vars["userID"]

	user := users.GetByID(userID)

	if user == nil {
		commons.WriteResponse(w, http.StatusNotFound, "User not found")
		return
	}

	commons.WriteJsonResponse(w, adaptUser(user))
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
		commons.WriteResponse(w, http.StatusNotFound, "")
		return
	}

	users.SaveUsersData()
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

func generateHash(password string) (string, string) {
	salt := make([]byte, 40)
	rand.Read(salt)
	saltString := base64.StdEncoding.EncodeToString(salt)

	return hash(password, saltString), saltString
}

func checkHash(password, hashString, saltString string) bool {
	return hash(password, saltString) == hashString
}

func hash(password, salt string) string {
	h := hmac.New(sha256.New, passwordSecretKey)
	h.Write([]byte(password + salt))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func adaptUser(user *users.User) *User {
	return &User{
		ID:               user.ID,
		DisplayName:      user.DisplayName,
		Login:            user.Login,
		CreatedAt:        user.CreatedAt,
		LastDisconection: user.LastDisconection,
		Role:             user.Role,
	}
}
