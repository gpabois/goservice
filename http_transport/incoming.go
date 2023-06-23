package http_transport

import (
	"net/http"

	"github.com/gpabois/goservice/bridge"
	"github.com/gpabois/gostd/option"
)

type Incoming bridge.Bridge

func (in Incoming) SetEndpointRequest(endpointRequest any) {
	in["endpoint_request"] = endpointRequest
}

func (in Incoming) GetEndpointRequest() option.Option[*any] {
	return bridge.Lookup[any]("endpoint_request", in)
}

func (in Incoming) GetHttpRequest() *http.Request {
	return *(bridge.Lookup[*http.Request]("endpoint_request", in).Expect())
}

func (in Incoming) SetDeserializedBody(decoded any) option.Option[any] {
	in["deserialized_body"] = decoded
}

func (in Incoming) GetDeserializedBody() option.Option[*any] {
	return bridge.Lookup[any]("deserialized_body", in)
}
