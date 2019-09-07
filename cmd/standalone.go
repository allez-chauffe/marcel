package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Zenika/marcel/config"
	"github.com/Zenika/marcel/standalone"
)

func init() {
	var standaloneConfig = viper.New()

	var standaloneCmd = &cobra.Command{
		Use:   "standalone",
		Short: "Starts marcel's standalone server",
		Args:  cobra.NoArgs,

		PreRun: func(cmd *cobra.Command, args []string) {
			log.SetOutput(os.Stdout)
			config.Init(standaloneConfig, configFile)
			setLogLevel()
			config.Debug()
		},

		Run: func(cmd *cobra.Command, args []string) {
			standalone.Start()
		},
	}

	var flags = standaloneCmd.Flags()

	flags.UintP("port", "p", config.Config.Standalone.Port, "Listening port")
	if err := standaloneConfig.BindPFlag("standalone.port", flags.Lookup("port")); err != nil {
		panic(err)
	}

	flags.String("apiBasePath", config.Config.API.BasePath, "API base path")
	if err := standaloneConfig.BindPFlag("api.basePath", flags.Lookup("apiBasePath")); err != nil {
		panic(err)
	}
	if err := standaloneConfig.BindPFlag("backoffice.apiURI", flags.Lookup("apiBasePath")); err != nil {
		panic(err)
	}
	if err := standaloneConfig.BindPFlag("frontend.apiURI", flags.Lookup("apiBasePath")); err != nil {
		panic(err)
	}

	flags.String("dbFile", config.Config.API.DBFile, "Database file")
	if err := standaloneConfig.BindPFlag("api.dbFile", flags.Lookup("dbFile")); err != nil {
		panic(err)
	}

	flags.String("pluginsDir", config.Config.API.PluginsDir, "Plugins directory")
	if err := standaloneConfig.BindPFlag("api.pluginsDir", flags.Lookup("pluginsDir")); err != nil {
		panic(err)
	}

	flags.String("mediasDir", config.Config.API.MediasDir, "Medias directory")
	if err := standaloneConfig.BindPFlag("api.mediasDir", flags.Lookup("mediasDir")); err != nil {
		panic(err)
	}

	flags.Bool("secure", config.Config.API.Auth.Secure, "Enable secure cookies")
	if err := standaloneConfig.BindPFlag("api.auth.secure", flags.Lookup("secure")); err != nil {
		panic(err)
	}

	flags.Duration("authExpiration", config.Config.API.Auth.Expiration, "Authentication token expiration")
	if err := standaloneConfig.BindPFlag("api.auth.expiration", flags.Lookup("authExpiration")); err != nil {
		panic(err)
	}

	flags.Duration("refreshExpiration", config.Config.API.Auth.RefreshExpiration, "Refresh token expiration")
	if err := standaloneConfig.BindPFlag("api.auth.refreshExpiration", flags.Lookup("refreshExpiration")); err != nil {
		panic(err)
	}

	flags.String("backofficeBasePath", config.Config.Backoffice.BasePath, "Backoffice base path")
	if err := standaloneConfig.BindPFlag("backoffice.basePath", flags.Lookup("backofficeBasePath")); err != nil {
		panic(err)
	}

	flags.String("frontendBasePath", config.Config.Frontend.BasePath, "Frontend base path")
	if err := standaloneConfig.BindPFlag("frontend.basePath", flags.Lookup("frontendBasePath")); err != nil {
		panic(err)
	}
	if err := standaloneConfig.BindPFlag("backoffice.frontendURI", flags.Lookup("frontendBasePath")); err != nil {
		panic(err)
	}

	Marcel.AddCommand(standaloneCmd)
}
