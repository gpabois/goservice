package endpoint_flow

import (
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/gostd/option"
)

func SetEndpointRequest(in flow.Flow, endpointRequest any) flow.Flow {
	in["Endpoint.Request"] = endpointRequest
	return in
}

func GetEndpointRequest[T any](in flow.Flow) option.Option[T] {
	return flow.Lookup[T]("Endpoint.Request", in)
}

func SetEndpointResponse(in flow.Flow, endpointRequest any) flow.Flow {
	in["Endpoint.Response"] = endpointRequest
	return in
}

func GetEndpointResponse[T any](in flow.Flow) option.Option[T] {
	return flow.Lookup[T]("Endpoint.Response", in)
}
