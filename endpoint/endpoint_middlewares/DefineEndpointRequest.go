package endpoint_middlewares

import (
	endpoint_flow "github.com/gpabois/goservice/endpoint/flow"
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/goservice/middlewares"
	"github.com/gpabois/gostd/result"
)

func DefineEndpointRequest[T any]() middlewares.Middleware {
	return middlewares.ByFunc(func(in flow.Flow) result.Result[flow.Flow] {
		var endpointRequest T
		endpoint_flow.Flow_SetEndpointRequest(in, endpointRequest)
		return result.Success(in)
	}, middlewares.FirstOrder)
}
