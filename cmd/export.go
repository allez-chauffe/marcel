package cmd

import (
	"github.com/spf13/cobra"

	"github.com/Zenika/marcel/api/db/export"
	"github.com/Zenika/marcel/config"
)

func init() {

	// === export command ===

	var exportFile string

	var exportCmd = &cobra.Command{
		Use:   "export",
		Short: "Exports data from marcel's database",
		Args:  cobra.NoArgs,

		PersistentPreRun: (func(cmd *cobra.Command, args []string) {
			if len(args) > 0 && args[0] != "-" {
				exportFile = args[0]
			}
		}),
	}

	exportCmd.PersistentFlags().StringVar(&config.Config.API.DBFile, "dbFile", config.Config.API.DBFile, "Database file name")

	Marcel.AddCommand(exportCmd)

	// === users command ===

	var usersWPassword bool

	var users = &cobra.Command{
		Use:   "users [FILE]",
		Short: "Exports users from marcel's database",
		Args:  cobra.MaximumNArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			return export.Users(usersWPassword, exportFile)
		},
	}

	users.Flags().BoolVar(&usersWPassword, "withPassword", false, "Include each user's password")

	exportCmd.AddCommand(users)

	// === medias command ===

	var medias = &cobra.Command{
		Use:   "medias [FILE]",
		Short: "Exports medias from marcel's database",
		Args:  cobra.MaximumNArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			return export.Medias(exportFile)
		},
	}

	exportCmd.AddCommand(medias)

	// === plugins command ===

	var plugins = &cobra.Command{
		Use:   "plugins [FILE]",
		Short: "Exports plugins from marcel's database",
		Args:  cobra.MaximumNArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			return export.Plugins(exportFile)
		},
	}

	exportCmd.AddCommand(plugins)

	// === all command ===

	var all = &cobra.Command{
		Use:   "all [FILE]",
		Short: "Exports all data from marcel's database",
		Args:  cobra.MaximumNArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			return export.All(usersWPassword, exportFile)
		},
	}

	all.Flags().BoolVar(&usersWPassword, "withPassword", false, "Include each user's password")

	exportCmd.AddCommand(all)
}
