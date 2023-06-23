package middlewares

import "github.com/gpabois/gostd/result"

type middlewareFunc[Input any, Output any] struct {
	inner func(in Input) result.Result[Output]
}

func ByFunc[Input any, Output any](fn func(in Input) result.Result[Output]) Middleware[Input, Output] {
	return middlewareFunc[Input, Output]{inner: fn}
}

func (m middlewareFunc[Input, Output]) Intercept(in Input) result.Result[Output] {
	return m.inner(in)
}

func (m middlewareFunc[Input, Output]) Connect(right Middleware[Output, Output]) Middleware[Input, Output] {
	return Connect[Input, Output](m, right)
}
