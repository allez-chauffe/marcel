package export

import (
	"github.com/allez-chauffe/marcel/pkg/db"
	"github.com/allez-chauffe/marcel/pkg/db/medias"
	"github.com/allez-chauffe/marcel/pkg/db/plugins"
)

type all struct {
	Users   interface{}      `json:"users"`
	Medias  []medias.Media   `json:"medias"`
	Plugins []plugins.Plugin `json:"plugins"`
}

func All(usersWPassword bool, outputFile string, pretty bool) error {
	return export(func() (interface{}, error) {
		var result *all

		return result, db.Transactional(func(tx *db.Tx) error {
			users, err := listUsers(usersWPassword)
			if err != nil {
				return err
			}

			medias, err := tx.Medias().List()
			if err != nil {
				return err
			}

			plugins, err := tx.Plugins().List()
			if err != nil {
				return err
			}

			result = &all{
				Users:   users,
				Medias:  medias,
				Plugins: plugins,
			}

			return nil
		})
	}, outputFile, pretty)
}
