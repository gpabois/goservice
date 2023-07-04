package endpoint_flow

import (
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/result"
)

func Flow_SetEndpointRequest(in flow.Flow, endpointRequest any) flow.Flow {
	in["Endpoint.Request"] = endpointRequest
	return in
}

func Flow_GetEndpointRequest[T any](in flow.Flow) option.Option[T] {
	return flow.Lookup[T]("Endpoint.Request", in)
}

func Flow_SetEndpointResult(in flow.Flow, endpointResponse any) flow.Flow {
	in["Endpoint.Result"] = endpointResponse
	return in
}

func Flow_GetEndpointResult[T any](in flow.Flow) option.Option[result.Result[T]] {
	return flow.Lookup[result.Result[T]]("Endpoint.Result", in)
}
