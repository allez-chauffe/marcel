package bolt

import (
	"github.com/allez-chauffe/marcel/api/db/internal/db"
	bh "github.com/timshannon/bolthold"
)

type boltStore struct {
	*boltDatabase
	newEntity func() db.Entity
}

func (database *boltDatabase) CreateStore(newEntity func() db.Entity) db.Store {
	return &boltStore{ database, newEntity }
}

func (store *boltStore) Get(id interface{}, result interface{}) error {
	err := store.bh.Get(id, result)
	if err != nil {
		result = nil
		if err == bh.ErrNotFound {
			return nil
		}
		return err
	}

	return nil
}

func (store *boltStore) List(result interface{}) error {
	return store.bh.Find(result, nil)
}

func (store *boltStore) Insert(item db.Entity) error {
	if db.ShouldAutoIncrement(item) {
		return store.bh.Insert(bh.NextSequence(), item)
	}

	return store.bh.Insert(item.GetID(), item)
}

func (store *boltStore) Update(item db.Entity) error {
	return store.bh.Update(item.GetID(), item)
}

func (store *boltStore) Delete(id interface{}) error {
	return store.bh.Delete(id, store.newEntity())
}

func (store *boltStore) DeleteAll() error {
	return store.bh.DeleteMatching(store.newEntity(), nil)
}

func (store *boltStore) Upsert(item db.Entity) error {
	if db.ShouldAutoIncrement(item) {
		return store.Insert(item)
	}
	return store.bh.Upsert(item.GetID(), item)
}

func (store *boltStore) Exists(id interface{}) (bool, error) {
	entity := store.newEntity()
	if err := store.Get(id, entity); err != nil {
		return false, err
	}

	return entity != nil, nil
}

func (store *boltStore) Find(result interface{}, filters map[string]interface{}) error {


	return store.bh.Find(result, queryFromFilters(filters))
}
