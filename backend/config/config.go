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
	UsersFile   string
	PluginsPath string
	Auth        auth
}{
	Port:        8090,
	LogLevel:    log.InfoLevel,
	DataPath:    "data",
	ClientsFile: "clients.json",
	MediasFile:  "medias.json",
	PluginsFile: "plugins.json",
	UsersFile:   "users.json",
	PluginsPath: "plugins",
	Auth: auth{
		Secured:           true,
		AuthExpiration:    8 * time.Hour,
		RefreshExpiration: 15 * 24 * time.Hour,
	},
}
