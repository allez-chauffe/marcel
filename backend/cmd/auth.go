package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Zenika/MARCEL/backend/auth/app"
	"github.com/Zenika/MARCEL/backend/config"
	"github.com/Zenika/MARCEL/backend/users"
)

func init() {
	auth.Flags().UintP("port", "p", config.Config.Auth.Port, "Port")
	viper.BindPFlag("auth.port", auth.Flags().Lookup("port"))

	auth.Flags().Bool("secured", config.Config.Auth.Secured, "Use secured cookies")
	viper.BindPFlag("auth.secured", auth.Flags().Lookup("secured"))

	auth.Flags().Duration("authExpiration", config.Config.Auth.AuthExpiration, "Authentication token expiration")
	viper.BindPFlag("auth.authExpiration", auth.Flags().Lookup("authExpiration"))

	auth.Flags().Duration("refreshExpiration", config.Config.Auth.RefreshExpiration, "Refresh token expiration")
	viper.BindPFlag("auth.refreshExpiration", auth.Flags().Lookup("refreshExpiration"))

	auth.Flags().String("domain", config.Config.Auth.Domain, "Cookies domain")
	viper.BindPFlag("auth.domain", auth.Flags().Lookup("domain"))

	auth.Flags().String("baseURL", config.Config.Auth.BaseURL, "Cookies base URL")
	viper.BindPFlag("auth.baseURL", auth.Flags().Lookup("baseURL"))

	auth.Flags().String("usersFile", config.Config.Auth.UsersFile, "Users data file")
	viper.BindPFlag("auth.usersFile", auth.Flags().Lookup("usersFile"))

	Marcel.AddCommand(auth)
}

var auth = &cobra.Command{
	Use:   "auth",
	Short: "Starts Marcel's auth server",
	Args:  cobra.NoArgs,

	Run: func(cmd *cobra.Command, args []string) {
		users.LoadUsersData()
		app.Run()
	},
}
