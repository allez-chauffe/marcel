package app

import (
	"testing"
	"os"
	"log"
)

var logFileTest *os.File
var logPathTest string = "testMarcelLog"

func TestApp_InitializeLog(t *testing.T) {
	t.Log("InitializeLog test")
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