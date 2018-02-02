package commons

import (
	"encoding/json"
	"log"
	"os"
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
	log.Println(err)
	Check(err)
}

func Commit(manager Manager) error {
	configFullPath, _, configFileName := manager.GetSaveFilePath()
	f, err := OpenSaveFile(manager, os.O_WRONLY)
	defer f.Close()

	if err != nil {
		log.Printf("Unable to open configuration file %s (%s) : %s", configFileName, configFullPath, err)
		return err
	}

	if err = json.NewEncoder(f).Encode(manager.GetConfig()); err != nil {
		log.Printf("Unable to write in configuration file %s (%s) : %s", configFileName, configFullPath, err)
	}

	return err
}

func OpenSaveFile(manager Manager, osFlag int) (*os.File, error) {
	configFullPath, _, _ := manager.GetSaveFilePath()

	return os.OpenFile(configFullPath, osFlag|os.O_CREATE|os.O_TRUNC, 0644)
}
