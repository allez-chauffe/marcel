package users

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/Zenika/MARCEL/backend/commons"
	"github.com/satori/go.uuid"
)

type User struct {
	ID               string `json:"id"`
	DisplayName      string `json:"displayName"`
	Login            string `json:"login"`
	Password         string `json:"password"`
	Role             string `json:"role"`
	CreatedAt        int64  `json:"createdAt"`
	LastDisconection int64  `json:"lastDisconnection"`
}

type UsersData struct {
	Users []*User `json:"users"`
}

const userFilePath = "data/users.json"

var usersData = &UsersData{[]*User{}}

func New(displayName, login, password string) *User {
	user := &User{
		ID:          uuid.NewV4().String(),
		DisplayName: displayName,
		Login:       login,
		Password:    password,
		Role:        "user",
		CreatedAt:   time.Now().Unix(),
	}

	usersData.Users = append(usersData.Users, user)
	return user
}

func GetAll() []*User {
	return usersData.Users
}

func GetByLogin(login, password string) *User {
	for _, user := range usersData.Users {
		if user.Login == login && user.Password == password {
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
	exists, i := commons.FindIndexInArray(
		func(userI interface{}) bool {
			user, ok := userI.(*User)
			return ok && user.ID == id
		},
		usersData.Users,
	)

	usersData.Users = append(usersData.Users[:i], usersData.Users[i+1:]...)

	return exists
}

func LoadUsersData() {
	if _, err := os.Stat(userFilePath); os.IsNotExist(err) {
		log.Println("WARNING: No users database file detected")
		return
	}

	data, err := ioutil.ReadFile(userFilePath)

	if err != nil {
		log.Printf("ERROR: Error while reading users database file (%s)", err.Error())
		return
	}

	if err := json.Unmarshal(data, usersData); err != nil {
		log.Printf("ERROR: Malformed JSON in users database file (%s)", err.Error())
		return
	}
}

func SaveUsersData() {
	data, err := json.Marshal(usersData)

	if err != nil {
		log.Printf("ERROR: Error while marshalling users data (%s)", err.Error())
		return
	}

	if err := ioutil.WriteFile(userFilePath, data, 0644); err != nil {
		log.Printf("ERROR: Error while saving users data in %s (%s)", userFilePath, err.Error())
		return
	}
}
