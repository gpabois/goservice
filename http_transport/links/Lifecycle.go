package http_links

import (
	"errors"

	"github.com/gpabois/goservice/chain"
	endpoint_flow "github.com/gpabois/goservice/endpoint/flow"
	"github.com/gpabois/goservice/flow"
	http_flow "github.com/gpabois/goservice/http_transport/flow"
	http_helpers "github.com/gpabois/goservice/http_transport/helpers"
)

// Handle the lifecycle of the http request/response
// Order: 0
func Lifecycle() chain.Link {
	return chain.ByFunc(func(flo flow.Flow, next chain.NextFunc) chain.Result {
		res := next(flo)

		w := http_flow.GetHttpResponseWriter(flo)
		r := http_flow.GetHttpRequest(flo)

		// It has failed, we write the error
		if res.HasFailed() {
			http_helpers.WriteResult(res.ToAny(), w, r)
			return res
		}

		// We write the endpoint response
		respRes := endpoint_flow.GetEndpointResponse[any](flo).IntoResult(errors.New("missing endpoint response"))
		http_helpers.WriteResult(respRes, w, r)
		if respRes.HasFailed() {
			return chain.Result{}.Failed(res.UnwrapError())
		}

		return res
	}, 1)
}
