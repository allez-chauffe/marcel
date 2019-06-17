package cmd

import (
	"github.com/spf13/cobra"

	"github.com/Zenika/MARCEL/api/db/imp0rt"
	"github.com/Zenika/MARCEL/config"
)

func init() {
	var importCmd = &cobra.Command{
		Use:   "import",
		Short: "Imports data from Marcel's database",
		Args:  cobra.NoArgs,
	}

	importCmd.PersistentFlags().StringVar(&config.Config.DBFile, "dbFile", config.Config.DBFile, "Database file name")

	Marcel.AddCommand(importCmd)

	var users = &cobra.Command{
		Use:   "users FILE",
		Short: "Imports users from Marcel's database",
		Args:  cobra.ExactArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			return imp0rt.Users(args[0])
		},
	}

	importCmd.AddCommand(users)

	var medias = &cobra.Command{
		Use:   "medias FILE",
		Short: "Imports medias from Marcel's database",
		Args:  cobra.ExactArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			return imp0rt.Medias(args[0])
		},
	}

	importCmd.AddCommand(medias)

	var plugins = &cobra.Command{
		Use:   "plugins FILE",
		Short: "Imports plugins from Marcel's database",
		Args:  cobra.ExactArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			return imp0rt.Plugins(args[0])
		},
	}

	importCmd.AddCommand(plugins)

	var all = &cobra.Command{
		Use:   "all FILE",
		Short: "Imports all data from Marcel's database",
		Args:  cobra.ExactArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			return imp0rt.All(args[0])
		},
	}

	importCmd.AddCommand(all)
}
