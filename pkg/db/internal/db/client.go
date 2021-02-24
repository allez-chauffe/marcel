package db

type Client interface {
	Store(entityFactory func() Entity) (Store, error)
	Begin() (Transaction, error)
	Close() error
}
