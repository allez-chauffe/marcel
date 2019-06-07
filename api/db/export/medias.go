package export

import (
	"github.com/Zenika/MARCEL/api/db/medias"
)

func Medias(outputFile string) error {
	return export(func() (interface{}, error) {
		return medias.List()
	}, outputFile)
}
