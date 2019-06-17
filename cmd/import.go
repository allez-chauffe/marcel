package cmd

import (
	"github.com/spf13/cobra"

	"github.com/Zenika/MARCEL/api/db/imp0rt"
	"github.com/Zenika/MARCEL/config"
)

func init() {

	// === import command ===

	var importFile string

	var importCmd = &cobra.Command{
		Use:   "import",
		Short: "Imports data from Marcel's database",
		Args:  cobra.NoArgs,

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 && args[0] != "-" {
				importFile = args[0]
			}
		},
	}

	importCmd.PersistentFlags().StringVar(&config.Config.DBFile, "dbFile", config.Config.DBFile, "Database file name")

	Marcel.AddCommand(importCmd)

	// === users command ===

	var users = &cobra.Command{
		Use:   "users [FILE]",
		Short: "Imports users from Marcel's database",
		Args:  cobra.MaximumNArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			return imp0rt.Users(importFile)
		},
	}

	importCmd.AddCommand(users)

	// === medias command ===

	var medias = &cobra.Command{
		Use:   "medias [FILE]",
		Short: "Imports medias from Marcel's database",
		Args:  cobra.MaximumNArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			return imp0rt.Medias(importFile)
		},
	}

	importCmd.AddCommand(medias)

	// === plugins command ===

	var plugins = &cobra.Command{
		Use:   "plugins [FILE]",
		Short: "Imports plugins from Marcel's database",
		Args:  cobra.MaximumNArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			return imp0rt.Plugins(importFile)
		},
	}

	importCmd.AddCommand(plugins)

	// === all command ===

	var all = &cobra.Command{
		Use:   "all [FILE]",
		Short: "Imports all data from Marcel's database",
		Args:  cobra.MaximumNArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			return imp0rt.All(importFile)
		},
	}

	importCmd.AddCommand(all)
}
