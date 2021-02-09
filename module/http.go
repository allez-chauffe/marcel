package module

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"

	"github.com/allez-chauffe/marcel/config"
	"github.com/allez-chauffe/marcel/httputil"
)

// HTTP describes a module's HTTP setup.
type HTTP struct {
	BasePath      string // FIXME replace by func receiving ctx?
	RedirectSlash bool
	Setup         func(ctx Context, basePath string, r *mux.Router)
	OnListen      func(ctx Context, l net.Listener, srv *http.Server)
}

func (m *Module) startHTTP(ctx *ctx) (*http.Server, error) {
	log.Infof("Starting %s's HTTP server...", m.Name)

	var rootRouter = mux.NewRouter()

	var basePath, router = mountSubrouter("", rootRouter, httputil.NormalizeBase(config.Default().HTTP().BasePath()), false)

	m.normalizeBasePaths()

	var hasHTTP = m.setupRouter(ctx, basePath, router)

	if !hasHTTP {
		return nil, nil
	}

	var addr = fmt.Sprintf(":%d", config.Default().HTTP().Port())

	var handler http.Handler = rootRouter

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

	var l, err = net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("Could not listen for %s's HTTP: %w", m.Name, err)
	}

	log.Infof("%s's HTTP server listening on %s", m.Name, l.Addr())

	go func() {
		if err := srv.Serve(l); err != nil {
			if err != http.ErrServerClosed {
				log.Errorf("Error while serving HTTP for %s module: %s", m.Name, err)
				// TODO maybe stop the module ?
			}
		}
	}()

	m.notifyOnListen(ctx, l, srv)

	return srv, nil
}

func (m *Module) stopHTTP(srv *http.Server) {
	log.Infof("Shutting down %s's HTTP server...", m.Name)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	// FIXME manage websockets
	if err := srv.Shutdown(ctx); err != nil {
		log.Errorf("Error while shutting down %s's HTTP: %s", m.Name, err)
	}

	log.Infof("%s's HTTP server stopped", m.Name)
}

func (m *Module) normalizeBasePaths() {
	m.BasePath = httputil.NormalizeBase(m.BasePath)

	for _, subM := range m.SubModules {
		subM.normalizeBasePaths()
	}
}

func (m *Module) mountSubrouter(parentBasePath string, parentRouter *mux.Router) (string, *mux.Router) {
	var basePath, r = mountSubrouter(parentBasePath, parentRouter, m.BasePath, m.RedirectSlash)

	if basePath != parentBasePath {
		uris[m.Name] = basePath

		log.Debugf("Mounted subrouter for %s at %s", m.Name, basePath)
	}

	return basePath, r
}

func mountSubrouter(parentBasePath string, parentRouter *mux.Router, basePath string, redirectSlash bool) (string, *mux.Router) {
	if basePath == "" {
		return parentBasePath, parentRouter
	}

	var r = parentRouter.PathPrefix(httputil.TrimTrailingSlash(basePath)).Subrouter()

	var absoluteBasePath = httputil.NormalizeBase(path.Join(parentBasePath, basePath))
	if redirectSlash {
		r.Handle("", http.RedirectHandler(absoluteBasePath, http.StatusMovedPermanently))
	}

	return absoluteBasePath, r
}

func (m *Module) setupRouter(ctx *ctx, parentBasePath string, parentRouter *mux.Router) bool {
	var hasHTTP = false

	var basePath, router = m.mountSubrouter(parentBasePath, parentRouter)

	if m.Setup != nil {
		hasHTTP = true
		router.HandleFunc("/uris", urisHandler)
		m.Setup(ctx, basePath, router)
		log.Debugf("Configured subrouter for %s", m.Name)
	}

	sort.Sort(byBasePath(m.SubModules))

	for _, subM := range m.SubModules {
		hasHTTP = subM.setupRouter(ctx, basePath, router) || hasHTTP
	}

	return hasHTTP
}

func (m *Module) notifyOnListen(ctx *ctx, l net.Listener, srv *http.Server) {
	if m.OnListen != nil {
		m.OnListen(ctx, l, srv)
	}

	for _, subM := range m.SubModules {
		subM.notifyOnListen(ctx, l, srv)
	}
}

type byBasePath []*Module

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
