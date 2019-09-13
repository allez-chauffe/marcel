package cmd

import (
	"github.com/Zenika/marcel/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func addCommonAPIFlags(conf *viper.Viper, command *cobra.Command) {
	flags := command.Flags()

	flags.UintP("port", "p", config.Config.API.Port, "Listening port")
	if err := conf.BindPFlag("api.port", flags.Lookup("port")); err != nil {
		panic(err)
	}

	flags.String("basePath", config.Config.API.BasePath, "Base path")
	if err := conf.BindPFlag("api.basePath", flags.Lookup("basePath")); err != nil {
		panic(err)
	}

	flags.Bool("cors", config.Config.API.CORS, "Enable CORS (all origins)")
	if err := conf.BindPFlag("api.cors", flags.Lookup("cors")); err != nil {
		panic(err)
	}

	flags.String("dbFile", config.Config.API.DBFile(), "Database file")
	if err := conf.BindPFlag("api.dbFile", flags.Lookup("dbFile")); err != nil {
		panic(err)
	}

	flags.String("pluginsDir", config.Config.API.PluginsDir(), "Plugins directory")
	if err := conf.BindPFlag("api.pluginsDir", flags.Lookup("pluginsDir")); err != nil {
		panic(err)
	}

	flags.String("mediasDir", config.Config.API.MediasDir(), "Medias directory")
	if err := conf.BindPFlag("api.mediasDir", flags.Lookup("mediasDir")); err != nil {
		panic(err)
	}

	flags.String("dataDir", config.Config.API.DataDir(), "Data directory")
	if err := conf.BindPFlag("api.dataDir", flags.Lookup("dataDir")); err != nil {
		panic(err)
	}

	flags.Bool("secure", config.Config.API.Auth.Secure, "Enable secure cookies")
	if err := conf.BindPFlag("api.auth.secure", flags.Lookup("secure")); err != nil {
		panic(err)
	}

	flags.Duration("authExpiration", config.Config.API.Auth.Expiration, "Authentication token expiration")
	if err := conf.BindPFlag("api.auth.expiration", flags.Lookup("authExpiration")); err != nil {
		panic(err)
	}

	flags.Duration("refreshExpiration", config.Config.API.Auth.RefreshExpiration, "Refresh token expiration")
	if err := conf.BindPFlag("api.auth.refreshExpiration", flags.Lookup("refreshExpiration")); err != nil {
		panic(err)
	}
}
