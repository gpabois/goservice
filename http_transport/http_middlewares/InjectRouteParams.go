package http_middlewares

import (
	"github.com/gorilla/mux"
	endpoint_flow "github.com/gpabois/goservice/endpoint/flow"
	"github.com/gpabois/goservice/flow"
	http_flow "github.com/gpabois/goservice/http_transport/flow"
	"github.com/gpabois/goservice/middlewares"
	"github.com/gpabois/gostd/result"
	"github.com/gpabois/gostd/serde/decoder"
	"github.com/gpabois/gostd/serde/norm"
)

// Inject route params into the endpoint request
func InjectRouteParams[T any]() middlewares.Middleware {
	return middlewares.ByFunc(func(in flow.Flow) result.Result[flow.Flow] {
		// Retrieve the http request
		r := http_flow.Flow_GetHttpRequest(in)

		// Get the endpoint request
		endpointRequest := endpoint_flow.Flow_GetEndpointRequest[T](in).Expect()

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
		endpoint_flow.Flow_SetEndpointRequest(in, res.Expect())

		return result.Success(in)
	}, 102)
}
