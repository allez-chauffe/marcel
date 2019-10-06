package cmd

import (
	"github.com/spf13/cobra"

	"github.com/Zenika/marcel/api/db/imp0rt"
	"github.com/Zenika/marcel/config"
)

func init() {

	var cfg = config.New()

	var importFile string

	var cmd = &cobra.Command{
		Use:   "import",
		Short: "Imports data from marcel's database",
		Args:  cobra.NoArgs,

		PersistentPreRun: func(_ *cobra.Command, args []string) {
			bindLogLevel(cfg)
			config.SetDefault(cfg)
			setLogLevel(cfg)

			if len(args) > 0 && args[0] != "-" {
				importFile = args[0]
			}
		},
	}

	var flags = cmd.PersistentFlags()

	if _, err := cfg.FlagString(flags, "dbFile", "marcel.db", "Database file name"); err != nil {
		panic(err)
	}

	Marcel.AddCommand(cmd)

	cmd.AddCommand(&cobra.Command{
		Use:   "users [FILE]",
		Short: "Imports users from marcel's database",
		Args:  cobra.MaximumNArgs(1),

		RunE: func(_ *cobra.Command, _ []string) error {
			return imp0rt.Users(importFile)
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "medias [FILE]",
		Short: "Imports medias from marcel's database",
		Args:  cobra.MaximumNArgs(1),

		RunE: func(_ *cobra.Command, _ []string) error {
			return imp0rt.Medias(importFile)
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "plugins [FILE]",
		Short: "Imports plugins from marcel's database",
		Args:  cobra.MaximumNArgs(1),

		RunE: func(_ *cobra.Command, _ []string) error {
			return imp0rt.Plugins(importFile)
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "all [FILE]",
		Short: "Imports all data from marcel's database",
		Args:  cobra.MaximumNArgs(1),

		RunE: func(_ *cobra.Command, _ []string) error {
			return imp0rt.All(importFile)
		},
	})
}
