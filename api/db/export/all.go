package export

import (
	"github.com/Zenika/MARCEL/api/db/medias"
	"github.com/Zenika/MARCEL/api/db/plugins"
)

type all struct {
	Users   interface{}      `json:"users"`
	Medias  []medias.Media   `json:"medias"`
	Plugins []plugins.Plugin `json:"plugins"`
}

func All(usersWPassword bool, outputFile string) error {
	return export(func() (interface{}, error) {
		users, err := listUsers(usersWPassword)
		if err != nil {
			return nil, err
		}

		medias, err := medias.List()
		if err != nil {
			return nil, err
		}

		plugins, err := plugins.List()
		if err != nil {
			return nil, err
		}

		return &all{
			Users:   users,
			Medias:  medias,
			Plugins: plugins,
		}, nil
	}, outputFile)
}
