package cmd

import (
	"github.com/spf13/cobra"

	"github.com/Zenika/MARCEL/backend/app"
	"github.com/Zenika/MARCEL/backend/config"
)

func init() {
	backend.Flags().UintVarP(&config.Global.Port, "port", "p", 8090, "Listening port")
	backend.Flags().StringVar(&config.Global.ConfigPath, "config-path", "data", "Directory containing config files")
	backend.Flags().StringVar(&config.Global.ClientsConfigFile, "clients-config-file", "clients.json", "Clients config file name")
	backend.Flags().StringVar(&config.Global.MediasConfigFile, "medias-config-file", "medias.json", "Medias config file name")
	backend.Flags().StringVar(&config.Global.PluginsConfigFile, "plugins-config-file", "plugins.json", "Plugins config file name")

	backend.Flags().StringVar(&config.Plugins.Path, "plugins-path", "plugins", "Plugins directory")

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
