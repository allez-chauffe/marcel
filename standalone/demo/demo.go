package demo

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/Zenika/marcel/httputil"
	"github.com/Zenika/marcel/osutil"
	"github.com/Zenika/marcel/standalone"

	log "github.com/sirupsen/logrus"

	"github.com/Zenika/marcel/api/auth"
	"github.com/Zenika/marcel/api/db/users"
	"github.com/Zenika/marcel/config"
	"github.com/Zenika/marcel/module"
)

func Run() int {
	dataDir, err := ioutil.TempDir("", "marcel")
	if err != nil {
		return 1
	}

	var cfg = config.New()

	cfg.HTTP().SetPort(0)
	cfg.API().SetDataDir(dataDir)
	cfg.API().Auth().SetSecure(false)

	config.SetDefault(cfg)

	log.SetReportCaller(true)
	log.SetFormatter(demoFormatter)

	var token string

	var demo = module.Module{
		Name: "Demo",
		Start: func(next module.StartNextFunc) (module.StopFunc, error) {
			log.Infoln("marcel is warming up...")

			if err := next(); err != nil {
				return nil, err
			}

			user := &users.User{
				DisplayName: "Demo",
				Login:       "demo",
				Role:        "user",
				CreatedAt:   time.Now(),
			}

			if err := users.Insert(user); err != nil {
				return nil, err
			}

			var err error
			token, err = auth.GenerateRefreshJWT(user)
			if err != nil {
				return nil, err
			}

			return nil, nil
		},
		SubModules: []module.Module{
			standalone.Module(),
		},
		Http: module.Http{
			OnListen: func(listener net.Listener, srv *http.Server) {
				url := fmt.Sprintf("http://%s%s?token=%s", listener.Addr(), httputil.NormalizeBase(cfg.Backoffice().BasePath()), token)

				log.Infof("marcel is running at %s\n", url)

				time.AfterFunc(100*time.Millisecond, func() {
					osutil.Open(url) // Discard error on purpose
				})
			},
		},
	}

	return demo.Run()
}
