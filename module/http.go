package module

import (
	"fmt"
	"net"
	"net/http"
	"sort"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"

	"github.com/Zenika/marcel/config"
)

// HTTP describes a module's HTTP setup.
type HTTP struct {
	BasePath string
	Setup    func(*mux.Router)
	OnListen func(net.Listener, *http.Server)
}

func (m *Module) startHTTP() (*http.Server, error) {
	log.Infof("Starting %s's HTTP server...", m.Name)

	var router = mux.NewRouter()

	var hasHTTP = m.setupRouter(router)

	if !hasHTTP {
		return nil, nil
	}

	var addr = fmt.Sprintf(":%d", config.Default().HTTP().Port())

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

	var listener, err = net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("Could not listen for %s's HTTP: %w", m.Name, err)
	}

	log.Infof("%s's HTTP server listening on %s", m.Name, listener.Addr())

	go func() {
		if err := srv.Serve(listener); err != nil {
			log.Errorf("Error while serving HTTP for %s module: %s", m.Name, err)
			// TODO maybe stop the module ?
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
