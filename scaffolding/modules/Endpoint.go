package modules_scaffolding

import (
	"github.com/gpabois/goservice/endpoint"
	endpoint_modules "github.com/gpabois/goservice/endpoint/modules"
)

func Endpoint[Request any, Response any](e endpoint.Endpoint[Request, Response]) endpoint_modules.Endpoint[Request, Response] {
	return endpoint_modules.NewEndpointModule(e)
}
