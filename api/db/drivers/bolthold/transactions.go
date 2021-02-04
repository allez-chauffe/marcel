package bolt

import (
	"github.com/allez-chauffe/marcel/api/db/internal/db"
	bh "github.com/timshannon/bolthold"

	bolt "go.etcd.io/bbolt"
)

type boltTransaction struct {
	bh *bolt.Tx
	store boltStore
}

func (store boltStore) Begin() (db.Transaction, error) {
	tx, err := store.bh.Bolt().Begin(true)
	if err != nil {
		return nil, err
	}

	return &boltTransaction{tx, store}, nil
}

func (tx *boltTransaction) Commit() error {
	return tx.bh.Commit()
}

func (tx *boltTransaction) Rollback() error {
	return tx.bh.Rollback()
}

func (tx *boltTransaction) Get(id interface{}, result interface{}) error {
	err := tx.store.bh.TxGet(tx.bh, id, result)
	if err != nil {
		result = nil
		if err == bh.ErrNotFound {
			return nil
		}
		return err
	}

	return nil
}

func (tx *boltTransaction) List(result interface{}) error {
	return tx.store.bh.TxFind(tx.bh, result, nil)
}

func (tx *boltTransaction) Insert(item db.Entity) error {
	if db.ShouldAutoIncrement(item) {
		return tx.store.bh.TxInsert(tx.bh, bh.NextSequence(), item)
	}
	return tx.store.bh.TxInsert(tx.bh, item.GetID(), item)
}

func (tx *boltTransaction) Update(item db.Entity) error {
	return tx.store.bh.TxUpdate(tx.bh, item.GetID(), item)
}

func (tx *boltTransaction) Delete(id interface{}) error {
	return tx.store.bh.TxDelete(tx.bh, id, tx.store.newEntity())
}

func (tx *boltTransaction) DeleteAll() error {
	return tx.store.bh.TxDeleteMatching(tx.bh, tx.store.newEntity(), nil)
}

func (tx *boltTransaction) Upsert(item db.Entity) error {
	return tx.store.bh.TxUpsert(tx.bh, item.GetID(), item)
}

func (tx *boltTransaction) Begin() (db.Transaction, error) {
	return tx, nil
}

func (tx *boltTransaction) Exists(id interface{}) (bool, error) {
	entity := tx.store.newEntity()
	if err := tx.Get(id, entity); err != nil {
		return false, err
	}

	return entity != nil, nil
}

func (tx *boltTransaction) Find(result interface{}, filters map[string]interface{}) error {
	return tx.store.bh.TxFind(tx.bh, result, queryFromFilters(filters))
}
