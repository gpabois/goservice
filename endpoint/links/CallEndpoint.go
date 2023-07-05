package endpoint_links

import (
	"context"
	"errors"

	"github.com/gpabois/goservice/chain"
	"github.com/gpabois/goservice/endpoint"
	endpoint_flow "github.com/gpabois/goservice/endpoint/flow"
	"github.com/gpabois/goservice/flow"
)

// Call the endpoint
// Order 1000
func CallEndpoint[Request any, Response any](e endpoint.Endpoint[Request, Response]) chain.Link {
	return chain.ByFunc(func(flo flow.Flow, next chain.NextFunc) chain.Result {
		req := endpoint_flow.GetEndpointRequest[Request](flo).IntoResult(errors.New("missing endpoint request"))

		if req.HasFailed() {
			return chain.Result{}.Failed(req.UnwrapError())
		}

		resp := e.Process(context.Background(), req.Expect())

		if resp.HasFailed() {
			return chain.Result{}.Failed(resp.UnwrapError())
		}

		// Save the response into the flow
		flo = endpoint_flow.SetEndpointResponse(flo, resp.Expect())

		return next(flo)
	}, 400)
}
