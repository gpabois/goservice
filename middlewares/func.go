package middlewares

import (
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/gostd/result"
)

type middlewareFunc struct {
	inner func(flow flow.Flow) result.Result[flow.Flow]
	order int
}

func ByFunc(fn func(flow flow.Flow) result.Result[flow.Flow], order int) Middleware {
	return middlewareFunc{inner: fn, order: order}
}

func (m middlewareFunc) Intercept(flow flow.Flow) result.Result[flow.Flow] {
	return m.inner(flow)
}

func (m middlewareFunc) Order() int {
	return m.order
}
