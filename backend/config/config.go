package config

import (
	"time"

	log "github.com/sirupsen/logrus"
)

// Config is the root configuration object
var Config = struct {
	Port        uint
	LogLevel    log.Level
	DataPath    string
	ClientsFile string
	MediasFile  string
	PluginsFile string
	PluginsPath string
	Auth        auth
}{
	Port:        8090,
	LogLevel:    log.InfoLevel,
	DataPath:    "data",
	ClientsFile: "clients.json",
	MediasFile:  "medias.json",
	PluginsFile: "plugins.json",
	PluginsPath: "plugins",
	Auth: auth{
		Port:              8090,
		Secured:           true,
		AuthExpiration:    8 * time.Hour,
		RefreshExpiration: 15 * 24 * time.Hour,
		UsersFile:         "data/users.json",
	},
}
