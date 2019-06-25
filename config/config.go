package config

import (
	"time"

	log "github.com/sirupsen/logrus"
)

// Config is the root configuration object
var Config = struct {
	LogLevel log.Level
	API      api
}{
	LogLevel: log.InfoLevel,
	API: api{
		Port:        8090,
		CORS:        false,
		DBFile:      "marcel.db",
		PluginsPath: "plugins",
		Auth: auth{
			Secure:            true,
			AuthExpiration:    8 * time.Hour,
			RefreshExpiration: 15 * 24 * time.Hour,
		},
	},
}
