package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Port           int  `json:"port"`
	SecuredCookies bool `json:"securedCookies"`
}

const configPath = "config/config.json"

func LoadConfig() *Config {

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return checkError(err, "WARNING : No config file detected, loading default configuration")
	}

	data, err := ioutil.ReadFile(configPath)

	if config := checkError(err, "ERROR: Unreadable config file"); config != nil {
		return config
	}

	config := &Config{}
	err = json.Unmarshal(data, config)

	if defaultConfig := checkError(err, "ERROR: Malformed JSON in config file"); defaultConfig != nil {
		return defaultConfig
	}

	return config
}

func checkError(err error, message string) *Config {
	if err != nil {
		log.Printf(message+" (%s)", err.Error())
		return &Config{
			Port:           8091,
			SecuredCookies: true,
		}
	}

	return nil
}
