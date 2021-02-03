package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/allez-chauffe/marcel/config"
	"github.com/allez-chauffe/marcel/standalone"
)

func init() {
	var cfg = config.New()

	var cmd = &cobra.Command{
		Use:   "standalone",
		Short: "Starts marcel's standalone server",
		Args:  cobra.NoArgs,

		PreRunE: preRunForServer(cfg),

		Run: func(_ *cobra.Command, _ []string) {
			os.Exit(standalone.Module().Run())
		},
	}

	var flags = cmd.Flags()

	commonAPIFlags(flags, cfg)

	if _, err := cfg.FlagUintP(flags, "port", "p", cfg.HTTP().Port(), "Listening port", "http.port"); err != nil {
		panic(err)
	}

	if _, err := cfg.FlagString(flags, "basePath", cfg.HTTP().BasePath(), "Base path", "http.basePath"); err != nil {
		panic(err)
	}

	if _, err := cfg.FlagString(flags, "apiBasePath", cfg.API().BasePath(), "API base path", "api.basePath"); err != nil {
		panic(err)
	}

	if _, err := cfg.FlagString(flags, "backofficeBasePath", cfg.Backoffice().BasePath(), "Backoffice base path", "backoffice.basePath"); err != nil {
		panic(err)
	}

	if _, err := cfg.FlagString(flags, "frontendBasePath", cfg.Frontend().BasePath(), "Frontend base path", "frontend.basePath"); err != nil {
		panic(err)
	}

	Marcel.AddCommand(cmd)
}
