package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/allez-chauffe/marcel/frontend"
	"github.com/allez-chauffe/marcel/pkg/config"
)

func init() {
	var cfg = config.New()

	var cmd = &cobra.Command{
		Use:   "frontend",
		Short: "Starts Marcel's Frontend server",
		Args:  cobra.NoArgs,

		PreRunE: preRunForServer(cfg),

		Run: func(_ *cobra.Command, _ []string) {
			os.Exit(frontend.Module().Run())
		},
	}

	var flags = cmd.Flags()

	if _, err := cfg.FlagUintP(flags, "port", "p", cfg.HTTP().Port(), "Listening port", "http.port"); err != nil {
		panic(err)
	}

	if _, err := cfg.FlagString(flags, "basePath", cfg.Frontend().BasePath(), "Base path", "frontend.basePath"); err != nil {
		panic(err)
	}

	if _, err := cfg.FlagString(flags, "apiURI", cfg.API().BasePath(), "API URI", "api.basePath"); err != nil {
		panic(err)
	}

	Marcel.AddCommand(cmd)
}
