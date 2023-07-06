package endpoint_scaffolding

import (
	"github.com/gpabois/goservice/endpoint"
	endpoint_modules "github.com/gpabois/goservice/endpoint/modules"
)

func ScaffoldEndpoint[Request any, Response any](e endpoint.Endpoint[Request, Response]) endpoint_modules.EndpointModule[Request, Response] {
	return endpoint_modules.NewEndpointModule(e)
}
