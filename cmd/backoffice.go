package cmd

import (
	"os"

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

		PreRunE: preRunForServer(cfg),

		Run: func(_ *cobra.Command, _ []string) {
			os.Exit(backoffice.Module().Run())
		},
	}

	var flags = cmd.Flags()

	if _, err := cfg.FlagUintP(flags, "port", "p", cfg.Backoffice().Port(), "Listening port", "backoffice.port"); err != nil {
		panic(err)
	}

	if _, err := cfg.FlagString(flags, "basePath", cfg.Backoffice().BasePath(), "Base path", "backoffice.basePath"); err != nil {
		panic(err)
	}

	if _, err := cfg.FlagString(flags, "apiURI", cfg.Backoffice().APIURI(), "API URI", "backoffice.apiURI"); err != nil {
		panic(err)
	}

	if _, err := cfg.FlagString(flags, "frontendURI", cfg.Backoffice().FrontendURI(), "Frontend URI", "backoffice.frontendURI"); err != nil {
		panic(err)
	}

	Marcel.AddCommand(cmd)
}
