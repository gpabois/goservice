package endpoint_modules

import (
	"github.com/gpabois/goservice/chain"
	"github.com/gpabois/goservice/endpoint"
	endpoint_links "github.com/gpabois/goservice/endpoint/links"
)

type EndpointModule[Request any, Response any] struct {
	Endpoint endpoint.Endpoint[Request, Response]
}

func NewEndpointModule[Request any, Response any](e endpoint.Endpoint[Request, Response]) chain.Module {
	return EndpointModule[Request, Response]{Endpoint: e}
}

func (mod EndpointModule[Request, Response]) Install(ch chain.Chain) chain.Chain {
	return ch.Link(endpoint_links.SetEndpointRequest[Request]()).Link(endpoint_links.CallEndpoint(mod.Endpoint))
}
