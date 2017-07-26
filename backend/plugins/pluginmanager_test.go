package plugins

import (
	"testing"
	"fmt"
	"os"
)

func TestNewManager(t *testing.T) {
	t.Log("NewManager test")
	m := NewManager(PLUGINS_CONFIG_PATH, PLUGINS_CONFIG_FILENAME)

	if m.Config == nil {
		t.Error("Configuration should not be nil")
	}

	configFullpath := fmt.Sprintf("%s%c%s", PLUGINS_CONFIG_PATH, os.PathSeparator, PLUGINS_CONFIG_FILENAME)
	if m.ConfigFullpath != configFullpath {
		t.Error("Full path for configuration is not correct")
	}
}
