package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/allez-chauffe/marcel/pkg/backoffice"
	"github.com/allez-chauffe/marcel/pkg/config"
)

func init() {
	var cfg = config.New()

	var cmd = &cobra.Command{
		Use:   "backoffice",
		Short: "Starts Marcel's Backoffice server",
		Args:  cobra.NoArgs,

		PreRunE: preRunForServer(cfg),

		Run: func(_ *cobra.Command, _ []string) {
			os.Exit(backoffice.Module().Run())
		},
	}

	var flags = cmd.Flags()

	if _, err := cfg.FlagUintP(flags, "port", "p", cfg.HTTP().Port(), "Listening port", "http.port"); err != nil {
		panic(err)
	}

	if _, err := cfg.FlagString(flags, "basePath", cfg.Backoffice().BasePath(), "Base path", "backoffice.basePath"); err != nil {
		panic(err)
	}

	if _, err := cfg.FlagString(flags, "apiURI", cfg.API().BasePath(), "API URI", "api.basePath"); err != nil {
		panic(err)
	}

	if _, err := cfg.FlagString(flags, "frontendURI", cfg.Frontend().BasePath(), "Frontend URI", "frontend.basePath"); err != nil {
		panic(err)
	}

	Marcel.AddCommand(cmd)
}
