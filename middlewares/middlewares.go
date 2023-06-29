package middlewares

import (
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/gostd/iter"
	"github.com/gpabois/gostd/result"
)

type Middleware[Input any, Output any] interface {
	Intercept(in Input) result.Result[Output]
	// Only works for iso-type connection (such as Middleware[B, B])
	Connect(right Middleware[Output, Output]) Middleware[Input, Output]
}

type FlowMiddleware interface {
	Middleware[flow.Flow, flow.Flow]
}

func Chain[T any](m1 Middleware[T, T], others ...Middleware[T, T]) Middleware[T, T] {
	return iter.Reduce(iter.IterSlice(&others), Connect[T, T, T], m1)
}

func Connect[Input any, Bridge any, Output any](m1 Middleware[Input, Bridge], m2 Middleware[Bridge, Output]) Middleware[Input, Output] {
	inner := func(in Input) result.Result[Output] {
		res := m1.Intercept(in)
		if res.HasFailed() {
			return result.Result[Output]{}.Failed(res.UnwrapError())
		}
		return m2.Intercept(res.Expect())
	}

	return middlewareFunc[Input, Output]{inner}
}
