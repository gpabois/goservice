package http_middlewares

import (
	endpoint_flow "github.com/gpabois/goservice/endpoint/flow"
	"github.com/gpabois/goservice/flow"
	http_flow "github.com/gpabois/goservice/http_transport/flow"
	"github.com/gpabois/goservice/middlewares"
	"github.com/gpabois/gostd/result"
	"github.com/gpabois/gostd/serde"
)

// Deserialize the http body based on the Content-Type header, into the endpoint request
func DeserializeBody[T any]() middlewares.Middleware {
	return middlewares.ByFunc(func(in flow.Flow) result.Result[flow.Flow] {
		httpRequest := http_flow.Flow_GetHttpRequest(in)
		contentType := httpRequest.Header.Get("Content-Type")

		endpointRequest := endpoint_flow.Flow_GetEndpointRequest[T](in).Expect()
		decodedRes := serde.DeserializeFromReaderInto(httpRequest.Body, contentType, &endpointRequest)

		if decodedRes.HasFailed() {
			return result.Result[flow.Flow]{}.Failed(decodedRes.UnwrapError())
		}

		endpoint_flow.Flow_SetEndpointRequest(in, endpointRequest)

		return result.Success(in)
	}, 102)
}
