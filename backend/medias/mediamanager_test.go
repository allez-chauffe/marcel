package medias

import (
	/*"testing"
	"os"
	"fmt"*/
	"testing"
)

func TestNewManager(t *testing.T) {
	t.Log("NewManager test")
	m := NewManager(MEDIAS_CONFIG_PATH, MEDIAS_CONFIG_FILENAME)

	if m.Config == nil {
		t.Error("Configuration should not be nil")
	}

	if m.Config.LastID != 0 {
		t.Error("LastID for a new Media Manager configuration should be 0")
	}
}

/*func TestCreateSaveFileIfNotExist(t *testing.T) {
	dir := "data_test"
	file := "media.config.test"

	//if test folder already exists, delete it
	if _, err := os.Stat(dir); err == nil {
		err = os.Remove(dir)
	}


	var fullPath string = fmt.Sprintf("%s%c%s", dir, os.PathSeparator, file)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		mediaManager := NewMediaManager()
		mediaManager.CreateSaveFileIfNotExist(dir, file)

		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			t.Error(err)
		}

		os.Remove(dir)
	}


	t.Logf("ok")
}
*/