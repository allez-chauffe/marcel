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

	// FIXME find a way to reuse default config from standalone cmd
	cfg.API().SetBasePath("/api")
	cfg.API().SetDBFile("marcel.db")
	cfg.API().SetDataDir(dataDir)
	cfg.API().SetMediasDir("medias")
	cfg.API().SetPluginsDir("plugins")

	cfg.API().Auth().SetExpiration(8 * time.Hour)
	cfg.API().Auth().SetRefreshExpiration(15 * 24 * time.Hour)
	cfg.API().Auth().SetSecure(false)

	cfg.Backoffice().SetAPIURI("/api")
	cfg.Backoffice().SetBasePath("/")
	cfg.Backoffice().SetFrontendURI("/front")

	cfg.Frontend().SetAPIURI("/api")
	cfg.Frontend().SetBasePath("/front")

	cfg.Standalone().SetPort(8090)

	config.SetDefault(cfg)

	log.SetLevel(log.FatalLevel)
	// FIXME rather change format
	log.SetOutput(ioutil.Discard)

	fmt.Println("marcel is starting...")

	done := make(chan error)
	if err := standalone.Start(done); err != nil {
		return err
	}

	admin, err := users.GetByLogin("admin")
	if err != nil {
		return err
	}

	token, err := auth.GenerateRefreshJWT(admin)
	if err != nil {
		return err
	}

	//FIXME use backoffice basePath
	url := fmt.Sprintf("http://localhost:%d/?token=%s", config.Default().Standalone().Port(), token)

	fmt.Printf("marcel is running at %s\n", url)

	osutil.Open(url) // Discard error on purpose

	return <-done
}
