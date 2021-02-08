package users

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"time"

	uuid "github.com/satori/go.uuid"
)

// FIXME something ?
var secret = []byte("This is the password secret key !")

type User struct {
	ID                string    `json:"id" boltholdKey:"ID" structs:"id"`
	DisplayName       string    `json:"displayName"`
	Login             string    `json:"login" boltholdIndex:"Login"`
	Role              string    `json:"role"`
	CreatedAt         time.Time `json:"createdAt"`
	LastDisconnection time.Time `json:"lastDisconnection"`
	PasswordHash      string    `json:"-"`
	PasswordSalt      string    `json:"-"`
}

func New() *User {
	return &User{
		ID:        uuid.NewV4().String(),
		Role:      "user",
		CreatedAt: time.Now(),
	}
}

func (u *User) CheckPassword(password string) (bool, error) {
	h, err := hash(password, u.PasswordSalt)
	if err != nil {
		return false, err
	}
	return h == u.PasswordHash, nil
}

func (u *User) SetPassword(password string) error {
	hash, salt, err := generateHash(password)
	if err != nil {
		return err
	}

	u.PasswordHash, u.PasswordSalt = hash, salt

	return nil
}

func generateHash(password string) (string, string, error) {
	salt := make([]byte, 40)

	if _, err := rand.Read(salt); err != nil {
		return "", "", err
	}

	saltString := base64.StdEncoding.EncodeToString(salt)

	h, err := hash(password, saltString)
	if err != nil {
		return "", "", err
	}

	return h, saltString, nil
}

func hash(password, salt string) (string, error) {
	h := hmac.New(sha256.New, secret)

	if _, err := h.Write([]byte(password + salt)); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}

func (u *User) GetID() interface{} {
	return u.ID
}

func (u *User) SetID(id interface{}) {
	u.ID = id.(string)
}
