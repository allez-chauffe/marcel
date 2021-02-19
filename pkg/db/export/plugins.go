package export

import "github.com/allez-chauffe/marcel/pkg/db"

func Plugins(outputFile string, pretty bool) error {
	return export(func() (interface{}, error) {
		return db.Plugins().List()
	}, outputFile, pretty)
}
