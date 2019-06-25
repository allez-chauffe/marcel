package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Zenika/marcel/backoffice"
	"github.com/Zenika/marcel/config"
)

func init() {
	var backofficeConfig = viper.New()

	var backofficeCmd = &cobra.Command{
		Use:   "backoffice",
		Short: "Starts Marcel's backoffice server",
		Args:  cobra.NoArgs,

		PreRun: func(cmd *cobra.Command, args []string) {
			log.SetOutput(os.Stdout)
			config.Init(backofficeConfig, configFile)
			setLogLevel()
			config.Debug()
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			return backoffice.Start()
		},
	}

	var flags = backofficeCmd.Flags()

	flags.Uint("port", config.Config.Backoffice.Port, "Listening port")
	if err := backofficeConfig.BindPFlag("backoffice.port", flags.Lookup("port")); err != nil {
		panic(err)
	}

	flags.String("basePath", config.Config.Backoffice.BasePath, "Base path")
	if err := backofficeConfig.BindPFlag("backoffice.basePath", flags.Lookup("basePath")); err != nil {
		panic(err)
	}

	flags.String("apiURI", config.Config.Backoffice.APIURI, "API URI")
	if err := backofficeConfig.BindPFlag("backoffice.apiURI", flags.Lookup("apiURI")); err != nil {
		panic(err)
	}

	flags.String("frontendURI", config.Config.Backoffice.FrontendURI, "Frontend URI")
	if err := backofficeConfig.BindPFlag("backoffice.frontendURI", flags.Lookup("frontendURI")); err != nil {
		panic(err)
	}

	Marcel.AddCommand(backofficeCmd)
}
