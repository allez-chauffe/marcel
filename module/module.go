package module

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Zenika/marcel/config"
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
	Setup    func(*mux.Router)
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

	// FIXME put back cors from api.Start...
	// var h http.Handler = r

	// if config.Default().API().CORS() {
	// 	h = cors.New(cors.Options{
	// 		AllowOriginFunc:  func(origin string) bool { return true },
	// 		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	// 		AllowedHeaders:   []string{"*"},
	// 		AllowCredentials: true,
	// 	}).Handler(h)
	// 	log.Warn("CORS is enabled")
	// }

	var srv = &http.Server{
		Handler: router,
	}

	log.Infof("%s's HTTP server listening on %s", m.Name, listener.Addr())

	go func() {
		if err := srv.Serve(listener); err != nil {
			// FIXME
		}
	}()

	return srv, nil
}

func (m *Module) setupRouter(parentRouter *mux.Router) bool {
	var hasHTTP = false
	var router = parentRouter

	if m.Http != nil {
		if m.Http.BasePath != "" {
			router = parentRouter.PathPrefix(m.Http.BasePath).Subrouter()
		}
		if m.Http.Setup != nil {
			hasHTTP = true
			m.Http.Setup(router)
		}
	}

	for _, subM := range m.SubModules {
		hasHTTP = hasHTTP || subM.setupRouter(router)
	}

	return hasHTTP
}
