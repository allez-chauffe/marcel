package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/Zenika/MARCEL/backend/config"
)

func init() {
	Marcel.PersistentFlags().Var((*LogLevel)(&config.Global.LogLevel), "log-level", fmt.Sprintf("Log level: %s, %s, %s, %s or %s", log.TraceLevel, log.DebugLevel, log.InfoLevel, log.WarnLevel, log.ErrorLevel))
}

// Marcel is the root command of Marcel
var Marcel = &cobra.Command{
	Use:   "marcel",
	Short: "Marcel is a configurable plugin based dashboard system",
	Args:  cobra.NoArgs,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log.SetLevel(config.Global.LogLevel)
	},
}

// LogLevel implements a pflag.Value with logrus.Level
type LogLevel log.Level

func (l *LogLevel) String() string {
	return log.Level(*l).String()
}

func (l *LogLevel) Set(s string) error {
	v, err := log.ParseLevel(s)
	if err != nil {
		return err
	}
	*l = LogLevel(v)
	return nil
}

func (l *LogLevel) Type() string {
	return "log.Level"
}
