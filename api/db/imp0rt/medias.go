package imp0rt

import (
	"github.com/allez-chauffe/marcel/api/db"
	"github.com/allez-chauffe/marcel/api/db/medias"
)

func Medias(inputFile string) error {
	var mediasList []medias.Media

	return imp0rt(inputFile, &mediasList, func() error {
		return db.Medias().UpsertAll(mediasList)
	})
}
