package commons

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/mitchellh/mapstructure"
)

type Manager interface {
	Commit() error
	GetSaveFilePath() (fullPath string, dirPath string, fileName string)
	GetConfig() interface{}
}

// CreateSaveFileIfNotExist check if the config file for the given manager exists and create it if not.
func CreateSaveFileIfNotExist(manager Manager) {
	fullPath, dirPath, fileName := manager.GetSaveFilePath()

	if !FileOrFolderExists(dirPath) {
		log.Println("Data directory did not exist. Create it.")
		os.Mkdir(dirPath, 0755)
	}

	if !FileOrFolderExists(fullPath) {

		f, err := os.Create(fullPath)
		Check(err)

		log.Printf("Configuration file %s created at %v", fileName, fullPath)

		f.Close()
		manager.Commit()
	}
}

func LoadFromDB(manager Manager) {
	configFullPath, _, _ := manager.GetSaveFilePath()

	CreateSaveFileIfNotExist(manager)
	content, err := ioutil.ReadFile(configFullPath)
	Check(err)

	var obj interface{}
	json.Unmarshal([]byte(content), &obj)
	if obj == nil {
		obj = make(map[string]interface{})
	}

	err = mapstructure.Decode(obj.(map[string]interface{}), manager.GetConfig())
	Check(err)
}

func Commit(manager Manager) error {
	configFullPath, _, configFileName := manager.GetSaveFilePath()
	content, _ := json.Marshal(manager.GetConfig())

	err := ioutil.WriteFile(configFullPath, content, 0644)

	if err != nil {
		log.Printf("Unable to write in configuration file %s (%s) : %s", configFileName, configFullPath, err)
	}

	return err
}
