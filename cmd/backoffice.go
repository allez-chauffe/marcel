package cmd

import (
	"github.com/spf13/cobra"

	"github.com/Zenika/marcel/backoffice"
	"github.com/Zenika/marcel/config"
)

func init() {
	Marcel.AddCommand(backofficeCmd)
}

var backofficeCmd = &cobra.Command{
	Use:   "backoffice",
	Short: "Starts Marcel's backoffice server",
	Args:  cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		return backoffice.ListenAndServe(config.Config.Port, "")
	},
}
