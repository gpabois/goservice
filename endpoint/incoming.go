package endpoint

import (
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/gostd/option"
)

func Flow_SetEndpointRequest(in flow.Flow, endpointRequest any) {
	in["Endpoint.Request"] = endpointRequest
}

func Flow_GetEndpointRequest[T any](in flow.Flow) option.Option[T] {
	return flow.Lookup[T]("Endpoint.Request", in)
}
