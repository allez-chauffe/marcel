package users

import (
	"sync"

	log "github.com/sirupsen/logrus"
	bolt "go.etcd.io/bbolt"

	"github.com/Zenika/MARCEL/api/db/internal/db"
	"github.com/Zenika/MARCEL/api/user"
)

var (
	usersBucketName       = []byte("users")
	ensureUsersBucketOnce sync.Once
)

func ensureUsersBucket() {
	ensureUsersBucketOnce.Do(func() {
		err := db.DB.Update(func(tx *bolt.Tx) error {
			b, err := tx.CreateBucketIfNotExists(usersBucketName)

			if b.Stats().KeyN == 0 {
				log.Info("No users, creating an admin user...")

				u, err := user.New("Admin", "admin", "admin", "admin")
				if err != nil {
					return err
				}

				if err := put(tx, u); err != nil {
					return err
				}
			}

			return err
		})

		if err != nil {
			log.Errorf("Users bucket initialization failed: %s", err)
		}
	})
}

func put(tx *bolt.Tx, u *user.User) error {
	value, err := u.Value()
	if err != nil {
		return err
	}

	return tx.Bucket(usersBucketName).Put(u.Key(), value)
}

func Put(u *user.User) error {
	ensureUsersBucket()

	return db.DB.Update(func(tx *bolt.Tx) error {
		return put(tx, u)
	})
}

func List() ([]*user.User, error) {
	ensureUsersBucket()

	var users []*user.User

	err := db.DB.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(usersBucketName).Cursor()

		for {
			_, value := c.Next()
			if value == nil {
				break
			}

			u, err := user.FromValue(value)
			if err != nil {
				return err
			}

			users = append(users, u)
		}

		return nil
	})

	return users, err
}

func Get(id string) (*user.User, error) {
	ensureUsersBucket()

	var u *user.User

	return u, db.DB.View(func(tx *bolt.Tx) error {
		value := tx.Bucket(usersBucketName).Get([]byte(id))
		if value == nil {
			return nil
		}

		var err error
		u, err = user.FromValue(value)

		return err
	})
}

func GetByLogin(login string) (*user.User, error) {
	ensureUsersBucket()

	var u *user.User

	return u, db.DB.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(usersBucketName).Cursor()

		for _, v := c.First(); v != nil; _, v = c.Next() {
			uu, err := user.FromValue(v)
			if err != nil {
				return err
			}

			if uu.Login == login {
				u = uu
				return nil
			}
		}

		return nil
	})
}

func Delete(id string) error {
	ensureUsersBucket()

	return db.DB.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(usersBucketName).Delete([]byte(id))
	})
}
