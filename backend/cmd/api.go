package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Zenika/MARCEL/backend/app"
	"github.com/Zenika/MARCEL/backend/config"
)

func init() {
	api.Flags().UintP("port", "p", config.Config.Port, "Listening port")
	viper.BindPFlag("port", api.Flags().Lookup("port"))

	api.Flags().String("dataPath", config.Config.DataPath, "Data files directory")
	viper.BindPFlag("dataPath", api.Flags().Lookup("dataPath"))

	api.Flags().String("clientsFile", config.Config.ClientsFile, "Clients data file name")
	viper.BindPFlag("clientsFile", api.Flags().Lookup("clientsFile"))

	api.Flags().String("mediasFile", config.Config.MediasFile, "Medias data file name")
	viper.BindPFlag("mediasFile", api.Flags().Lookup("mediasFile"))

	api.Flags().String("pluginsFile", config.Config.PluginsFile, "Plugins data file name")
	viper.BindPFlag("pluginsFile", api.Flags().Lookup("pluginsFile"))

	api.Flags().String("usersFile", config.Config.UsersFile, "Users data file")
	viper.BindPFlag("usersFile", api.Flags().Lookup("usersFile"))

	api.Flags().String("pluginsPath", config.Config.PluginsPath, "Plugins directory")
	viper.BindPFlag("pluginsPath", api.Flags().Lookup("pluginsPath"))

	api.Flags().Bool("secured", config.Config.Auth.Secured, "Use secured cookies")
	viper.BindPFlag("auth.secured", api.Flags().Lookup("secured"))

	api.Flags().Duration("authExpiration", config.Config.Auth.AuthExpiration, "Authentication token expiration")
	viper.BindPFlag("auth.authExpiration", api.Flags().Lookup("authExpiration"))

	api.Flags().Duration("refreshExpiration", config.Config.Auth.RefreshExpiration, "Refresh token expiration")
	viper.BindPFlag("auth.refreshExpiration", api.Flags().Lookup("refreshExpiration"))

	api.Flags().String("domain", config.Config.Auth.Domain, "Cookies domain")
	viper.BindPFlag("auth.domain", api.Flags().Lookup("domain"))

	Marcel.AddCommand(api)
}

var api = &cobra.Command{
	Use:   "api",
	Short: "Starts Marcel's api server",
	Args:  cobra.NoArgs,

	Run: func(cmd *cobra.Command, args []string) {
		a := new(app.App)
		a.Initialize()
		a.Run()
	},
}
