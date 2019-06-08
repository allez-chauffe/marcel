package cmd

import (
	"github.com/spf13/cobra"

	"github.com/Zenika/MARCEL/api/db/export"
	"github.com/Zenika/MARCEL/config"
)

var usersWPassword bool

func init() {
	var exportCmd = &cobra.Command{
		Use:   "export",
		Short: "Exports data from Marcel's database",
		Args:  cobra.NoArgs,
	}

	exportCmd.PersistentFlags().StringVar(&config.Config.DBFile, "dbFile", config.Config.DBFile, "Database file name")

	Marcel.AddCommand(exportCmd)

	var users = &cobra.Command{
		Use:   "users FILE",
		Short: "Exports users from Marcel's database",
		Args:  cobra.ExactArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			return export.Users(usersWPassword, args[0])
		},
	}

	users.Flags().BoolVar(&usersWPassword, "withPassword", false, "Include each user's password")

	exportCmd.AddCommand(users)

	var medias = &cobra.Command{
		Use:   "medias FILE",
		Short: "Exports medias from Marcel's database",
		Args:  cobra.ExactArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			return export.Medias(args[0])
		},
	}

	exportCmd.AddCommand(medias)

	var plugins = &cobra.Command{
		Use:   "plugins FILE",
		Short: "Exports plugins from Marcel's database",
		Args:  cobra.ExactArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			return export.Plugins(args[0])
		},
	}

	exportCmd.AddCommand(plugins)

	var all = &cobra.Command{
		Use:   "all FILE",
		Short: "Exports all data from Marcel's database",
		Args:  cobra.ExactArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			return export.All(usersWPassword, args[0])
		},
	}

	all.Flags().BoolVar(&usersWPassword, "withPassword", false, "Include each user's password")

	exportCmd.AddCommand(all)
}
