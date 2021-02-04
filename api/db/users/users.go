package users

import (
	"time"

	rand "github.com/Pallinder/go-randomdata"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"

	"github.com/allez-chauffe/marcel/api/db/internal/db"
)

var store db.Store

func CreateStore(database db.Databse) {
	store = database.CreateStore(func() db.Entity {
		return new(User)
	})
}

func EnsureOneUser() error {
	return db.Transactional(store, func(tx db.Transaction) error {
		users := &[]User{}

		if err := tx.List(users); err != nil || len(*users) != 0 {
			return err
		}

		log.Info("No users in database, creating admin...")

		u := &User{
			DisplayName: "Admin",
			Login:       "admin",
			Role:        "admin",
			CreatedAt:   time.Now(),
		}

		password := rand.RandStringRunes(10)
		if err := u.SetPassword(password); err != nil {
			return err
		}

		if err := insert(tx, u); err != nil {
			return err
		}

		log.Infof("User admin created with password %s", password)

		return nil
	})
}

func Insert(u *User) error {
	return insert(store, u)
}

func insert(store db.Store, u *User) error {
	u.ID = uuid.NewV4().String()
	if u.Role == "" {
		u.Role = "user"
	}
	u.CreatedAt = time.Now()

	return store.Insert(u)
}

func List() ([]User, error) {
	var users []User
	return users, store.List(&users)
}

func Get(id string) (*User, error) {
	u := new(User)
	return u, store.Get(id, u)
}

func GetByLogin(login string) (*User, error) {
	var users []User

	filters := map[string]interface{}{
		"Login": login,
	}

	if err := store.Find(&users, filters); err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	return &users[0], nil
}

func Delete(id string) error {
	return store.Delete(id)
}

func Disconnect(id string) error {
	return db.Transactional(store, func(tx db.Transaction) error {
		u := &User{}
		if err := tx.Get(id, u);  u == nil || err != nil {
			return err
		}

		u.LastDisconnection = time.Now()
		return tx.Update(u)
	})
}

func Update(user *User) error {
	return store.Update(user)
}

func UpsertAll(users []User) error {
	return db.Transactional(store, func(tx db.Transaction) error {
		for _, u := range users {
			if err := tx.Upsert(&u); err != nil {
				return err
			}
		}

		return nil
	})
}
