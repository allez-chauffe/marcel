package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Zenika/marcel/api"
	"github.com/Zenika/marcel/config"
)

func init() {
	var apiConfig = viper.New()

	var apiCmd = &cobra.Command{
		Use:   "api",
		Short: "Starts marcel's api server",
		Args:  cobra.NoArgs,

		PreRun: func(cmd *cobra.Command, args []string) {
			log.SetOutput(os.Stdout)
			config.Init(apiConfig, configFile)
			setLogLevel()
			config.Debug()
		},

		Run: func(cmd *cobra.Command, args []string) {
			a := api.New()
			a.Initialize()
			a.Start()
		},
	}

	addCommonAPIFlags(apiConfig, apiCmd)

	Marcel.AddCommand(apiCmd)
}
