package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Zenika/marcel/config"
	"github.com/Zenika/marcel/standalone"
)

func init() {
	var standaloneConfig = viper.New()

	var standaloneCmd = &cobra.Command{
		Use:   "standalone",
		Short: "Starts marcel's standalone server",
		Args:  cobra.NoArgs,

		PreRun: func(cmd *cobra.Command, args []string) {
			log.SetOutput(os.Stdout)
			config.Init(standaloneConfig, configFile)
			setLogLevel()
			config.Debug()
		},

		Run: func(cmd *cobra.Command, args []string) {
			standalone.Start()
		},
	}

	addCommonAPIFlags(standaloneConfig, standaloneCmd)

	var flags = standaloneCmd.Flags()

	if err := standaloneConfig.BindPFlag("backoffice.apiURI", flags.Lookup("basePath")); err != nil {
		panic(err)
	}
	if err := standaloneConfig.BindPFlag("frontend.apiURI", flags.Lookup("basePath")); err != nil {
		panic(err)
	}

	flags.String("backofficeBasePath", config.Config.Backoffice.BasePath, "Backoffice base path")
	if err := standaloneConfig.BindPFlag("backoffice.basePath", flags.Lookup("backofficeBasePath")); err != nil {
		panic(err)
	}

	flags.String("frontendBasePath", config.Config.Frontend.BasePath, "Frontend base path")
	if err := standaloneConfig.BindPFlag("frontend.basePath", flags.Lookup("frontendBasePath")); err != nil {
		panic(err)
	}
	if err := standaloneConfig.BindPFlag("backoffice.frontendURI", flags.Lookup("frontendBasePath")); err != nil {
		panic(err)
	}

	Marcel.AddCommand(standaloneCmd)
}
