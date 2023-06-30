package http_middlewares

import (
	"github.com/gorilla/mux"
	"github.com/gpabois/goservice/endpoint"
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/goservice/http_transport"
	"github.com/gpabois/goservice/middlewares"
	"github.com/gpabois/gostd/result"
	"github.com/gpabois/gostd/serde/decoder"
	"github.com/gpabois/gostd/serde/norm"
)

// Inject route params into the endpoint request
func InjectRouteParams[T any]() middlewares.FlowMiddleware {
	return middlewares.ByFunc(func(in flow.Flow) result.Result[flow.Flow] {
		r := http_transport.Flow_GetHttpRequest(in)

		endpointRequest := endpoint.Flow_GetEndpointRequest[T](in).Expect()
		var routeParams map[string]any
		for key, value := range mux.Vars(r) {
			routeParams[key] = value
		}
		d := norm.NewDecoder(routeParams)
		res := decoder.DecodeInto(d, &endpointRequest)
		if res.HasFailed() {
			return result.Result[flow.Flow]{}.Failed(res.UnwrapError())
		}
		endpoint.Flow_SetEndpointRequest(in, res.Expect())
		return result.Success(in)
	})
}
