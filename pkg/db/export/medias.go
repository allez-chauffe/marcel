package export

import (
	"github.com/allez-chauffe/marcel/pkg/db"
)

func Medias(outputFile string, pretty bool) error {
	return export(func() (interface{}, error) {
		return db.Medias().List()
	}, outputFile, pretty)
}
