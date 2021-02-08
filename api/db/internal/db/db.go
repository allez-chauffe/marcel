package db

import "errors"

var DB Database

type Store interface {
	Get(id interface{}, result interface{}) error

	Exists(id interface{}) (bool, error)

	List(result interface{}) error

	Find(result interface{}, filters map[string]interface{}) error

	Insert(item Entity) error

	Update(item Entity) error

	Upsert(item Entity) error

	Delete(id interface{}) error

	DeleteAll() error

	Transactional(tx Transaction) Store

	IsTransactional() bool
}

type Database interface {
	CreateStore(newItem func() Entity) (Store, error)
	Open() error
	Close() error
	Begin() (Transaction, error)
}

type Entity interface {
	GetID() interface{}
	SetID(id interface{})
}

var EntityNotFound = errors.New("Entity not found")

func ShouldAutoIncrement(entity Entity) bool {
	id, isInt := entity.GetID().(int)

	return isInt && id == -1
}
