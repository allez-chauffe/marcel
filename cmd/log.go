package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/allez-chauffe/marcel/pkg/config"
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

func setLogLevel(cfg *config.Config) {
	log.SetLevel(cfg.LogLevel())
	log.Infof("Log level set to %s", cfg.LogLevel())
}

func bindLogLevel(cfg *config.Config) {
	cfg.BindPFlag(Marcel.PersistentFlags(), "logLevel", "logLevel")
}
