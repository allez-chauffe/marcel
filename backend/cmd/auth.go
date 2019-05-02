package cmd

import (
	"github.com/spf13/cobra"

	"github.com/Zenika/MARCEL/backend/auth/app"
	"github.com/Zenika/MARCEL/backend/auth/conf"
	"github.com/Zenika/MARCEL/backend/auth/users"
)

var (
	configPath string
)

func init() {
	auth.Flags().StringVarP(&configPath, "config-file", "c", "auth/config/config.json", "Auth config file")
	auth.Flags().StringVar(&users.UsersFilePath, "users-file", "auth/config/users.json", "Users data file")

	Marcel.AddCommand(auth)
}

var auth = &cobra.Command{
	Use:   "auth",
	Short: "Starts Marcel's auth server",
	Args:  cobra.NoArgs,

	Run: func(cmd *cobra.Command, args []string) {
		config := conf.LoadConfig(configPath)
		users.LoadUsersData()
		app.Run(config)
	},
}
