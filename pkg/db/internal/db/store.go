package db

import "errors"

type StoreBase interface {
	Get(id interface{}, result interface{}) error

	Exists(id interface{}) (bool, error) // Get + check err

	List(result interface{}) error // FIXME find w/ nil filters

	Find(filters map[string]interface{}, result interface{}) error

	Insert(item Entity) error

	Update(item Entity) error

	Upsert(item Entity) error // FIXME Get + Insert/Update

	Delete(id interface{}) error

	DeleteAll() error // FIXME List + Delete

	WithTransaction(tx Transaction) StoreBase

	IsWithTransaction() bool
}

type Store interface {
	Get(id interface{}) (interface{}, error)

	Exists(id interface{}) (bool, error) // Get + check err

	List() (interface{}, error)

	Find(filters map[string]interface{}) (interface{}, error)

	Insert(item Entity) error

	Update(item Entity) error

	Upsert(item Entity) error // FIXME Get + Insert/Update

	Delete(id interface{}) error

	DeleteAll() error // FIXME List + Delete

	// FIXME Transactional()
}

type Entity interface {
	GetID() interface{}
	SetID(id interface{})
}

var EntityNotFound = errors.New("Entity not found")

func ShouldAutoIncrement(entity Entity) bool {
	// FIXME this shouldn't be determined on ID value, but rather set as an option on the store
	id, isInt := entity.GetID().(int)

	return isInt && id == -1
}
