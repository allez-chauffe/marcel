package users

import (
	"encoding/json"
	"log"
	"os"
	"time"

	uuid "github.com/satori/go.uuid"
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

func GetByLoginAndPassword(login, password string) *User {
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
	f, err := os.OpenFile(userFilePath, os.O_CREATE, 0755)
	defer f.Close()

	if err != nil {
		log.Println("Error while loading users database", err.Error())
		return
	}

	if err := json.NewDecoder(f).Decode(usersData); err != nil {
		log.Printf("ERROR: Malformed JSON in users database file (%s)", err.Error())
		return
	}
}

func SaveUsersData() {
	f, err := os.OpenFile(userFilePath, os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()

	if err != nil {
		log.Printf("ERROR: Error while opening users database file %s (%s)", userFilePath, err.Error())
		return
	}

	if err := json.NewEncoder(f).Encode(usersData); err != nil {
		log.Printf("ERROR: Error while saving users data in %s (%s)", userFilePath, err.Error())
		return
	}
}
