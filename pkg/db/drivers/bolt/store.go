package bolt

import (
	"reflect"

	bh "github.com/timshannon/bolthold"
	bolt "go.etcd.io/bbolt"

	"github.com/allez-chauffe/marcel/pkg/db/internal/db"
)

type store struct {
	*client
	entityFactory func() db.Entity
	typeName      string
	tx            *bolt.Tx
}

func (s *store) WithTransaction(tx db.Transaction) db.StoreBase {
	if tx == nil {
		return s
	}

	boltTx, ok := tx.(*bolt.Tx)
	if !ok {
		panic("Bolt driver received non bolt transaction (*bbolt.Tx required)")
	}

	return &store{s.client, s.entityFactory, s.typeName, boltTx}
}

func (s *store) IsWithTransaction() bool {
	return s.tx != nil
}

func (s *store) Get(id interface{}, result interface{}) error {
	var err error

	resultType := reflect.TypeOf(result)
	if resultType.Kind() != reflect.Ptr ||
		resultType.Elem().Kind() != reflect.Ptr &&
			resultType.Elem().Kind() != reflect.Interface {
		panic("result should be a pointer of pointer of the targeted entity (**Client by example)")
	}

	resultPointer := reflect.ValueOf(result).Elem()

	if s.IsWithTransaction() {
		err = s.bh.TxGet(s.tx, id, resultPointer.Interface())
	} else {
		err = s.bh.Get(id, resultPointer.Interface())
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

func (s *store) List(result interface{}) error {
	if s.IsWithTransaction() {
		return s.bh.TxFind(s.tx, result, nil)
	}
	return s.bh.Find(result, nil)
}

func (s *store) Insert(item db.Entity) error {
	// FIXME no need for that, see FIXME on s.nextSequence()
	return s.ensureTransaction(func(s *store) error {
		if db.ShouldAutoIncrement(item) {
			id, err := s.nextSequence()
			if err != nil {
				return err
			}
			item.SetID(id)
		}

		return s.bh.TxInsert(s.tx, item.GetID(), item)
	})
}

func (s *store) Update(item db.Entity) error {
	if s.IsWithTransaction() {
		return s.bh.TxUpdate(s.tx, item.GetID(), item)
	}
	return s.bh.Update(item.GetID(), item)
}

func (s *store) Delete(id interface{}) error {
	if s.IsWithTransaction() {
		return s.bh.TxDelete(s.tx, id, s.entityFactory())
	}
	return s.bh.Delete(id, s.entityFactory())
}

func (s *store) DeleteAll() error {
	if s.IsWithTransaction() {
		return s.bh.TxDeleteMatching(s.tx, s.entityFactory(), nil)
	}
	return s.bh.DeleteMatching(s.entityFactory(), nil)
}

func (s *store) Upsert(item db.Entity) error {
	return s.ensureTransaction(func(s *store) error {
		if db.ShouldAutoIncrement(item) {
			id, err := s.nextSequence()
			if err != nil {
				return err
			}
			item.SetID(id)
		}

		return s.bh.TxUpsert(s.tx, item.GetID(), item)
	})
}

func (s *store) Exists(id interface{}) (bool, error) {
	entity := s.entityFactory()
	if err := s.Get(id, &entity); err != nil {
		return false, err
	}

	return entity != nil, nil
}

func (s *store) Find(result interface{}, filters map[string]interface{}) error {
	query := queryFromFilters(filters)
	if s.IsWithTransaction() {
		return s.bh.TxFind(s.tx, result, query)
	}
	return s.bh.Find(result, query)
}

func (s *store) ensureTransaction(task func(*store) error) error {
	if s.IsWithTransaction() {
		return task(s)
	}

	return s.bh.Bolt().Update(func(tx *bolt.Tx) error {
		return task(&store{s.client, s.entityFactory, s.typeName, tx})
	})
}

// FIXME use bh.NextSequence() to give to bh.Insert()
func (s *store) nextSequence() (int, error) {
	// FIXME not really the right place
	bucket, err := s.tx.CreateBucketIfNotExists([]byte(s.typeName))
	if err != nil {
		return 0, err
	}

	id, err := bucket.NextSequence()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
