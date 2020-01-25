package cmd

import (
	"github.com/spf13/cobra"

	"github.com/Zenika/marcel/api/db/export"
	"github.com/Zenika/marcel/config"
)

func init() {

	var cfg = config.New()

	var exportFile string

	var cmd = &cobra.Command{
		Use:   "export",
		Short: "Exports data from marcel's database",
		Args:  cobra.NoArgs,

		PersistentPreRun: func(_ *cobra.Command, args []string) {
			bindLogLevel(cfg)
			config.SetDefault(cfg)
			setLogLevel(cfg)

			if len(args) > 0 && args[0] != "-" {
				exportFile = args[0]
			}
		},
	}

	if _, err := cfg.FlagString(cmd.PersistentFlags(), "dbFile", cfg.API().DBFile(), "Database file name", "api.dbFile"); err != nil {
		panic(err)
	}

	Marcel.AddCommand(cmd)

	var usersWPassword bool
	var pretty bool

	var users = &cobra.Command{
		Use:   "users [FILE]",
		Short: "Exports users from marcel's database",
		Args:  cobra.MaximumNArgs(1),

		RunE: func(_ *cobra.Command, _ []string) error {
			return export.Users(pretty, usersWPassword, exportFile)
		},
	}

	users.Flags().BoolVar(&usersWPassword, "withPassword", false, "Include each user's password")
	users.Flags().BoolVar(&pretty, "pretty", false, "Indent json export")

	cmd.AddCommand(users)

	var medias = &cobra.Command{
		Use:   "medias [FILE]",
		Short: "Exports medias from marcel's database",
		Args:  cobra.MaximumNArgs(1),

		RunE: func(_ *cobra.Command, _ []string) error {
			return export.Medias(pretty, exportFile)
		},
	}

	medias.Flags().BoolVar(&pretty, "pretty", false, "Indent json export")
	cmd.AddCommand(medias)

	var plugins = &cobra.Command{
		Use:   "plugins [FILE]",
		Short: "Exports plugins from marcel's database",
		Args:  cobra.MaximumNArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			return export.Plugins(pretty, exportFile)
		},
	}

	plugins.Flags().BoolVar(&pretty, "pretty", false, "Indent json export")
	cmd.AddCommand(plugins)

	var all = &cobra.Command{
		Use:   "all [FILE]",
		Short: "Exports all data from marcel's database",
		Args:  cobra.MaximumNArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			return export.All(pretty, usersWPassword, exportFile)
		},
	}

	all.Flags().BoolVar(&usersWPassword, "withPassword", false, "Include each user's password")
	all.Flags().BoolVar(&pretty, "pretty", false, "Indent json export")

	cmd.AddCommand(all)
}
