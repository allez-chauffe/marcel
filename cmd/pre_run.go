package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/Zenika/marcel/config"
)

func preRunForServer(cfg *config.ConfigType) func(*cobra.Command, []string) {
	return func(_ *cobra.Command, _ []string) {
		log.SetOutput(os.Stdout)
		bindLogLevel(cfg)
		cfg.Read(configFile)
		config.SetConfig(cfg)
		setLogLevel(cfg)
		cfg.Debug()
	}
}
