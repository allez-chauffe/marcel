package config

import (
	"time"

	log "github.com/sirupsen/logrus"
)

// Config is the root configuration object
var Config = struct {
	LogLevel   log.Level
	API        api
	Backoffice backoffice
	Frontend   frontend
	Standalone standalone
}{
	LogLevel: log.InfoLevel,
	API: api{
		Port:       8090,
		BasePath:   "/api",
		CORS:       false,
		DBFile:     "marcel.db",
		PluginsDir: "plugins",
		Auth: auth{
			Secure:            true,
			Expiration:        8 * time.Hour,
			RefreshExpiration: 15 * 24 * time.Hour,
		},
	},
	Backoffice: backoffice{
		Port:        8090,
		BasePath:    "/",
		APIURI:      "/api",
		FrontendURI: "/front",
	},
	Frontend: frontend{
		Port:     8090,
		BasePath: "/front",
		APIURI:   "/api",
	},
	Standalone: standalone{
		Port: 8090,
	},
}
