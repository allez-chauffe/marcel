package config

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ConfigType viper.Viper

var cfg *ConfigType

func SetConfig(newCfg *ConfigType) {
	cfg = newCfg
}

func Config() *ConfigType {
	return cfg
}

func New() *ConfigType {
	var cfg = viper.New()

	cfg.SetEnvPrefix("marcel")
	cfg.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	cfg.AutomaticEnv()

	cfg.AddConfigPath("/etc/marcel")
	cfg.AddConfigPath(".")
	cfg.SetConfigName("config")

	return (*ConfigType)(cfg)
}

func (c *ConfigType) Read(configFile string) error {
	var cfg = (*viper.Viper)(c)

	if configFile != "" {
		cfg.SetConfigFile(configFile)
	}

	if err := cfg.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if configFile != "" {
				log.Fatalf("Config file %s not found", configFile)
			}
			log.Warnln(err)
		} else {
			log.Fatalln("Error loading config file:", err)
		}
	} else {
		log.Infof("Using config file %s", cfg.ConfigFileUsed())
	}

	return nil
}

func (c *ConfigType) Debug() {
	log.Debug("FIXME")
}

// func (c *ConfigType) LogLevel() log.Level {
// 	l, err := log.ParseLevel(cfg.GetString("logLevel"))
// 	if err != nil {
// 		panic(err)
// 	}
// 	return l
// }

// func (c *ConfigType) SetLogLevel(l log.Level) {
// 	cfg.Set("logLevel", l.String())
// }

func (c *ConfigType) API() *API {
	return (*API)(c)
}

func (c *ConfigType) Backoffice() *Backoffice {
	return (*Backoffice)(c)
}

func (c *ConfigType) Frontend() *Frontend {
	return (*Frontend)(c)
}

func (c *ConfigType) Standalone() *Standalone {
	return (*Standalone)(c)
}
