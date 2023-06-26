package http_transport

import (
	"net/http"

	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/gostd/option"
)

func Flow_GetHttpRequest(in flow.Flow) *http.Request {
	return (flow.Lookup[*http.Request]("Http.Request", in).Expect())
}

func Flow_SetHttpRequest(in flow.Flow, req *http.Request) flow.Flow {
	in["Http.Request"] = req
	return in
}

func Flow_SetDeserializedBody(in flow.Flow, decoded any) flow.Flow {
	in["Http.Body.Deserialized"] = decoded
	return in
}

func Flow_GetDeserializedBody[T any](in flow.Flow) option.Option[T] {
	return flow.Lookup[T]("Http.Body.Deserialized", in)
}
