package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Zenika/MARCEL/backend/app"
	"github.com/Zenika/MARCEL/backend/config"
)

var (
	port uint
)

func init() {
	Cmd.Flags().UintVarP(&port, "port", "p", 8090, "Listening port")

	Cmd.Flags().StringVar(&config.Global.ConfigPath, "config-path", "data", "Directory containing config files")
	Cmd.Flags().StringVar(&config.Global.ClientsConfigFile, "clients-config-file", "clients.json", "Clients config file name")
	Cmd.Flags().StringVar(&config.Global.MediasConfigFile, "medias-config-file", "medias.json", "Medias config file name")
	Cmd.Flags().StringVar(&config.Global.PluginsConfigFile, "plugins-config-file", "plugins.json", "Plugins config file name")

	Cmd.Flags().StringVar(&config.Plugins.Path, "plugins-path", "plugins", "Plugins directory")
}

// Cmd is the root command of Marcel
var Cmd = &cobra.Command{
	Use:   "marcel",
	Short: "Marcel is a configurable plugin based dashboard system",
	Args:  cobra.NoArgs,

	Run: func(cmd *cobra.Command, args []string) {
		a := new(app.App)
		a.Initialize()

		a.Run(fmt.Sprintf(":%d", port))
	},
}
