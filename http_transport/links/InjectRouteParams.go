package http_links

import (
	"reflect"

	"github.com/gorilla/mux"
	"github.com/gpabois/goservice/chain"
	endpoint_flow "github.com/gpabois/goservice/endpoint/flow"
	"github.com/gpabois/goservice/flow"
	http_flow "github.com/gpabois/goservice/http_transport/flow"
	"github.com/gpabois/gostd/result"
	"github.com/gpabois/gostd/serde/decoder"
	"github.com/gpabois/gostd/serde/norm"
)

func Reflect_InjectRouteParams() chain.Link {
	return chain.ByFunc(func(flo flow.Flow, next chain.NextFunc) chain.Result {
		// Retrieve the http request
		r := http_flow.GetHttpRequest(flo)

		// Get the endpoint request
		anyReq := endpoint_flow.GetEndpointRequest[any](flo).Expect()
		req := reflect.New(reflect.TypeOf(anyReq))
		req.Elem().Set(reflect.ValueOf(anyReq))

		// Encode the route params into a normalised map (map[string]any)
		routeParams := make(map[string]any)
		for key, value := range mux.Vars(r) {
			routeParams[key] = value
		}

		// Decode the normalised map into the endpoint request
		d := norm.NewDecoder(routeParams)
		res := decoder.Reflect_DecodeInto(d, req.Interface())
		if res.HasFailed() {
			return result.Result[flow.Flow]{}.Failed(res.UnwrapError())
		}

		// Store the updated endpoint request into the flow
		flo = endpoint_flow.SetEndpointRequest(flo, req.Elem().Interface())

		return next(flo)
	}, 102)
}

// Inject route params into the endpoint request
func InjectRouteParams[T any]() chain.Link {
	return chain.ByFunc(func(flo flow.Flow, next chain.NextFunc) chain.Result {
		// Retrieve the http request
		r := http_flow.GetHttpRequest(flo)

		// Get the endpoint request
		endpointRequest := endpoint_flow.GetEndpointRequest[T](flo).Expect()

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
		flo = endpoint_flow.SetEndpointRequest(flo, endpointRequest)

		return next(flo)
	}, 102)
}
