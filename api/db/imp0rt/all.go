package imp0rt

import (
	"github.com/allez-chauffe/marcel/api/db"
	"github.com/allez-chauffe/marcel/api/db/medias"
	"github.com/allez-chauffe/marcel/api/db/plugins"
)

type all struct {
	Users   []userPassword   `json:"users"`
	Medias  []medias.Media   `json:"medias"`
	Plugins []plugins.Plugin `json:"plugins"`
}

func All(inputFile string) error {
	var data all

	return imp0rt(inputFile, &data, func() error {
		return db.Transactional(func(tx *db.Tx) error {
			if err := importUsers(data.Users); err != nil {
				return err
			}

			if err := tx.Medias().UpsertAll(data.Medias); err != nil {
				return err
			}

			return tx.Plugins().UpsertAll(data.Plugins)
		})
	})
}
