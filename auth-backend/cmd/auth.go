package cmd

import (
	"github.com/spf13/cobra"

	"github.com/Zenika/MARCEL/auth-backend/app"
	"github.com/Zenika/MARCEL/auth-backend/conf"
	"github.com/Zenika/MARCEL/auth-backend/users"
)

var (
	configPath string
)

func init() {
	AuthCmd.Flags().StringVarP(&configPath, "config", "c", "config/config.json", "Config file")
	AuthCmd.Flags().StringVar(&users.UsersFilePath, "users-file", "config/users.json", "Users data file")
}

// AuthCmd is the auth server command
var AuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "Starts Marcel's auth server",
	Args:  cobra.NoArgs,

	Run: func(cmd *cobra.Command, args []string) {
		config := conf.LoadConfig(configPath)
		users.LoadUsersData()
		app.Run(config)
	},
}
