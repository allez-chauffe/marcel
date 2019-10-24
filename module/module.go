package module

import (
	"github.com/gorilla/mux"
)

type Module struct {
	Name       string
	Init       func() error
	Router     *Router
	SubModules []*Module
}

type Router struct {
	BasePath  string
	Configure func(*mux.Router) error
}

func (m *Module) ListenAndServe(host string, uint port) error {

}
