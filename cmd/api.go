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
	apiCmd.Flags().UintP("port", "p", config.Config.Port, "Listening port")
	viper.BindPFlag("port", apiCmd.Flags().Lookup("port"))

	apiCmd.Flags().Bool("cors", config.Config.CORS, "Enable CORS for all origins")
	viper.BindPFlag("cors", apiCmd.Flags().Lookup("cors"))

	apiCmd.Flags().String("dbFile", config.Config.DBFile, "Database file name")
	viper.BindPFlag("dbFile", apiCmd.Flags().Lookup("dbFile"))

	apiCmd.Flags().String("pluginsPath", config.Config.PluginsPath, "Plugins directory")
	viper.BindPFlag("pluginsPath", apiCmd.Flags().Lookup("pluginsPath"))

	apiCmd.Flags().Bool("secure", config.Config.Auth.Secure, "Use secured cookies")
	viper.BindPFlag("auth.secure", apiCmd.Flags().Lookup("secure"))

	apiCmd.Flags().Duration("authExpiration", config.Config.Auth.AuthExpiration, "Authentication token expiration")
	viper.BindPFlag("auth.authExpiration", apiCmd.Flags().Lookup("authExpiration"))

	apiCmd.Flags().Duration("refreshExpiration", config.Config.Auth.RefreshExpiration, "Refresh token expiration")
	viper.BindPFlag("auth.refreshExpiration", apiCmd.Flags().Lookup("refreshExpiration"))

	Marcel.AddCommand(apiCmd)
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Starts marcel's api server",
	Args:  cobra.NoArgs,

	PreRun: func(cmd *cobra.Command, args []string) {
		log.SetOutput(os.Stdout)
		config.Init(configFile)
		setLogLevel()
		debugConfig()
	},

	Run: func(cmd *cobra.Command, args []string) {
		a := new(api.App)
		a.Initialize()
		a.Run()
	},
}

// LogLevel implements a pflag.Value with logrus.Level
type LogLevel log.Level

func (l *LogLevel) String() string {
	return log.Level(*l).String()
}

func (l *LogLevel) Set(s string) error {
	v, err := log.ParseLevel(s)
	if err != nil {
		return err
	}
	*l = LogLevel(v)
	return nil
}

func (l *LogLevel) Type() string {
	return "log.Level"
}

func setLogLevel() {
	log.SetLevel(config.Config.LogLevel)
	log.Infof("Log level set to %s", config.Config.LogLevel)
}

func debugConfig() {
	log.Debugf("Config: %+v", config.Config)
}
