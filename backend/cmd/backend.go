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

	backend.Flags().String("pluginsPath", config.Config.PluginsPath, "Plugins directory")
	viper.BindPFlag("pluginsPath", backend.Flags().Lookup("pluginsPath"))

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
