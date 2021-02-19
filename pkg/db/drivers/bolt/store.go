package bolt

import (
	"reflect"

	bh "github.com/timshannon/bolthold"
	bolt "go.etcd.io/bbolt"

	"github.com/allez-chauffe/marcel/pkg/db/internal/db"
)

type boltStoreConfig struct {
	*boltDatabase
	newEntity func() db.Entity
	typeName  string
}

type boltStore struct {
	*boltStoreConfig
	tx *bolt.Tx
}

func (store *boltStore) Transactional(tx db.Transaction) db.Store {
	if tx == nil {
		return store
	}

	boltTx, ok := tx.(*bolt.Tx)

	if !ok {
		panic("Bolt driver received non bolt transaction (*bbolt.Tx required)")
	}

	return &boltStore{store.boltStoreConfig, boltTx}
}

func (store *boltStore) IsTransactional() bool {
	return store.tx != nil
}

func (store *boltStore) Get(id interface{}, result interface{}) error {
	var err error

	resultType := reflect.TypeOf(result)
	if resultType.Kind() != reflect.Ptr ||
		resultType.Elem().Kind() != reflect.Ptr &&
			resultType.Elem().Kind() != reflect.Interface {
		panic("result should be a pointer of pointer of the targeted entity (**Client by example)")
	}

	resultPointer := reflect.ValueOf(result).Elem()

	if store.IsTransactional() {
		err = store.bh.TxGet(store.tx, id, resultPointer.Interface())
	} else {
		err = store.bh.Get(id, resultPointer.Interface())
	}

	if err != nil {
		resultPointer.Set(reflect.Zero(resultPointer.Type()))
		if err == bh.ErrNotFound {
			return nil
		}
		return err
	}

	return nil
}

func (store *boltStore) List(result interface{}) error {
	if store.IsTransactional() {
		return store.bh.TxFind(store.tx, result, nil)
	}
	return store.bh.Find(result, nil)
}

func (store *boltStore) Insert(item db.Entity) error {
	return store.ensureTransaction(func(store *boltStore) error {
		if db.ShouldAutoIncrement(item) {
			id, err := store.nextSequence()
			if err != nil {
				return err
			}
			item.SetID(id)
		}

		return store.bh.TxInsert(store.tx, item.GetID(), item)
	})
}

func (store *boltStore) Update(item db.Entity) error {
	if store.IsTransactional() {
		return store.bh.TxUpdate(store.tx, item.GetID(), item)
	}
	return store.bh.Update(item.GetID(), item)
}

func (store *boltStore) Delete(id interface{}) error {
	if store.IsTransactional() {
		return store.bh.TxDelete(store.tx, id, store.newEntity())
	}
	return store.bh.Delete(id, store.newEntity())
}

func (store *boltStore) DeleteAll() error {
	if store.IsTransactional() {
		return store.bh.TxDeleteMatching(store.tx, store.newEntity(), nil)
	}
	return store.bh.DeleteMatching(store.newEntity(), nil)
}

func (store *boltStore) Upsert(item db.Entity) error {
	return store.ensureTransaction(func(bs *boltStore) error {
		if db.ShouldAutoIncrement(item) {
			id, err := store.nextSequence()
			if err != nil {
				return err
			}
			item.SetID(id)
		}

		return store.bh.TxUpsert(store.tx, item.GetID(), item)
	})
}

func (store *boltStore) Exists(id interface{}) (bool, error) {
	entity := store.newEntity()
	if err := store.Get(id, &entity); err != nil {
		return false, err
	}

	return entity != nil, nil
}

func (store *boltStore) Find(result interface{}, filters map[string]interface{}) error {
	query := queryFromFilters(filters)
	if store.IsTransactional() {
		return store.bh.TxFind(store.tx, result, query)
	}
	return store.bh.Find(result, query)
}
