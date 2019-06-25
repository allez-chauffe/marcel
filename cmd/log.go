package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/Zenika/marcel/config"
)

type logLevel log.Level

func (l *logLevel) String() string {
	return log.Level(*l).String()
}

func (l *logLevel) Set(s string) error {
	v, err := log.ParseLevel(s)
	if err != nil {
		return err
	}
	*l = logLevel(v)
	return nil
}

func (l *logLevel) Type() string {
	return "log.Level"
}

func setLogLevel() {
	log.SetLevel(config.Config.LogLevel)
	log.Infof("Log level set to %s", config.Config.LogLevel)
}
