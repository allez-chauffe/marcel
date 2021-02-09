package module

import (
	"fmt"
	"reflect"
)

// Context contains information about the execution context.
type Context interface {
	Register(v interface{})
	Lookup(p interface{}) (interface{}, bool)
}

type ctx struct {
	values map[reflect.Type]interface{}
}

var _ Context = new(ctx)

func (ctx *ctx) Register(v interface{}) {
	t := reflect.TypeOf(v)
	if _, ok := ctx.values[t]; ok {
		panic(fmt.Errorf("%s already registered in modules context", t))
	}
	ctx.values[t] = v
}

func (ctx *ctx) Lookup(p interface{}) (interface{}, bool) {
	t := reflect.TypeOf(p)
	v, ok := ctx.values[t]
	return v, ok
}
