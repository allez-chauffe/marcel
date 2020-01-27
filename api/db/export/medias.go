package export

import (
	"github.com/Zenika/marcel/api/db/medias"
)

func Medias(outputFile string, pretty bool) error {
	return export(func() (interface{}, error) {
		return medias.List()
	}, outputFile, pretty)
}
