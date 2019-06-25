package config

import (
	log "github.com/sirupsen/logrus"
)

func Debug() {
	log.Debugf("Config: %+v", Config)
}
