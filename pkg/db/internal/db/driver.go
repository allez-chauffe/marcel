package db

type Driver interface {
	Open() (Client, error)
}
