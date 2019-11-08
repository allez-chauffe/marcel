package module

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sort"
	"strings"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"

	"github.com/Zenika/marcel/config"
)

type StartNextFunc func() error

type StopFunc func() error

type StartFunc func(StartNextFunc) (StopFunc, error)

type Module struct {
	Name       string
	Start      StartFunc
	SubModules []Module
	Http
}

type Http struct {
	BasePath string
	Setup    func(*mux.Router)
	OnListen func(net.Listener, *http.Server)
}

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

	fmt.Print("\r") // FIXME only in tty
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

	log.Infof("Module %s started", m.Name)

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
	log.Infof("Starting %s's HTTP server...", m.Name)

	var router = mux.NewRouter()

	var hasHTTP = m.setupRouter(router)

	if !hasHTTP {
		return nil, nil
	}

	var addr = fmt.Sprintf(":%d", config.Default().HTTP().Port())
	var listener, err = net.Listen("tcp", addr)
	if err != nil {
		return nil, err // FIXME wrap
	}

	var handler http.Handler = router

	if config.Default().HTTP().CORS() {
		handler = cors.New(cors.Options{
			AllowOriginFunc:  func(origin string) bool { return true },
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
			AllowedHeaders:   []string{"*"},
			AllowCredentials: true,
		}).Handler(handler)
		log.Warn("CORS is enabled")
	}

	var srv = &http.Server{
		Handler: handler,
	}

	log.Infof("%s's HTTP server listening on %s", m.Name, listener.Addr())

	go func() {
		if err := srv.Serve(listener); err != nil {
			// FIXME
		}
	}()

	m.notifyOnServe(listener, srv)

	return srv, nil
}

func (m *Module) setupRouter(parentRouter *mux.Router) bool {
	var hasHTTP = false
	var router = parentRouter

	if m.BasePath != "" {
		router = parentRouter.PathPrefix(m.BasePath).Subrouter()
		log.Debugf("Created subrouter for %s at %s", m.Name, m.BasePath)
	}

	if m.Setup != nil {
		hasHTTP = true
		m.Setup(router)
		log.Debugf("Configured subrouter for %s", m.Name)
	}

	sort.Sort(byBasePath(m.SubModules))

	for _, subM := range m.SubModules {
		hasHTTP = subM.setupRouter(router) || hasHTTP
	}

	return hasHTTP
}

func (m *Module) notifyOnServe(listener net.Listener, srv *http.Server) {
	if m.OnListen != nil {
		m.OnListen(listener, srv)
	}

	for _, subM := range m.SubModules {
		subM.notifyOnServe(listener, srv)
	}
}

type byBasePath []Module

var _ sort.Interface = byBasePath(nil)

func (modules byBasePath) Len() int {
	return len(modules)
}

func (modules byBasePath) Less(i, j int) bool {
	return strings.Count(modules[i].BasePath, "/") > strings.Count(modules[j].BasePath, "/")
}

func (modules byBasePath) Swap(i, j int) {
	modules[i], modules[j] = modules[j], modules[i]
}
