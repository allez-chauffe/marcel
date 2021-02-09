package config

import (
	"encoding/json"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config viper.Viper

var _default *Config

func SetDefault(cfg *Config) {
	_default = cfg
}

func Default() *Config {
	return _default
}

func New() *Config {
	var cfg = viper.New()

	cfg.SetEnvPrefix("marcel")
	cfg.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	cfg.AutomaticEnv()

	cfg.AddConfigPath("/etc/marcel")
	cfg.AddConfigPath(".")
	cfg.SetConfigName("config")

	c := (*Config)(cfg)

	c.API().SetDefaults()
	c.Backoffice().SetDefaults()
	c.Frontend().SetDefaults()
	c.HTTP().SetDefaults()

	return c
}

func (c *Config) viper() *viper.Viper {
	return (*viper.Viper)(c)
}

func (c *Config) Read(configFile string) error {
	var cfg = c.viper()

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

func (c *Config) Debug() {
	cfgString, err := json.MarshalIndent(c.viper().AllSettings(), "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshall config : %s", err)
	}

	log.Debugf("Current configuration : %s", string(cfgString))
}

func (c *Config) LogLevel() log.Level {
	logLevelString := c.viper().GetString("logLevel")

	// Default to info if no log level is given
	if logLevelString == "" {
		return log.InfoLevel
	}

	l, err := log.ParseLevel(logLevelString)
	if err != nil {
		panic(err)
	}
	return l
}

func (c *Config) SetLogLevel(l log.Level) {
	c.viper().Set("logLevel", l.String())
	log.SetLevel(l)
}

func (c *Config) API() *API {
	return (*API)(c)
}

func (c *Config) Backoffice() *Backoffice {
	return (*Backoffice)(c)
}

func (c *Config) Frontend() *Frontend {
	return (*Frontend)(c)
}

func (c *Config) HTTP() *HTTP {
	return (*HTTP)(c)
}
