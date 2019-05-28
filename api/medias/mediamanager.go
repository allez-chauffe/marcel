package medias

// // CreateMedia Create a new Media, save it into memory and commit
// func CreateEmpty(owner string) *Media {

// 	log.Debugln("Creating media")

// 	newMedia := NewMedia()
// 	newMedia.ID = m.GetNextID()
// 	newMedia.Name = "Media " + strconv.Itoa(newMedia.ID)
// 	newMedia.Owner = owner

// 	//save it into the MediasConfiguration
// 	m.SaveIntoDB(newMedia)
// 	m.Commit()

// 	return newMedia
// }

// func Activate(media *Media) error {
// 	errorMessages := ""

// 	for _, mp := range media.Plugins {
// 		// duplicate plugin files into "medias/{idMedia}/{plugins_EltName}/{idInstance}"
// 		mpPath := m.GetPluginDirectory(media, mp.EltName, mp.InstanceId)
// 		if err := m.copyNewInstanceOfPlugin(media, &mp, mpPath); err != nil {
// 			log.Errorln(err.Error())
// 			//Don't return an error now, we need to activate the other plugins
// 			errorMessages += err.Error() + "\n"
// 		}
// 	}

// 	media.IsActive = true

// 	m.SaveIntoDB(media)

// 	if errorMessages != "" {
// 		return errors.New(errorMessages)
// 	}

// 	return nil
// }

// func Deactivate(media *Media) error {

// 	media.IsActive = false

// 	m.SaveIntoDB(media)

// 	return nil
// }

// func Delete(media *Media) error {

// 	m.Deactivate(media)

// 	m.RemoveFromDB(media)
// 	m.Commit()

// 	//remove plugins files
// 	err := os.RemoveAll(filepath.Join("medias", strconv.Itoa(media.ID)))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func copyNewInstanceOfPlugin(media *Media, mp *MediaPlugin, path string) error {
// 	// Copy only frontend dir since it is the only relevant files
// 	if err := commons.CopyDir(filepath.Join("plugins", mp.EltName, "frontend"), filepath.Join(path, "frontend")); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func GetPluginDirectory(media *Media, eltName string, instanceId string) string {
// 	return filepath.Join("medias", strconv.Itoa(media.ID), eltName, instanceId)
// }
