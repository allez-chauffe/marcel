package config

import (
	log "github.com/sirupsen/logrus"
)

// Global configuration
var Global = struct {
	Port        uint
	LogLevel    log.Level
	DataPath    string
	ClientsFile string
	MediasFile  string
	PluginsFile string
}{
	LogLevel: log.InfoLevel,
}
