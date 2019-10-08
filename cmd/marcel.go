package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/Zenika/marcel/api/auth"
	"github.com/Zenika/marcel/api/db/users"
	"github.com/Zenika/marcel/httputil"

	"github.com/Zenika/marcel/osutil"

	"github.com/Zenika/marcel/config"
	"github.com/Zenika/marcel/standalone"

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
				return startDemoServer()
			case strings.HasPrefix(answer, "n"):
				fallthrough
			case strings.HasPrefix(answer, "h"):
				fmt.Println()
				return usage()
			default:
				fmt.Printf("Answer %#v is invalid.\n", scanner.Text())
			}
		} else {
			break
		}
	}

	return nil
}

func startDemoServer() error {
	dataDir, err := ioutil.TempDir("", "marcel")
	if err != nil {
		return err
	}

	var cfg = config.New()

	cfg.API().SetDataDir(dataDir)
	cfg.API().Auth().SetSecure(false)

	config.SetDefault(cfg)

	log.SetLevel(log.FatalLevel)
	// FIXME rather change format
	log.SetOutput(ioutil.Discard)

	fmt.Println("marcel is warming up...")

	done := make(chan error)
	if err := standalone.Start(done); err != nil {
		return err
	}

	user := &users.User{
		DisplayName: "Demo",
		Login:       "demo",
		Role:        "user",
		CreatedAt:   time.Now(),
	}

	if err := users.Insert(user); err != nil {
		return err
	}

	token, err := auth.GenerateRefreshJWT(user)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("http://localhost:%d%s?token=%s", config.Default().Standalone().Port(), httputil.NormalizeBase(cfg.Backoffice().BasePath()), token)

	fmt.Printf("marcel is running at %s\n", url)

	osutil.Open(url) // Discard error on purpose

	return <-done
}
