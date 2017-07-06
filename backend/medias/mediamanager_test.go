package medias

import (
	/*"testing"
	"os"
	"fmt"*/
)

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