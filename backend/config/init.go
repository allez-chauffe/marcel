package config

import (
	"reflect"

	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Init(configFile string) {
	if configFile == "" {
		viper.AddConfigPath("/etc/marcel")
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	} else {
		viper.SetConfigFile(configFile)
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if configFile != "" {
				log.Fatalf("Config file %s not found", configFile)
			}
			log.Warnln(err)
		} else {
			log.Fatalln("Error loading config file:", err)
		}
	} else {
		log.Infof("Using config file %s", viper.ConfigFileUsed())
	}

	if err := viper.Unmarshal(&Config, registerLogLevelDecoder); err != nil {
		log.Fatalln("Error in config file:", err)
	}
}

func registerLogLevelDecoder(dc *mapstructure.DecoderConfig) {
	dc.DecodeHook = mapstructure.ComposeDecodeHookFunc(
		dc.DecodeHook,
		func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
			if f.Kind() != reflect.String {
				return data, nil
			}
			if t != reflect.TypeOf(log.InfoLevel) {
				return data, nil
			}
			return log.ParseLevel(data.(string))
		},
	)
}
