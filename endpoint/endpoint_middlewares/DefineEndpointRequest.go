package endpoint_middlewares

import (
	"github.com/gpabois/goservice/endpoint"
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/goservice/middlewares"
	"github.com/gpabois/gostd/result"
)

func DefineEndpointRequest[T any]() middlewares.FlowMiddleware {
	return middlewares.ByFunc(func(in flow.Flow) result.Result[flow.Flow] {
		var endpointRequest T
		endpoint.Flow_SetEndpointRequest(in, endpointRequest)
		return result.Success(in)
	})
}
