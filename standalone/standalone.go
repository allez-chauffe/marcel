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
	configure func(*mux.Router)
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

func Start() {
	var a = api.New()
	a.Initialize()

	var r = mux.NewRouter()

	var configurers = routerConfigurers{
		routerConfigurer{httputil.NormalizeBase(config.Config.API.BasePath), a.ConfigureRouter},
		routerConfigurer{httputil.NormalizeBase(config.Config.Frontend.BasePath), frontend.ConfigureRouter},
		routerConfigurer{httputil.NormalizeBase(config.Config.Backoffice.BasePath), backoffice.ConfigureRouter},
	}

	sort.Sort(configurers)

	for _, configurer := range configurers {
		configurer.configure(r)
	}

	log.Infof("Starting standalone server on port %d...", config.Config.Standalone.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Standalone.Port), r))
}
