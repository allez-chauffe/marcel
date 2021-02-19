package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/allez-chauffe/marcel/pkg/config"
)

func preRunForServer(cfg *config.Config) func(*cobra.Command, []string) error {
	return func(_ *cobra.Command, _ []string) error {
		log.SetOutput(os.Stdout)

		bindLogLevel(cfg)

		if err := cfg.Read(configFile); err != nil {
			return err
		}

		config.SetDefault(cfg)

		setLogLevel(cfg)

		cfg.Debug()

		return nil
	}
}
