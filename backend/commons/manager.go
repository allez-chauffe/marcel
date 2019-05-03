package commons

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
)

type Manager interface {
	Commit() error
	GetSaveFilePath() (fullPath string, dirPath string, fileName string)
	GetConfig() interface{}
}

func LoadFromDB(manager Manager) {
	f, err := OpenSaveFile(manager, os.O_RDONLY)
	defer f.Close()
	Check(err)
	err = json.NewDecoder(f).Decode(manager.GetConfig())
	Check(err)
}

func Commit(manager Manager) error {
	configFullPath, _, configFileName := manager.GetSaveFilePath()
	f, err := OpenSaveFile(manager, os.O_WRONLY|os.O_TRUNC)
	defer f.Close()

	if err != nil {
		log.Errorf("Unable to open configuration file %s (%s) : %s", configFileName, configFullPath, err)
		return err
	}

	if err = json.NewEncoder(f).Encode(manager.GetConfig()); err != nil {
		log.Errorf("Unable to write in configuration file %s (%s) : %s", configFileName, configFullPath, err)
	}

	return err
}

func OpenSaveFile(manager Manager, osFlag int) (*os.File, error) {
	configFullPath, _, _ := manager.GetSaveFilePath()

	return os.OpenFile(configFullPath, osFlag|os.O_CREATE, 0644)
}
