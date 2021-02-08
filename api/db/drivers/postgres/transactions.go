package postgres

import "github.com/allez-chauffe/marcel/api/db/internal/db"

type postgresTransaction struct {

}

func (tx postgresTransaction) Get(id interface{}, result interface{}) error {
	panic("not implemented") // TODO: Implement
}

func (tx postgresTransaction) Exists(id interface{}) (bool, error) {
	panic("not implemented") // TODO: Implement
}

func (tx postgresTransaction) List(result interface{}) error {
	panic("not implemented") // TODO: Implement
}

func (tx postgresTransaction) Find(result interface{}, filters map[string]interface{}) error {
	panic("not implemented") // TODO: Implement
}

func (tx postgresTransaction) Insert(item db.Entity) error {
	panic("not implemented") // TODO: Implement
}

func (tx postgresTransaction) Update(item db.Entity) error {
	panic("not implemented") // TODO: Implement
}

func (tx postgresTransaction) Upsert(item db.Entity) error {
	panic("not implemented") // TODO: Implement
}

func (tx postgresTransaction) Delete(id interface{}) error {
	panic("not implemented") // TODO: Implement
}

func (tx postgresTransaction) DeleteAll() error {
	panic("not implemented") // TODO: Implement
}

func (tx postgresTransaction) Begin() (db.Transaction, error) {
	panic("not implemented") // TODO: Implement
}

func (tx postgresTransaction) Commit() error {
	panic("not implemented") // TODO: Implement
}

func (tx postgresTransaction) Rollback() error {
	panic("not implemented") // TODO: Implement
}

