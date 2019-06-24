package cmd

import (
	"github.com/spf13/cobra"

	"github.com/Zenika/marcel/backoffice"
)

func init() {
	var basePath = "/"

	var backofficeCmd = &cobra.Command{
		Use:   "backoffice",
		Short: "Starts Marcel's backoffice server",
		Args:  cobra.NoArgs,

		RunE: func(cmd *cobra.Command, args []string) error {
			return backoffice.Start(basePath)
		},
	}

	backofficeCmd.Flags().StringVar(&basePath, "basePath", basePath, "Base path")

	Marcel.AddCommand(backofficeCmd)
}
