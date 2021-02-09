package export

import "github.com/allez-chauffe/marcel/api/db"

func Plugins(outputFile string, pretty bool) error {
	return export(func() (interface{}, error) {
		return db.Plugins().List()
	}, outputFile, pretty)
}
