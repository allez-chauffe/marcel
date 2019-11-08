package module

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

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

	// FIXME HTTP

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

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
