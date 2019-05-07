package user

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"time"

	uuid "github.com/satori/go.uuid"
)

// FIXME something ?
var secret = []byte("This is the password secret key !")

type User struct {
	ID               string    `json:"id"`
	DisplayName      string    `json:"displayName"`
	Login            string    `json:"login"`
	Role             string    `json:"role"`
	CreatedAt        time.Time `json:"createdAt"`
	LastDisconection time.Time `json:"lastDisconnection"`
	PasswordHash     string    `json:"passwordHash"`
	PasswordSalt     string    `json:"passwordSalt"`
}

func New(displayName, login, pRole, password string) (*User, error) {
	role := pRole
	if role == "" {
		role = "user"
	}

	hash, salt, err := generateHash(password)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:           uuid.NewV4().String(),
		DisplayName:  displayName,
		Login:        login,
		Role:         role,
		CreatedAt:    time.Now(),
		PasswordHash: hash,
		PasswordSalt: salt,
	}, nil
}

func FromValue(value []byte) (*User, error) {
	user := &User{}

	return user, json.NewDecoder(bytes.NewReader(value)).Decode(user)
}

func (u *User) Key() []byte {
	return []byte(u.ID)
}

func (u *User) Value() ([]byte, error) {
	b := bytes.NewBuffer([]byte{})

	err := json.NewEncoder(b).Encode(u)

	return b.Bytes(), err
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
