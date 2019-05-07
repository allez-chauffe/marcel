package plugins

import (
	"path/filepath"
	"testing"
)

func TestNewManager(t *testing.T) {
	m := NewManager(PLUGINS_CONFIG_PATH, PLUGINS_CONFIG_FILENAME)

	if m.Config == nil {
		t.Error("Configuration should not be nil")
	}

	configFullpath := filepath.Join(PLUGINS_CONFIG_PATH, PLUGINS_CONFIG_FILENAME)
	if m.ConfigFullpath != configFullpath {
		t.Error("Full path for configuration is not correct")
	}
}
