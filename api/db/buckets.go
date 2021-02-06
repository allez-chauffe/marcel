package db

import (
	"github.com/allez-chauffe/marcel/api/db/clients"
	"github.com/allez-chauffe/marcel/api/db/internal/db"
	"github.com/allez-chauffe/marcel/api/db/medias"
	"github.com/allez-chauffe/marcel/api/db/plugins"
	"github.com/allez-chauffe/marcel/api/db/users"
)

func Clients() *clients.Bucket {
	return clients.DefaultBucket
}

func Medias() *medias.Bucket {
	return medias.DefaultBucket
}

func Plugins() *plugins.Bucket {
	return plugins.DefaultBucket
}

func Users() *users.Bucket {
	return users.DefaultBucket
}

type Tx struct {
	db.Transaction
}

func (tx *Tx) Clients() *clients.Bucket {
	return clients.Transactional(tx.Transaction)
}

func (tx *Tx) Medias() *medias.Bucket {
	return medias.Transactional(tx.Transaction)
}

func (tx *Tx) Plugins() *plugins.Bucket {
	return plugins.Transactional(tx.Transaction)
}

func (tx *Tx) Users() *users.Bucket {
	return users.Transactional(tx.Transaction)
}
