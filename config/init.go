package config

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Init(config *viper.Viper, configFile string) {
	config.SetEnvPrefix("marcel")
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	bindEnvs(config, Config)

	if configFile == "" {
		config.AddConfigPath("/etc/marcel")
		config.AddConfigPath(".")
		config.SetConfigName("config")
	} else {
		config.SetConfigFile(configFile)
	}

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if configFile != "" {
				log.Fatalf("Config file %s not found", configFile)
			}
			log.Warnln(err)
		} else {
			log.Fatalln("Error loading config file:", err)
		}
	} else {
		log.Infof("Using config file %s", config.ConfigFileUsed())
	}

	if err := config.Unmarshal(&Config, registerLogLevelDecoder); err != nil {
		log.Fatalln("Error in config file:", err)
	}
}
