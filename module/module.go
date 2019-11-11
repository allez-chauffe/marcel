package module

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Zenika/marcel/osutil"

	log "github.com/sirupsen/logrus"
)

// StartFunc is a module's start function.
// It must call next.
// It may return a StopFunc if necessary.
type StartFunc func(next NextFunc) (StopFunc, error)

// StopFunc is a module's stop function.
type StopFunc func() error // FIXME use a NextFunc

// NextFunc starts the submodules of the current module.
type NextFunc func() error

// Module describes how a module should run.
type Module struct {
	Name       string
	Start      StartFunc
	SubModules []Module
	HTTP
}

// Run run's a module tree.
func (m Module) Run() (exitCode int) {
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
		log.Errorf("Error while starting %s's HTTP: %s", m.Name, err)
		exitCode = 1
		return
	}
	defer func() {
		log.Infof("Shutting down %s's HTTP server...", m.Name)

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()

		// FIXME manage websockets
		if err := httpSrv.Shutdown(ctx); err != nil {
			log.Errorf("Error while shutting down %s's HTTP: %s", m.Name, err)
		}

		log.Infof("%s's HTTP server stopped", m.Name)
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	var signal = <-ch

	if osutil.IsInteractive() {
		fmt.Print("\r")
	}
	log.Infof("Caught signal %s", signal)

	return
}

type startResult struct {
	stop StopFunc
	err  error
}

func (m Module) start() startResult {
	log.Infof("Starting module %s...", m.Name)

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

	if err == nil {
		log.Infof("Module %s started", m.Name)
	}

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

func (m Module) callStart(next NextFunc) (StopFunc, error) {
	if m.Start == nil {
		return nil, next()
	}

	var nextCalled = false
	var callNext NextFunc = func() error {
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
