package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Zenika/MARCEL/api/db/imp0rt"
	"github.com/Zenika/MARCEL/config"
)

func init() {
	var importCmd = &cobra.Command{
		Use:   "import",
		Short: "Imports data from Marcel's database",
		Args:  cobra.NoArgs,
	}

	importCmd.PersistentFlags().String("dbFile", config.Config.DBFile, "Database file name")
	viper.BindPFlag("dbFile", importCmd.PersistentFlags().Lookup("dbFile"))

	Marcel.AddCommand(importCmd)

	var usersCmd = &cobra.Command{
		Use:   "users FILE",
		Short: "Imports users from Marcel's database",
		Args:  cobra.ExactArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			return imp0rt.Users(args[0])
		},
	}

	importCmd.AddCommand(usersCmd)

	// mediasCmd = &cobra.Command{
	// 	Use:   "medias FILE",
	// 	Short: "Imports medias from Marcel's database",
	// 	Args:  cobra.ExactArgs(1),

	// 	RunE: func(cmd *cobra.Command, args []string) error {
	// 		return export.Medias(args[0])
	// 	},
	// }

	// importCmd.AddCommand(mediasCmd)

	// pluginsCmd = &cobra.Command{
	// 	Use:   "plugins FILE",
	// 	Short: "Imports plugins from Marcel's database",
	// 	Args:  cobra.ExactArgs(1),

	// 	RunE: func(cmd *cobra.Command, args []string) error {
	// 		return export.Plugins(args[0])
	// 	},
	// }

	// importCmd.AddCommand(pluginsCmd)

	// allCmd = &cobra.Command{
	// 	Use:   "all FILE",
	// 	Short: "Imports all data from Marcel's database",
	// 	Args:  cobra.ExactArgs(1),

	// 	RunE: func(cmd *cobra.Command, args []string) error {
	// 		return export.All(usersWPassword, args[0])
	// 	},
	// }

	// importCmd.AddCommand(allCmd)
}
