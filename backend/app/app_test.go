package app

import (
	"log"
	"os"
	"testing"
)

var logFileTest *os.File
var logPathTest = "testMarcelLog"

func TestApp_InitializeLog(t *testing.T) {
	var a = new(App)

	err := a.InitializeLog(logPathTest, logFileTest)
	defer closeLogFile()

	if err != nil {
		t.Error(err)
	}

	log.Println("test output")
}

func closeLogFile() {
	logFile.Close()
	os.Remove(logPathTest)
}
