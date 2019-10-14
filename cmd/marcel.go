package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Zenika/marcel/standalone/demo"

	isatty "github.com/mattn/go-isatty"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string

func init() {
	var ll = log.InfoLevel
	Marcel.PersistentFlags().Var((*logLevel)(&ll), "logLevel", fmt.Sprintf("Log level: %s, %s, %s, %s or %s", log.TraceLevel, log.DebugLevel, log.InfoLevel, log.WarnLevel, log.ErrorLevel))

	Marcel.PersistentFlags().StringVarP(&configFile, "configFile", "c", "", fmt.Sprintf("Config file (default /etc/marcel/config.xxx or ./config.xxx, supports %s)", strings.Join(viper.SupportedExts, " ")))
}

// Marcel is the root command of marcel
var Marcel = &cobra.Command{
	Use:           "marcel",
	Short:         "marcel is a configurable plugin based dashboard system",
	Args:          cobra.NoArgs,
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, _ []string) error {
		if isatty.IsTerminal(os.Stdin.Fd()) && isatty.IsTerminal(os.Stdout.Fd()) {
			return startInteractive(cmd.Usage)
		}

		return cmd.Usage()
	},
}

func startInteractive(usage func() error) error {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("You haven't specified any command, would you like to start a demonstration server ? [Y/n/h] ")

		if scanner.Scan() {
			answer := strings.ToLower(strings.TrimLeft(scanner.Text(), " "))

			switch {
			case answer == "":
				fallthrough
			case strings.HasPrefix(answer, "y"):
				fmt.Println()
				return demo.StartServer()
			case strings.HasPrefix(answer, "n"):
				fallthrough
			case strings.HasPrefix(answer, "h"):
				fmt.Println()
				return usage()
			default:
				fmt.Printf("Answer %#v is invalid.\n\n", scanner.Text())
			}
		} else {
			break
		}
	}

	return nil
}
