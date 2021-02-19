package bolt

import (
	bh "github.com/timshannon/bolthold"
	bolt "go.etcd.io/bbolt"
)

func queryFromFilters(filters map[string]interface{}) *bh.Query {
	var query *bh.Query

	for key, value := range filters {
		if query == nil {
			query = bh.Where(key).Eq(value).Index(key)
		} else {
			query = query.And(key).Eq(value).Index(key)
		}
	}

	return query
}

func (store *boltStore) ensureTransaction(task func(*boltStore) error) error {
	if store.IsTransactional() {
		return task(store)
	}

	return store.bh.Bolt().Update(func(tx *bolt.Tx) error {
		return task(&boltStore{store.boltStoreConfig, tx})
	})
}

func (store *boltStore) nextSequence() (int, error) {
	bucket, err := store.tx.CreateBucketIfNotExists([]byte(store.typeName))
	if err != nil {
		return 0, err
	}

	id, err := bucket.NextSequence()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
