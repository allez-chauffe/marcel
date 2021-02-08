package users

import (
	"time"

	rand "github.com/Pallinder/go-randomdata"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"

	"github.com/allez-chauffe/marcel/api/db/internal/db"
)

var DefaultStore *Store

type Store struct {
	store db.Store
}

func CreateStore() error {
	store, err := db.DB.CreateStore(func() db.Entity {
		return new(User)
	})

	if err != nil {
		return err
	}

	DefaultStore = &Store{store}
	return nil
}

func Transactional(tx db.Transaction) *Store {
	return &Store{DefaultStore.store.Transactional(tx)}
}

func (b *Store) EnsureOneUser() error {
	return db.EnsureTransaction(b.store, func(store db.Store) error {
		users := &[]User{}

		if err := store.List(users); err != nil || len(*users) != 0 {
			return err
		}

		log.Info("No users in database, creating admin...")

		u := &User{
			ID:          uuid.NewV4().String(),
			DisplayName: "Admin",
			Login:       "admin",
			Role:        "admin",
			CreatedAt:   time.Now(),
		}

		password := rand.RandStringRunes(10)
		if err := u.SetPassword(password); err != nil {
			return err
		}

		if err := store.Insert(u); err != nil {
			return err
		}

		log.Infof("User admin created with password %s", password)

		return nil
	})
}

func (b *Store) Insert(u *User) error {
	return b.store.Insert(u)
}

func (b *Store) List() ([]User, error) {
	var users []User
	return users, b.store.List(&users)
}

func (b *Store) Get(id string) (*User, error) {
	u := new(User)
	return u, b.store.Get(id, &u)
}

func (b *Store) GetByLogin(login string) (*User, error) {
	var users []User

	filters := map[string]interface{}{
		"Login": login,
	}

	if err := b.store.Find(&users, filters); err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	return &users[0], nil
}

func (b *Store) Delete(id string) error {
	return b.store.Delete(id)
}

func (b *Store) Disconnect(id string) error {
	return db.EnsureTransaction(b.store, func(store db.Store) error {
		u := &User{}
		if err := store.Get(id, &u); u == nil || err != nil {
			return err
		}

		u.LastDisconnection = time.Now()
		return store.Update(u)
	})
}

func (b *Store) Update(user *User) error {
	return b.store.Update(user)
}

func (b *Store) UpsertAll(users []User) error {
	return db.EnsureTransaction(b.store, func(store db.Store) error {
		for _, u := range users {
			if err := store.Upsert(&u); err != nil {
				return err
			}
		}

		return nil
	})
}

func (b *Store) Exists(id string) (bool, error) {
	return b.store.Exists(id)
}
