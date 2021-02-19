package demo

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/allez-chauffe/marcel/pkg/standalone"

	log "github.com/sirupsen/logrus"

	"github.com/allez-chauffe/marcel/api/auth"
	"github.com/allez-chauffe/marcel/api/db"
	"github.com/allez-chauffe/marcel/api/db/users"
	"github.com/allez-chauffe/marcel/pkg/config"
	"github.com/allez-chauffe/marcel/pkg/module"
	xexec "github.com/allez-chauffe/marcel/pkg/os/exec"
)

// Module creates a demonstration standalone server module.
func Module() (*module.Module, error) {
	dataDir, err := ioutil.TempDir("", "marcel")
	if err != nil {
		return nil, fmt.Errorf("Could not create temporary directory: %w", err)
	}

	var cfg = config.New()

	cfg.HTTP().SetPort(0)
	cfg.API().SetDataDir(dataDir)
	cfg.API().Auth().SetSecure(false)

	config.SetDefault(cfg)

	log.SetReportCaller(true) // Allows filtering logs by package
	log.SetFormatter(demoFormatter)

	var token string

	return &module.Module{
		Name: "Demo",
		Start: func(_ module.Context, next module.NextFunc) (module.StopFunc, error) {
			log.Infoln("marcel is warming up...")

			if err := next(); err != nil {
				return nil, err
			}

			user := &users.User{
				DisplayName: "Demo",
				Login:       "demo",
				Role:        "admin",
				CreatedAt:   time.Now(),
			}

			if err := db.Users().Insert(user); err != nil {
				return nil, err
			}

			var err error
			token, err = auth.GenerateRefreshJWT(user)
			if err != nil {
				return nil, err
			}

			return nil, nil
		},
		SubModules: []*module.Module{
			standalone.Module(),
		},
		HTTP: module.HTTP{
			OnListen: func(_ module.Context, l net.Listener, srv *http.Server) {
				url := fmt.Sprintf("http://%s%s?token=%s", l.Addr(), module.URI("Backoffice"), token)

				log.Infof("marcel is running at %s", url)

				if err := xexec.Open(url); err != nil {
					log.Errorf("Error while opening browser: %s", err)
				}
			},
		},
	}, nil
}
