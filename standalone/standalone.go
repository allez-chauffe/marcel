package standalone

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
	"sync"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/Zenika/marcel/api"
	"github.com/Zenika/marcel/config"
)

type initializer interface {
	Init() error
}

type routerConfigurer interface {
	BasePath() string
	Configure(*mux.Router) error
}

type routerConfigurers []routerConfigurer

var _ sort.Interface = routerConfigurers(nil)

func (configurers routerConfigurers) Len() int {
	return len(configurers)
}

func (configurers routerConfigurers) Less(i, j int) bool {
	return strings.Count(configurers[i].BasePath(), "/") > strings.Count(configurers[j].BasePath(), "/")
}

func (configurers routerConfigurers) Swap(i, j int) {
	configurers[i], configurers[j] = configurers[j], configurers[i]
}

var modules = []interface{}{
	api.New(),
	// FIXME backoffice
	// FIXME frontend
}

func Start(done chan<- error) error {
	if done == nil {
		return start()
	}

	go func() {
		done <- start()
	}()

	return nil
}

func start() error {
	if err := initModules(); err != nil {
		return err
	}

	r, err := initRouter()
	if err != nil {
		return err
	}

	log.Infof("Standalone server listening on %d...", config.Default().Standalone().Port())

	return http.ListenAndServe(fmt.Sprintf(":%d", config.Default().Standalone().Port()), r)
}

func initModules() error {
	var wg sync.WaitGroup
	errs := make(chan error)

	for _, m := range modules {
		if i, ok := m.(initializer); ok {
			println("init")
			wg.Add(1)
			go func(i initializer) {
				var err error
				defer func() {
					if err != nil {
						errs <- fmt.Errorf("Error during initialization: %w", err)
					}
					if r := recover(); r != nil {
						if err, ok := r.(error); ok {
							errs <- fmt.Errorf("Error during initialization: %w", err)
						} else {
							errs <- fmt.Errorf("Error during initialization: %s", r)
						}
					}
					wg.Done()
				}()
				err = i.Init()
			}(i)
		}
	}

	go func() {
		wg.Wait()
		close(errs)
	}()

	println("oho")

	err := <-errs

	println("aha")

	// Drain errs
	go func() {
		for range errs {
		}
	}()

	return err
}

func initRouter() (*mux.Router, error) {
	var r = mux.NewRouter()

	var configurers = routerConfigurers{}

	for _, m := range modules {
		if rc, ok := m.(routerConfigurer); ok {
			configurers = append(configurers, rc)
		}
	}

	sort.Sort(configurers)

	for _, configurer := range configurers {
		if err := configurer.Configure(r); err != nil {
			return nil, err
		}
	}

	return r, nil
}
