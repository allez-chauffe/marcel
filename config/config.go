package config

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type privateConfig struct {
	LogLevel   log.Level
	API        privateAPI
	Backoffice privateBackoffice
	Frontend   privateFrontend
	Standalone privateStandalone
}

var defaultConfig = privateConfig{
	LogLevel: log.InfoLevel,
	API: privateAPI{
		Port:       8090,
		BasePath:   "/api",
		CORS:       false,
		DBFile:     "marcel.db",
		PluginsDir: "plugins",
		MediasDir:  "medias",
		DataDir:    "",
		Auth: privateAuth{
			Secure:            true,
			Expiration:        8 * time.Hour,
			RefreshExpiration: 15 * 24 * time.Hour,
		},
	},
	Backoffice: privateBackoffice{
		Port:        8090,
		BasePath:    "/",
		APIURI:      "/api",
		FrontendURI: "/front",
	},

	Frontend: privateFrontend{
		Port:     8090,
		BasePath: "/front",
		APIURI:   "/api",
	},

	Standalone: privateStandalone{
		Port: 8090,
	},
}

type publicConfig struct {
	*privateConfig
	API        api
	Backoffice backoffice
	Frontend   frontend
	Standalone standalone
}

// Config is the root configuration object
var Config = publicConfig{
	&defaultConfig,
	api{&defaultConfig.API, auth{&defaultConfig.API.Auth}},
	backoffice{&defaultConfig.Backoffice},
	frontend{&defaultConfig.Frontend},
	standalone{&defaultConfig.Standalone},
}
