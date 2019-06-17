package export

import (
	"github.com/Zenika/MARCEL/api/db/plugins"
)

func Plugins(outputFile string) error {
	return export(func() (interface{}, error) {
		return plugins.List()
	}, outputFile)
}
