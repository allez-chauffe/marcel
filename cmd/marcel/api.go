package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/allez-chauffe/marcel/pkg/api"
	"github.com/allez-chauffe/marcel/pkg/config"
)

func init() {
	var cfg = config.New()

	var cmd = &cobra.Command{
		Use:   "api",
		Short: "Starts Marcel's API server",
		Args:  cobra.NoArgs,

		PreRunE: preRunForServer(cfg),

		Run: func(_ *cobra.Command, _ []string) {
			os.Exit(api.Module().Run())
		},
	}

	var flags = cmd.Flags()

	commonAPIFlags(flags, cfg)

	if _, err := cfg.FlagUintP(flags, "port", "p", cfg.HTTP().Port(), "Listening port", "http.port"); err != nil {
		panic(err)
	}

	if _, err := cfg.FlagString(flags, "basePath", cfg.API().BasePath(), "Base path", "api.basePath"); err != nil {
		panic(err)
	}

	if _, err := cfg.FlagBool(flags, "cors", cfg.HTTP().CORS(), "Enable CORS (all origins)", "http.cors"); err != nil {
		panic(err)
	}

	rootCmd.AddCommand(cmd)
}

func commonAPIFlags(flags *pflag.FlagSet, cfg *config.Config) {
	if _, err := cfg.FlagString(flags, "dbFile", cfg.API().DB().Bolt().File(), "Database file", "api.db.bolt.file"); err != nil {
		panic(err)
	}

	if _, err := cfg.FlagString(flags, "pluginsDir", cfg.API().PluginsDir(), "Plugins directory", "api.pluginsDir"); err != nil {
		panic(err)
	}

	if _, err := cfg.FlagString(flags, "mediasDir", cfg.API().MediasDir(), "Medias directory", "api.mediasDir"); err != nil {
		panic(err)
	}

	if _, err := cfg.FlagString(flags, "dataDir", cfg.API().DataDir(), "Data directory", "api.dataDir"); err != nil {
		panic(err)
	}

	if _, err := cfg.FlagBool(flags, "secure", cfg.API().Auth().Secure(), "Enable secure cookies", "api.auth.secure"); err != nil {
		panic(err)
	}

	if _, err := cfg.FlagDuration(flags, "authExpiration", cfg.API().Auth().Expiration(), "Authentication token expiration", "api.auth.expiration"); err != nil {
		panic(err)
	}

	if _, err := cfg.FlagDuration(flags, "refreshExpiration", cfg.API().Auth().RefreshExpiration(), "Refresh token expiration", "api.auth.refreshExpiration"); err != nil {
		panic(err)
	}
}
