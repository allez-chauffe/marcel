package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Zenika/MARCEL/api"
	"github.com/Zenika/MARCEL/config"
)

func init() {
	apiCmd.Flags().UintP("port", "p", config.Config.Port, "Listening port")
	viper.BindPFlag("port", apiCmd.Flags().Lookup("port"))

	apiCmd.Flags().String("dataPath", config.Config.DataPath, "Data files directory")
	viper.BindPFlag("dataPath", apiCmd.Flags().Lookup("dataPath"))

	apiCmd.Flags().String("dbFile", config.Config.DBFile, "Database file name")
	viper.BindPFlag("dbFile", apiCmd.Flags().Lookup("dbFile"))

	apiCmd.Flags().String("mediasFile", config.Config.MediasFile, "Medias data file name")
	viper.BindPFlag("mediasFile", apiCmd.Flags().Lookup("mediasFile"))

	apiCmd.Flags().String("pluginsPath", config.Config.PluginsPath, "Plugins directory")
	viper.BindPFlag("pluginsPath", apiCmd.Flags().Lookup("pluginsPath"))

	apiCmd.Flags().Bool("secure", config.Config.Auth.Secure, "Use secured cookies")
	viper.BindPFlag("auth.secure", apiCmd.Flags().Lookup("secure"))

	apiCmd.Flags().Duration("authExpiration", config.Config.Auth.AuthExpiration, "Authentication token expiration")
	viper.BindPFlag("auth.authExpiration", apiCmd.Flags().Lookup("authExpiration"))

	apiCmd.Flags().Duration("refreshExpiration", config.Config.Auth.RefreshExpiration, "Refresh token expiration")
	viper.BindPFlag("auth.refreshExpiration", apiCmd.Flags().Lookup("refreshExpiration"))

	Marcel.AddCommand(apiCmd)
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Starts Marcel's api server",
	Args:  cobra.NoArgs,

	Run: func(cmd *cobra.Command, args []string) {
		a := new(api.App)
		a.Initialize()
		a.Run()
	},
}
