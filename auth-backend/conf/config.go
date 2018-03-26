package conf

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Port              int    `json:"port"`
	SecuredCookies    bool   `json:"securedCookies"`
	AuthExpiration    int64  `json:"authExpiration"`
	RefreshExpiration int64  `json:"refreshExpiration"`
	Domain            string `json:"domain"`
	BaseURL           string `json:"baseURL"`
	CookiesNamePrefix string `json:"cookiesNamePrefix"`
}

const configPath = "config/config.json"

func LoadConfig() *Config {

	f, err := os.Open(configPath)

	if os.IsNotExist(err) {
		return checkError(err, "WARNING : No config file detected, loading default configuration")
	}

	if defaultConfig := checkError(err, ""); defaultConfig != nil {
		return defaultConfig
	}

	config := &Config{}
	err = json.NewDecoder(f).Decode(config)

	if defaultConfig := checkError(err, "ERROR: Malformed JSON in config file"); defaultConfig != nil {
		return defaultConfig
	}

	return config
}

func checkError(err error, message string) *Config {
	if err != nil {
		log.Printf(message+" (%s)", err.Error())
		return &Config{
			Port:              8091,
			SecuredCookies:    true,
			AuthExpiration:    3600 * 8,
			RefreshExpiration: 3600 * 24 * 15,
		}
	}

	return nil
}
