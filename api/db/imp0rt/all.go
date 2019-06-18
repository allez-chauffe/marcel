package imp0rt

import (
	"github.com/Zenika/marcel/api/db/medias"
	"github.com/Zenika/marcel/api/db/plugins"
)

type all struct {
	Users   []userPassword   `json:"users"`
	Medias  []medias.Media   `json:"medias"`
	Plugins []plugins.Plugin `json:"plugins"`
}

func All(inputFile string) error {
	var data all

	return imp0rt(inputFile, &data, func() error {
		if err := importUsers(data.Users); err != nil {
			return err
		}

		if err := medias.UpsertAll(data.Medias); err != nil {
			return err
		}

		return plugins.UpsertAll(data.Plugins)
	})
}
