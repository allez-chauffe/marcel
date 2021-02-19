package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/allez-chauffe/marcel/pkg/db/internal/db"
)

type postgresStoreConfig struct {
	table     string
	idType    string
	newEntity func() db.Entity
	pg        *sql.DB
}

type postgresStore struct {
	*postgresStoreConfig
	tx *sql.Tx
}

func (store *postgresStore) Transactional(tx db.Transaction) db.Store {
	if tx == nil {
		return store
	}

	pgTx, ok := tx.(*sql.Tx)

	if !ok {
		panic("Postgres driver received non postgres transaction (*sql.Tx required)")
	}

	return &postgresStore{store.postgresStoreConfig, pgTx}
}

func (store *postgresStore) IsTransactional() bool {
	return store.tx != nil
}

func (store *postgresStore) selectQuery(query string) string {
	return fmt.Sprintf(`SELECT id, data FROM "%s" %s`, store.table, query)
}

func (store *postgresStore) Get(id interface{}, result interface{}) error {
	row := store.client().QueryRow(store.selectQuery("WHERE id = $1"), id)

	if err := store.unmarshallRow(row, result); err != nil {
		return err
	}

	return nil
}

func (store *postgresStore) Exists(id interface{}) (bool, error) {
	entity := store.newEntity()

	if err := store.Get(id, &entity); err != nil {
		return false, err
	}

	return entity != nil, nil
}

func (store *postgresStore) List(result interface{}) error {
	return store.Find(result, nil)
}

func (store *postgresStore) Find(result interface{}, filters map[string]interface{}) error {
	resultVal := reflect.ValueOf(result)
	if resultVal.Kind() != reflect.Ptr || resultVal.Elem().Kind() != reflect.Slice {
		panic("List result should be a slice pointer")
	}
	sliceVal := resultVal.Elem()

	if filters == nil {
		filters = map[string]interface{}{}
	}

	jsonFilters, err := json.Marshal(filters)
	if err != nil {
		return nil
	}

	rows, err := store.client().Query(store.selectQuery("WHERE data @> $1::jsonb"), jsonFilters)
	if err != nil {
		return err
	}

	for rows.Next() {
		entity := store.newEntity()
		if err := store.unmarshallRow(rows, &entity); err != nil {
			return err
		}
		sliceVal = reflect.Append(sliceVal, reflect.ValueOf(entity).Elem())
	}

	resultVal.Elem().Set(sliceVal.Slice(0, sliceVal.Len()))
	return nil
}

func (store *postgresStore) Insert(item db.Entity) error {
	id, data := prepare(item)

	if db.ShouldAutoIncrement(item) {
		row := store.client().QueryRow(fmt.Sprintf(`INSERT INTO "%s" (data) VALUES ($1) RETURNING id`, store.table), data)
		if err := row.Scan(&id); err != nil {
			return err
		}
		item.SetID(id)
	} else {
		if _, err := store.client().Exec(fmt.Sprintf(`INSERT INTO "%s" (id, data) VALUES ($1, $2)`, store.table), id, data); err != nil {
			return err
		}
	}

	return nil
}

func (store *postgresStore) Update(item db.Entity) error {
	id, data := prepare(item)

	result, err := store.client().Exec(fmt.Sprintf(`UPDATE "%s" SET data = $1 WHERE id = $2`, store.table), data, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected != 1 {
		return db.EntityNotFound
	}

	return nil
}

func (store *postgresStore) Upsert(item db.Entity) error {
	id, data := prepare(item)

	if db.ShouldAutoIncrement(item) {
		row := store.client().QueryRow(fmt.Sprintf(`INSERT INTO "%s" (data) VALUES ($1) RETURNING id`, store.table), data)
		if err := row.Scan(&id); err != nil {
			return err
		}
		item.SetID(id)
	} else {
		if _, err := store.client().Exec(fmt.Sprintf(`
				INSERT INTO "%s" (id, data) VALUES ($1, $2)
				ON CONFLICT (id) DO UPDATE SET data = $2
			`, store.table), id, data); err != nil {
			return err
		}
	}

	return nil
}

func (store *postgresStore) Delete(id interface{}) error {
	_, err := store.client().Exec(fmt.Sprintf(`DELETE FROM "%s" WHERE id = $1`, store.table), id)
	return err
}

func (store *postgresStore) DeleteAll() error {
	_, err := store.client().Exec(fmt.Sprintf(`DELETE FROM "%s"`, store.table))
	return err
}
