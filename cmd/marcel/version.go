package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/allez-chauffe/marcel/pkg/version"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Displays version information",
		Args:  cobra.NoArgs,

		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("%s rev: %s\n", version.Version(), version.Revision())
		},
	})
}
