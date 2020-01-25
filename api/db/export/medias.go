package export

import (
	"github.com/Zenika/marcel/api/db/medias"
)

func Medias(pretty bool, outputFile string) error {
	return export(func() (interface{}, error) {
		return medias.List()
	}, pretty, outputFile)
}
