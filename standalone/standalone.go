package standalone

import (
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/Zenika/marcel/api"
	"github.com/Zenika/marcel/backoffice"
	"github.com/Zenika/marcel/config"
	"github.com/Zenika/marcel/frontend"
	"github.com/Zenika/marcel/httputil"
)

type routerConfigurer struct {
	base      string
	configure func(*mux.Router) error
}

type routerConfigurers []routerConfigurer

var _ sort.Interface = routerConfigurers(nil)

func (configurers routerConfigurers) Len() int {
	return len(configurers)
}

func (configurers routerConfigurers) Less(i, j int) bool {
	return strings.Count(configurers[i].base, "/") > strings.Count(configurers[j].base, "/")
}

func (configurers routerConfigurers) Swap(i, j int) {
	configurers[i], configurers[j] = configurers[j], configurers[i]
}

func Start(done chan<- error) error {
	var a = api.New()
	if err := a.Init(); err != nil {
		return err
	}

	var r = mux.NewRouter()

	var configurers = routerConfigurers{
		routerConfigurer{httputil.NormalizeBase(config.Default().API().BasePath()), a.ConfigureRouter},
		routerConfigurer{httputil.NormalizeBase(config.Default().Frontend().BasePath()), frontend.ConfigureRouter},
		routerConfigurer{httputil.NormalizeBase(config.Default().Backoffice().BasePath()), backoffice.ConfigureRouter},
	}

	sort.Sort(configurers)

	for _, configurer := range configurers {
		if err := configurer.configure(r); err != nil {
			return err
		}
	}

	log.Infof("Standalone server listening on %d...", config.Default().Standalone().Port())

	if done == nil {
		return http.ListenAndServe(fmt.Sprintf(":%d", config.Default().Standalone().Port()), r)
	}

	go func() {
		done <- http.ListenAndServe(fmt.Sprintf(":%d", config.Default().Standalone().Port()), r)
	}()

	return nil
}
