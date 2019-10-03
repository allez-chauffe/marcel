package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Zenika/marcel/version"
)

func init() {
	Marcel.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Displays version information",
		Args:  cobra.NoArgs,

		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("%s rev: %s\n", version.Version, version.Revision)
		},
	})
}
