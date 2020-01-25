package export

import (
	"github.com/Zenika/marcel/api/db/plugins"
)

func Plugins(pretty bool, outputFile string) error {
	return export(func() (interface{}, error) {
		return plugins.List()
	}, pretty, outputFile)
}
