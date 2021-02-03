package export

import (
	"github.com/allez-chauffe/marcel/api/db/plugins"
)

func Plugins(outputFile string, pretty bool) error {
	return export(func() (interface{}, error) {
		return plugins.List()
	}, outputFile, pretty)
}
