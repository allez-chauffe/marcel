package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Zenika/marcel/api"
	"github.com/Zenika/marcel/backoffice"
	"github.com/Zenika/marcel/config"
	"github.com/Zenika/marcel/frontend"
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
			a := api.New()
			a.Initialize()

			r := mux.NewRouter()

			// FIXME routers should be added by basePath order !!!

			a.ConfigureRouter(r)

			frontend.ConfigureRouter(r)

			backoffice.ConfigureRouter(r)

			log.Infof("Starting standalone server on port %d...", config.Config.Standalone.Port)

			log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Standalone.Port), r))
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

	flags.String("dbFile", config.Config.API.DBFile, "Database file")
	if err := standaloneConfig.BindPFlag("api.dbFile", flags.Lookup("dbFile")); err != nil {
		panic(err)
	}

	flags.String("pluginsDir", config.Config.API.PluginsDir, "Plugins directory")
	if err := standaloneConfig.BindPFlag("api.pluginsDir", flags.Lookup("pluginsDir")); err != nil {
		panic(err)
	}

	// FIXME add flag and config for mediasDir...

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
