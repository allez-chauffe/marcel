package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Zenika/marcel/api"
	"github.com/Zenika/marcel/config"
)

func init() {
	var apiConfig = viper.New()

	var apiCmd = &cobra.Command{
		Use:   "api",
		Short: "Starts marcel's api server",
		Args:  cobra.NoArgs,

		PreRun: func(cmd *cobra.Command, args []string) {
			log.SetOutput(os.Stdout)
			config.Init(apiConfig, configFile)
			setLogLevel()
			config.Debug()
		},

		Run: func(cmd *cobra.Command, args []string) {
			a := api.New()
			a.Initialize()
			a.Start()
		},
	}

	var flags = apiCmd.Flags()

	// FIXME Use a flagSet to mutualize common flags with standalone

	flags.UintP("port", "p", config.Config.API.Port, "Listening port")
	if err := apiConfig.BindPFlag("api.port", flags.Lookup("port")); err != nil {
		panic(err)
	}

	flags.String("basePath", config.Config.API.BasePath, "Base path")
	if err := apiConfig.BindPFlag("api.basePath", flags.Lookup("basePath")); err != nil {
		panic(err)
	}

	flags.Bool("cors", config.Config.API.CORS, "Enable CORS (all origins)")
	if err := apiConfig.BindPFlag("api.cors", flags.Lookup("cors")); err != nil {
		panic(err)
	}

	flags.String("dbFile", config.Config.API.DBFile, "Database file")
	if err := apiConfig.BindPFlag("api.dbFile", flags.Lookup("dbFile")); err != nil {
		panic(err)
	}

	flags.String("pluginsDir", config.Config.API.PluginsDir, "Plugins directory")
	if err := apiConfig.BindPFlag("api.pluginsDir", flags.Lookup("pluginsDir")); err != nil {
		panic(err)
	}

	flags.String("mediasDir", config.Config.API.MediasDir, "Medias directory")
	if err := apiConfig.BindPFlag("api.mediasDir", flags.Lookup("mediasDir")); err != nil {
		panic(err)
	}

	flags.String("dataDir", config.Config.API.DataDir, "Data directory")
	if err := apiConfig.BindPFlag("api.dataDir", flags.Lookup("dataDir")); err != nil {
		panic(err)
	}

	flags.Bool("secure", config.Config.API.Auth.Secure, "Enable secure cookies")
	if err := apiConfig.BindPFlag("api.auth.secure", flags.Lookup("secure")); err != nil {
		panic(err)
	}

	flags.Duration("authExpiration", config.Config.API.Auth.Expiration, "Authentication token expiration")
	if err := apiConfig.BindPFlag("api.auth.expiration", flags.Lookup("authExpiration")); err != nil {
		panic(err)
	}

	flags.Duration("refreshExpiration", config.Config.API.Auth.RefreshExpiration, "Refresh token expiration")
	if err := apiConfig.BindPFlag("api.auth.refreshExpiration", flags.Lookup("refreshExpiration")); err != nil {
		panic(err)
	}

	Marcel.AddCommand(apiCmd)
}
