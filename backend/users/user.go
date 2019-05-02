package users

import (
	"encoding/json"
	"os"
	"time"

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

type User struct {
	ID               string `json:"id"`
	DisplayName      string `json:"displayName"`
	Login            string `json:"login"`
	Role             string `json:"role"`
	CreatedAt        int64  `json:"createdAt"`
	LastDisconection int64  `json:"lastDisconnection"`
	PasswordHash     string `json:"passwordHash"`
	PasswordSalt     string `json:"passwordSalt"`
}

type UsersData struct {
	Users []*User `json:"users"`
}

var (
	UsersFilePath string
	usersData     = &UsersData{[]*User{}}
)

func New(displayName, login, pRole, hash, salt string) *User {
	role := "user"
	if pRole != "" {
		role = pRole
	}
	user := &User{
		ID:           uuid.NewV4().String(),
		DisplayName:  displayName,
		Login:        login,
		PasswordHash: hash,
		PasswordSalt: salt,
		Role:         role,
		CreatedAt:    time.Now().Unix(),
	}

	usersData.Users = append(usersData.Users, user)
	return user
}

func GetAll() []*User {
	return usersData.Users
}

func GetByLogin(login string) *User {
	for _, user := range usersData.Users {
		if user.Login == login {
			return user
		}
	}
	return nil
}

func GetByID(id string) *User {
	for _, user := range usersData.Users {
		if user.ID == id {
			return user
		}
	}
	return nil
}

func Delete(id string) bool {
	index := -1

	for i, user := range usersData.Users {
		if user.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		return false
	}

	copy(usersData.Users[index:], usersData.Users[index+1:])
	usersData.Users = usersData.Users[:len(usersData.Users)-1]

	return true
}

func LoadUsersData() {
	f, err := os.OpenFile(UsersFilePath, os.O_CREATE, 0755)
	defer f.Close()

	if err != nil {
		log.Errorln("Error while loading users database", err.Error())
		return
	}

	if err := json.NewDecoder(f).Decode(usersData); err != nil {
		log.Errorf("ERROR: Malformed JSON in users database file (%s)", err.Error())
		return
	}
}

func SaveUsersData() {
	f, err := os.OpenFile(UsersFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer f.Close()

	if err != nil {
		log.Errorf("ERROR: Error while opening users database file %s (%s)", UsersFilePath, err.Error())
		return
	}

	if err := json.NewEncoder(f).Encode(usersData); err != nil {
		log.Errorf("ERROR: Error while saving users data in %s (%s)", UsersFilePath, err.Error())
		return
	}
}
