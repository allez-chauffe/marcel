package demo

import (
	"fmt"
	"io/ioutil"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Zenika/marcel/api/auth"
	"github.com/Zenika/marcel/api/db/users"
	"github.com/Zenika/marcel/config"
	"github.com/Zenika/marcel/httputil"
	"github.com/Zenika/marcel/osutil"
	"github.com/Zenika/marcel/standalone"
)

func StartServer() error {
	dataDir, err := ioutil.TempDir("", "marcel")
	if err != nil {
		return err
	}

	var cfg = config.New()

	cfg.API().SetDataDir(dataDir)
	cfg.API().Auth().SetSecure(false)

	config.SetDefault(cfg)

	// log.SetLevel(log.FatalLevel)
	log.SetReportCaller(true)
	log.SetFormatter(demoFormatter)

	log.Infoln("marcel is warming up...")

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

	select {
	case err := <-done:
		return err
	case <-time.After(50 * time.Millisecond):
	}

	url := fmt.Sprintf("http://localhost:%d%s?token=%s", config.Default().HTTP().Port(), httputil.NormalizeBase(cfg.Backoffice().BasePath()), token)

	log.Infof("marcel is running at %s\n", url)

	osutil.Open(url) // Discard error on purpose

	return <-done
}
