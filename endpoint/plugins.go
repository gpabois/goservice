package endpoint

import (
	"github.com/gpabois/goservice/endpoint/endpoint_middlewares"
	"github.com/gpabois/goservice/middlewares"
)

type EndpointPlugin[Request any, Response any] struct{}

func (p EndpointPlugin[Request, Response]) Install(io *middlewares.IO) {
	io.Incoming = io.Incoming.Add(endpoint_middlewares.DefineEndpointRequest[Request]())
}
