package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Zenika/MARCEL/backend/app"
)

var (
	port uint
)

func init() {
	Cmd.Flags().UintVarP(&port, "port", "p", 8090, "Listening port")
}

// Cmd is the root command of Marcel
var Cmd = &cobra.Command{
	Use:   "marcel",
	Short: "Marcel is a configurable plugin based dashboard system",

	Run: func(cmd *cobra.Command, args []string) {
		a := new(app.App)
		a.Initialize()

		a.Run(fmt.Sprintf(":%d", port))
	},
}
