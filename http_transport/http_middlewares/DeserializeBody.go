package http_middlewares

import (
	"github.com/gpabois/goservice/endpoint"
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/goservice/http_transport"
	"github.com/gpabois/goservice/middlewares"
	"github.com/gpabois/gostd/result"
	"github.com/gpabois/gostd/serde"
)

// Deserialize the http body based on the Content-Type header, into the endpoint request
func DeserializeBody[T any]() middlewares.FlowMiddleware {
	return middlewares.ByFunc(func(in flow.Flow) result.Result[flow.Flow] {
		httpRequest := http_transport.Flow_GetHttpRequest(in)
		contentType := httpRequest.Header.Get("Content-Type")

		endpointRequest := endpoint.Flow_GetEndpointRequest[T](in).Expect()
		decodedRes := serde.DeserializeFromReaderInto(httpRequest.Body, contentType, &endpointRequest)

		if decodedRes.HasFailed() {
			return result.Result[flow.Flow]{}.Failed(decodedRes.UnwrapError())
		}

		endpoint.Flow_SetEndpointRequest(in, endpointRequest)

		return result.Success(in)
	})
}
