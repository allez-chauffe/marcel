package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Zenika/MARCEL/api/db/export"
	"github.com/Zenika/MARCEL/config"
)

var usersWPassword bool

func init() {
	exportCmd.PersistentFlags().String("dbFile", config.Config.DBFile, "Database file name")
	viper.BindPFlag("dbFile", exportCmd.PersistentFlags().Lookup("dbFile"))

	Marcel.AddCommand(exportCmd)

	usersCmd.Flags().BoolVar(&usersWPassword, "withPassword", false, "Include each user's password")

	exportCmd.AddCommand(usersCmd)

	exportCmd.AddCommand(mediasCmd)

	exportCmd.AddCommand(pluginsCmd)

	allCmd.Flags().BoolVar(&usersWPassword, "withPassword", false, "Include each user's password")

	exportCmd.AddCommand(allCmd)
}

var (
	exportCmd = &cobra.Command{
		Use:   "export",
		Short: "Exports data from Marcel's database",
		Args:  cobra.NoArgs,
	}

	usersCmd = &cobra.Command{
		Use:   "users FILE",
		Short: "Exports users from Marcel's database",
		Args:  cobra.ExactArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			return export.Users(usersWPassword, args[0])
		},
	}

	mediasCmd = &cobra.Command{
		Use:   "medias FILE",
		Short: "Exports medias from Marcel's database",
		Args:  cobra.ExactArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			return export.Medias(args[0])
		},
	}

	pluginsCmd = &cobra.Command{
		Use:   "plugins FILE",
		Short: "Exports plugins from Marcel's database",
		Args:  cobra.ExactArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			return export.Plugins(args[0])
		},
	}

	allCmd = &cobra.Command{
		Use:   "all FILE",
		Short: "Exports all data from Marcel's database",
		Args:  cobra.ExactArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			return export.All(usersWPassword, args[0])
		},
	}
)
