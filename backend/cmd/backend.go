package cmd

import (
	"github.com/spf13/cobra"

	"github.com/Zenika/MARCEL/backend/app"
	"github.com/Zenika/MARCEL/backend/config"
)

func init() {
	backend.Flags().UintVarP(&config.Global.Port, "port", "p", 8090, "Listening port")
	backend.Flags().StringVar(&config.Global.DataPath, "data-path", "data", "Directory containing data files")
	backend.Flags().StringVar(&config.Global.ClientsFile, "clients-file", "clients.json", "Clients data file name")
	backend.Flags().StringVar(&config.Global.MediasFile, "medias-file", "medias.json", "Medias data file name")
	backend.Flags().StringVar(&config.Global.PluginsFile, "plugins-file", "plugins.json", "Plugins data file name")

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
