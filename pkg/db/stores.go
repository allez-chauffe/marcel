package db

import (
	"github.com/allez-chauffe/marcel/pkg/db/clients"
	"github.com/allez-chauffe/marcel/pkg/db/internal/db"
	"github.com/allez-chauffe/marcel/pkg/db/medias"
	"github.com/allez-chauffe/marcel/pkg/db/plugins"
	"github.com/allez-chauffe/marcel/pkg/db/users"
)

func Clients() *clients.Store {
	return clients.DefaultStore
}

func Medias() *medias.Store {
	return medias.DefaultStore
}

func Plugins() *plugins.Store {
	return plugins.DefaultStore
}

func Users() *users.Store {
	return users.DefaultStore
}

type Tx struct {
	db.Transaction
}

func (tx *Tx) Clients() *clients.Store {
	return clients.Transactional(tx.Transaction)
}

func (tx *Tx) Medias() *medias.Store {
	return medias.Transactional(tx.Transaction)
}

func (tx *Tx) Plugins() *plugins.Store {
	return plugins.Transactional(tx.Transaction)
}

func (tx *Tx) Users() *users.Store {
	return users.Transactional(tx.Transaction)
}
