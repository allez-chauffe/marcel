package module

// Context contains information about the execution context.
type Context interface{}

type ctx struct{}

var _ Context = new(ctx)
