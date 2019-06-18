package imp0rt

import "github.com/Zenika/marcel/api/db/medias"

func Medias(inputFile string) error {
	var mediasList []medias.Media

	return imp0rt(inputFile, &mediasList, func() error {
		return medias.UpsertAll(mediasList)
	})
}
