package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/Zenika/marcel/backoffice"
	"github.com/Zenika/marcel/config"
)

func init() {
	var cfg = config.New()

	var cmd = &cobra.Command{
		Use:   "backoffice",
		Short: "Starts Marcel's backoffice server",
		Args:  cobra.NoArgs,

		PreRun: func(_ *cobra.Command, _ []string) {
			log.SetOutput(os.Stdout)
			cfg.Read(configFile)
			config.SetConfig(cfg)
			setLogLevel(cfg)
			cfg.Debug()
		},

		RunE: func(_ *cobra.Command, _ []string) error {
			return backoffice.Start()
		},
	}

	var flags = cmd.Flags()

	if err := cfg.FlagUintP(flags, "port", "p", 8090, "Listening port", "backoffice.port"); err != nil {
		panic(err)
	}

	if err := cfg.FlagString(flags, "basePath", "/", "Base path", "backoffice.basePath"); err != nil {
		panic(err)
	}

	if err := cfg.FlagString(flags, "apiURI", "/api", "API URI", "backoffice.apiURI"); err != nil {
		panic(err)
	}

	if err := cfg.FlagString(flags, "frontendURI", "/front", "Frontend URI", "backoffice.frontendURI"); err != nil {
		panic(err)
	}

	Marcel.AddCommand(cmd)
}
