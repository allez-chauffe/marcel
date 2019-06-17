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
}
