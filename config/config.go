package config

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ConfigType struct {
	v *viper.Viper
}

var cfg ConfigType

func SetConfig(newCfg ConfigType) {
	cfg = newCfg
}

func Config() ConfigType {
	return cfg
}

func New() ConfigType {
	var cfg = viper.New()

	cfg.SetEnvPrefix("marcel")
	cfg.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	cfg.AutomaticEnv()

	cfg.AddConfigPath("/etc/marcel")
	cfg.AddConfigPath(".")
	cfg.SetConfigName("config")

	return ConfigType{cfg}
}

func (c ConfigType) Read(configFile string) error {
	if configFile != "" {
		c.v.SetConfigFile(configFile)
	}

	if err := c.v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if configFile != "" {
				log.Fatalf("Config file %s not found", configFile)
			}
			log.Warnln(err)
		} else {
			log.Fatalln("Error loading config file:", err)
		}
	} else {
		log.Infof("Using config file %s", c.v.ConfigFileUsed())
	}

	return nil
}

func (c ConfigType) Debug() {
	log.Debug("FIXME")
}

// func (c ConfigType) LogLevel() log.Level {
// 	l, err := log.ParseLevel(c.v.GetString("logLevel"))
// 	if err != nil {
// 		panic(err)
// 	}
// 	return l
// }

// func (c ConfigType) SetLogLevel(l log.Level) {
// 	c.v.Set("logLevel", l.String())
// }

func (c ConfigType) API() API {
	return API{c.v}
}

func (c ConfigType) Backoffice() Backoffice {
	return Backoffice{c.v}
}

func (c ConfigType) Frontend() Frontend {
	return Frontend{c.v}
}

func (c ConfigType) Standalone() Standalone {
	return Standalone{c.v}
}

// var defaultConfig = _config{
// 	LogLevel: log.InfoLevel,
// 	API: _API{
// 		Port:       8090,
// 		BasePath:   "/api",
// 		CORS:       false,
// 		DBFile:     "marcel.db",
// 		PluginsDir: "plugins",
// 		MediasDir:  "medias",
// 		DataDir:    "",
// 		Auth: _Auth{
// 			Secure:            true,
// 			Expiration:        8 * time.Hour,
// 			RefreshExpiration: 15 * 24 * time.Hour,
// 		},
// 	},
// 	Backoffice: _Backoffice{
// 		Port:        8090,
// 		BasePath:    "/",
// 		APIURI:      "/api",
// 		FrontendURI: "/front",
// 	},

// 	Frontend: _Frontend{
// 		Port:     8090,
// 		BasePath: "/front",
// 		APIURI:   "/api",
// 	},

// 	Standalone: _Standalone{
// 		Port: 8090,
// 	},
// }
