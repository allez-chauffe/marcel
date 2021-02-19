package db

import (
	log "github.com/sirupsen/logrus"
)

type Transaction interface {
	Commit() error
	Rollback() error
}

func Transactional(task func(Transaction) error) (err error) {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	log.Debugf("begin transaction")

	defer FinishTransaction(tx, &err)

	return task(tx)
}

func FinishTransaction(tx Transaction, err *error) {
	r := recover()

	if *err == nil && r == nil {
		*err = tx.Commit()
		log.Debugf("commit transaction")
		return
	}

	if err := tx.Rollback(); err != nil {
		log.Errorf("Error while rolling back: %s", err)
	}
	log.Debugf("rollback transaction")

	if r != nil {
		panic(r)
	}
}

func EnsureTransaction(store Store, task func(Store) error) (err error) {
	if store.IsTransactional() {
		return task(store)
	}

	return Transactional(func(tx Transaction) error {
		return task(store.Transactional(tx))
	})
}
