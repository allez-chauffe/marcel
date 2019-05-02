package cmd

import (
	"github.com/spf13/cobra"
)

// Marcel is the root command of Marcel
var Marcel = &cobra.Command{
	Use:   "marcel",
	Short: "Marcel is a configurable plugin based dashboard system",
	Args:  cobra.NoArgs,
}
