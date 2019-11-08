package module

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type StartNextFunc func() error

type StopFunc func() error

type StartFunc func(StartNextFunc) (StopFunc, error)

type Module struct {
	Name       string
	Start      StartFunc
	SubModules []Module
	*Http
}

type Http struct {
	BasePath string
	Setup    func(*mux.Router) error
}

func (m Module) Run() (exitCode int) {
	log.Infof("Starting module %s...", m.Name)

	var startRes = m.start()
	if startRes.err != nil {
		log.Errorln(startRes.err)
		exitCode = 1
		return
	}
	defer func() {
		if err := startRes.stop(); err != nil {
			log.Errorln(err)
			exitCode = 1
		}
	}()

	var httpSrv, err = m.startHTTP()
	if err != nil {
		// FIXME log
		exitCode = 1
		return
	}
	defer func() {
		log.Infof("Shutting down %s's HTTP server...", m.Name)

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()

		// FIXME manage websockets
		if err := httpSrv.Shutdown(ctx); err != nil {
			// FIXME log
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	var signal = <-ch

	fmt.Print("\r")
	log.Infof("Caught signal %s", signal)

	return
}

type startResult struct {
	stop StopFunc
	err  error
}

func (m Module) start() startResult {
	log.Debugf("Starting module %s...", m.Name)

	var stopFuncs = make([]StopFunc, 0, len(m.SubModules))

	var next = func() error {
		if len(m.SubModules) == 0 {
			return nil
		}

		startResCh := make(chan startResult)
		for _, subM := range m.SubModules {
			go func(subM Module) {
				defer func() {
					if r := recover(); r != nil {
						if err, ok := r.(error); ok {
							startResCh <- startResult{nil, fmt.Errorf("panic while starting module %s: %w", subM.Name, err)}
						} else {
							startResCh <- startResult{nil, fmt.Errorf("panic while starting module %s: %s", subM.Name, r)}
						}
					}
				}()
				startResCh <- subM.start()
			}(subM)
		}

		var hasError bool
		for range m.SubModules {
			var startRes = <-startResCh
			if startRes.err != nil {
				log.Errorf("Error while starting %s's submodules: %s", m.Name, startRes.err)
				hasError = true
			}
			if startRes.stop != nil {
				stopFuncs = append(stopFuncs, startRes.stop)
			}
		}

		if hasError {
			return fmt.Errorf("Error while starting %s's submodules", m.Name)
		}

		return nil
	}

	var stop, err = m.callStart(next)

	return startResult{
		func() error {
			var hasError = false

			if stop != nil {
				if err := stop(); err != nil {
					log.Errorf("Error while stopping module %s: %s", m.Name, err)
					hasError = true
				}
			}

			for _, subStop := range stopFuncs {
				// FIXME concurrent subStops
				if err := subStop(); err != nil {
					log.Errorf("Error while stopping %s's submodules: %s", m.Name, err)
					hasError = true
				}
			}

			if hasError {
				return fmt.Errorf("Error while stopping module %s", m.Name)
			}

			return nil
		},
		err,
	}
}

func (m Module) callStart(next StartNextFunc) (StopFunc, error) {
	if m.Start == nil {
		return nil, next()
	}

	var nextCalled = false
	var callNext StartNextFunc = func() error {
		if nextCalled {
			return fmt.Errorf("next already called while starting module %s", m.Name)
		}
		nextCalled = true

		return next()
	}

	var stop, err = m.Start(callNext)
	if err != nil {
		return stop, fmt.Errorf("Error while starting module %s: %w", m.Name, err)
	}

	if !nextCalled {
		return stop, fmt.Errorf("next never called while starting module %s", m.Name)
	}

	return stop, nil
}

func (m *Module) startHTTP() (*http.Server, error) {
	var router = mux.NewRouter()

	var hasHTTP, err = m.setupRouter(router)
	if err != nil {
		return nil, err
	}

	if !hasHTTP {
		return nil, nil
	}

	var srv = &http.Server{
		Addr:    ":8090", // FIXME Addr
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			// FIXME
		}
	}()

	return srv, nil
}

func (m *Module) setupRouter(parentRouter *mux.Router) (bool, error) {
	var hasHTTP = false
	var router = parentRouter

	if m.Http != nil {
		if m.Http.BasePath != "" {
			router = parentRouter.PathPrefix(m.Http.BasePath).Subrouter()
		}
		if m.Http.Setup != nil {
			hasHTTP = true
			if err := m.Http.Setup(router); err != nil {
				return false, err // FIXME wrap
			}
		}
	}

	for _, subM := range m.SubModules {
		var hasSubHTTP, err = subM.setupRouter(router)
		if err != nil {
			return false, err
		}
		hasHTTP = hasHTTP || hasSubHTTP
	}

	return hasHTTP, nil
}
