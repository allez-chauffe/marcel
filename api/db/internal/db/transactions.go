package db

type Transaction interface {
	Store
	Commit() error
	Rollback() error
}

func Transactional(store Store, task func(Transaction) error) (err error){
	tx, err := store.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if recoverd := recover(); err != nil && recoverd != nil {
			tx.Rollback()

			if recoverd != nil {
				panic(recoverd)
			}
		}
	}()

	if err = task(tx); err != nil {
		return
	}

	return tx.Commit()
}
