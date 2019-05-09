package users

import (
	"time"

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/timshannon/bolthold"
	bolt "go.etcd.io/bbolt"

	"github.com/Zenika/MARCEL/api/db/internal/db"
)

func EnsureOneUser() error {
	return db.Store.Bolt().Update(func(tx *bolt.Tx) error {
		res, err := db.Store.TxFindAggregate(tx, &User{}, nil)
		if err != nil {
			return err
		}

		if res[0].Count() != 0 {
			return nil
		}

		log.Info("No users, creating an admin user...")

		u := &User{
			DisplayName: "Admin",
			Login:       "admin",
		}

		// FIXME generate password
		if err := u.SetPassword("admin"); err != nil {
			return err
		}

		return db.Store.TxInsert(tx, uuid.NewV4().String(), u)
	})
}

func Insert(u *User) error {
	if u.Role == "" {
		u.Role = "user"
	}
	u.CreatedAt = time.Now()

	return db.Store.Insert(uuid.NewV4().String(), u)
}

func List() ([]User, error) {
	users := []User{}

	return users, db.Store.Find(&users, nil)
}

func Get(id string) (*User, error) {
	u := &User{}

	return u, db.Store.Get(id, &u)
}

func GetByLogin(login string) (*User, error) {
	var users []User

	err := db.Store.Find(&users, bolthold.Where("Login").Eq(login))
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	return &users[0], nil
}

func Delete(id string) error {
	return db.Store.Delete(id, &User{})
}