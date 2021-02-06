package imp0rt

import (
	"github.com/allez-chauffe/marcel/api/db"
	"github.com/allez-chauffe/marcel/api/db/plugins"
)

func Plugins(inputFile string) error {
	var pluginsList []plugins.Plugin

	return imp0rt(inputFile, &pluginsList, func() error {
		return db.Plugins().UpsertAll(pluginsList)
	})
}
