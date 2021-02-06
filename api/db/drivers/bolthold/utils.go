package bolt

import (
	log "github.com/sirupsen/logrus"
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
	log.Debug("ensure start")
	if store.IsTransactional() {
		log.Debug("is transactional")
		return task(store)
	}

	log.Debug("is not transactional")

	return store.bh.Bolt().Update(func(tx *bolt.Tx) error {
		return task(&boltStore{store.boltStoreConfig, tx})
	})
}

func (store *boltStore) nextSequence() (int, error) {
	bucket := store.tx.Bucket([]byte(store.typeName))
	// If it is the first time we save this type of entity, the bucket doesn't exists
	if bucket == nil {
		return 0, nil
	}

	id, err := bucket.NextSequence()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
