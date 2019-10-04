package cmd

import (
	"github.com/spf13/cobra"

	"github.com/Zenika/marcel/config"
	"github.com/Zenika/marcel/standalone"
)

func init() {
	var cfg = config.New()

	var cmd = &cobra.Command{
		Use:   "standalone",
		Short: "Starts marcel's standalone server",
		Args:  cobra.NoArgs,

		PreRun: preRunForServer(cfg),

		RunE: func(_ *cobra.Command, _ []string) error {
			return standalone.Start()
		},
	}

	var flags = cmd.Flags()

	commonAPIFlags(flags, cfg)

	if err := cfg.FlagUintP(flags, "port", "p", 8090, "Listening port", "standalone.port"); err != nil {
		panic(err)
	}

	if err := cfg.FlagString(flags, "apiBasePath", "/api", "Base path", "api.basePath", "backoffice.apiURI", "frontend.apiURI"); err != nil {
		panic(err)
	}

	if err := cfg.FlagString(flags, "backofficeBasePath", "/", "Backoffice base path", "backoffice.basePath"); err != nil {
		panic(err)
	}

	if err := cfg.FlagString(flags, "frontendBasePath", "/front", "Frontend base path", "frontend.basePath", "backoffice.frontendURI"); err != nil {
		panic(err)
	}

	Marcel.AddCommand(cmd)
}
