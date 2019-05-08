package users

import (
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/timshannon/bolthold"
	bolt "go.etcd.io/bbolt"

	"github.com/Zenika/MARCEL/api/db/internal/db"
	"github.com/Zenika/MARCEL/api/user"
)

func EnsureOneUser() error {
	return db.Store.Bolt().Update(func(tx *bolt.Tx) error {
		res, err := db.Store.TxFindAggregate(tx, &user.User{}, nil)
		if err != nil {
			return err
		}

		if res[0].Count() != 0 {
			return nil
		}

		log.Info("No users, creating an admin user...")

		// FIXME generate password
		u, err := user.New("Admin", "admin", "admin", "admin")
		if err != nil {
			return err
		}

		return db.Store.TxInsert(tx, uuid.NewV4().String(), u)
	})
}

func Insert(u *user.User) error {
	return db.Store.Insert(uuid.NewV4().String(), u)
}

func List() ([]user.User, error) {
	users := []user.User{}

	return users, db.Store.Find(&users, nil)
}

func Get(id string) (*user.User, error) {
	u := &user.User{}

	return u, db.Store.Get(id, &u)
}

func GetByLogin(login string) (*user.User, error) {
	var users []user.User

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
	return db.Store.Delete(id, &user.User{})
}
