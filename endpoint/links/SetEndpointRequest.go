package endpoint_links

import (
	"github.com/gpabois/goservice/chain"
	endpoint_flow "github.com/gpabois/goservice/endpoint/flow"
	"github.com/gpabois/goservice/flow"
)

// Call the endpoint
// Order 1000
func SetEndpointRequest[Request any]() chain.Link {
	return chain.ByFunc(func(flo flow.Flow, next chain.NextFunc) chain.Result {
		var req Request
		flo = endpoint_flow.SetEndpointRequest(flo, req)
		return next(flo)
	}, 0)
}
