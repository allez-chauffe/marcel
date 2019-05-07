package cmd

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Zenika/MARCEL/config"
)

var configFile string

func init() {
	Marcel.PersistentFlags().Var((*LogLevel)(&config.Config.LogLevel), "logLevel", fmt.Sprintf("Log level: %s, %s, %s, %s or %s", log.TraceLevel, log.DebugLevel, log.InfoLevel, log.WarnLevel, log.ErrorLevel))
	viper.BindPFlag("logLevel", Marcel.PersistentFlags().Lookup("logLevel"))

	Marcel.PersistentFlags().StringVarP(&configFile, "configFile", "c", "", fmt.Sprintf("Config file (default /etc/marcel/config.xxx or ./config.xxx, supports %s)", strings.Join(viper.SupportedExts, " ")))
}

// Marcel is the root command of Marcel
var Marcel = &cobra.Command{
	Use:           "marcel",
	Short:         "Marcel is a configurable plugin based dashboard system",
	Args:          cobra.NoArgs,
	SilenceErrors: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if cmd.Name() == "help" {
			return
		}
		config.Init(configFile)
		setLogLevel()
		debugConfig()
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

func setLogLevel() {
	log.SetLevel(config.Config.LogLevel)
	log.Infof("Log level set to %s", config.Config.LogLevel)
}

func debugConfig() {
	log.Debugf("Config: %+v", config.Config)
}
