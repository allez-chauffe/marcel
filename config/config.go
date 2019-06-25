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
		BasePath:   "/api/v1",
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
		APIURI:      "/api/v1",
		FrontendURI: "/front",
	},
	Frontend: frontend{
		Port:     8090,
		BasePath: "/front",
		APIURI:   "/api/v1",
	},
	Standalone: standalone{
		Port: 8090,
	},
}
