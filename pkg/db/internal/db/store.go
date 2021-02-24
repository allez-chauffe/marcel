package db

import "errors"

type StoreBase interface {
	Get(id interface{}, result interface{}) error

	Exists(id interface{}) (bool, error) // Get + check err

	List(result interface{}) error

	Find(result interface{}, filters map[string]interface{}) error

	Insert(item Entity) error

	Update(item Entity) error

	Upsert(item Entity) error // FIXME Get + Insert/Update

	Delete(id interface{}) error

	DeleteAll() error // FIXME List + Delete

	WithTransaction(tx Transaction) StoreBase

	IsWithTransaction() bool
}

type Store interface {
	StoreBase
}

type Entity interface {
	GetID() interface{}
	SetID(id interface{})
}

var EntityNotFound = errors.New("Entity not found")

func ShouldAutoIncrement(entity Entity) bool {
	id, isInt := entity.GetID().(int)

	// FIXME use a special autoincrement type like bolthold does

	return isInt && id == -1
}
