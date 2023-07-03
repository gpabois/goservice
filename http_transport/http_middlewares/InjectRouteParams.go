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
		// Retrieve the http request
		r := http_transport.Flow_GetHttpRequest(in)

		// Get the endpoint request
		endpointRequest := endpoint.Flow_GetEndpointRequest[T](in).Expect()

		// Encode the route params into a normalised map (map[string]any)
		var routeParams map[string]any
		for key, value := range mux.Vars(r) {
			routeParams[key] = value
		}

		// Decode the normalised map into the endpoint request
		d := norm.NewDecoder(routeParams)
		res := decoder.DecodeInto(d, &endpointRequest)
		if res.HasFailed() {
			return result.Result[flow.Flow]{}.Failed(res.UnwrapError())
		}

		// Store the updated endpoint request into the flow
		endpoint.Flow_SetEndpointRequest(in, res.Expect())

		return result.Success(in)
	})
}
