package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Zenika/MARCEL/backend/app"
	"github.com/Zenika/MARCEL/backend/config"
)

func init() {
	backend.Flags().UintP("port", "p", config.Config.Port, "Listening port")
	viper.BindPFlag("port", backend.Flags().Lookup("port"))

	backend.Flags().String("dataPath", config.Config.DataPath, "Data files directory")
	viper.BindPFlag("dataPath", backend.Flags().Lookup("dataPath"))

	backend.Flags().String("clientsFile", config.Config.ClientsFile, "Clients data file name")
	viper.BindPFlag("clientsFile", backend.Flags().Lookup("clientsFile"))

	backend.Flags().String("mediasFile", config.Config.MediasFile, "Medias data file name")
	viper.BindPFlag("mediasFile", backend.Flags().Lookup("mediasFile"))

	backend.Flags().String("pluginsFile", config.Config.PluginsFile, "Plugins data file name")
	viper.BindPFlag("pluginsFile", backend.Flags().Lookup("pluginsFile"))

	backend.Flags().String("usersFile", config.Config.UsersFile, "Users data file")
	viper.BindPFlag("usersFile", backend.Flags().Lookup("usersFile"))

	backend.Flags().String("pluginsPath", config.Config.PluginsPath, "Plugins directory")
	viper.BindPFlag("pluginsPath", backend.Flags().Lookup("pluginsPath"))

	backend.Flags().Bool("secured", config.Config.Auth.Secured, "Use secured cookies")
	viper.BindPFlag("auth.secured", backend.Flags().Lookup("secured"))

	backend.Flags().Duration("authExpiration", config.Config.Auth.AuthExpiration, "Authentication token expiration")
	viper.BindPFlag("auth.authExpiration", backend.Flags().Lookup("authExpiration"))

	backend.Flags().Duration("refreshExpiration", config.Config.Auth.RefreshExpiration, "Refresh token expiration")
	viper.BindPFlag("auth.refreshExpiration", backend.Flags().Lookup("refreshExpiration"))

	backend.Flags().String("domain", config.Config.Auth.Domain, "Cookies domain")
	viper.BindPFlag("auth.domain", backend.Flags().Lookup("domain"))

	backend.Flags().String("baseURL", config.Config.Auth.BaseURL, "Cookies base URL")
	viper.BindPFlag("auth.baseURL", backend.Flags().Lookup("baseURL"))

	Marcel.AddCommand(backend)
}

var backend = &cobra.Command{
	Use:   "backend",
	Short: "Starts Marcel's backend server",
	Args:  cobra.NoArgs,

	Run: func(cmd *cobra.Command, args []string) {
		a := new(app.App)
		a.Initialize()
		a.Run()
	},
}
