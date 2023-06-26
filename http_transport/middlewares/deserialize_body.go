package middlewares

import (
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/goservice/http_transport"
	"github.com/gpabois/goservice/middlewares"
	"github.com/gpabois/gostd/result"
	"github.com/gpabois/gostd/serde"
)

// Deserialize the http body based on the Content-Type header
// Store the result in "Http.Body.Deserialized"
func DeserializeHttpBody[T any]() middlewares.FlowMiddleware {
	return middlewares.ByFunc(func(in flow.Flow) result.Result[flow.Flow] {
		httpRequest := http_transport.Flow_GetHttpRequest(in)
		contentType := httpRequest.Header.Get("Content-Type")
		decodedRes := serde.DeserializeFromReader[T](httpRequest.Body, contentType)
		if decodedRes.HasFailed() {
			return result.Result[flow.Flow]{}.Failed(decodedRes.UnwrapError())
		}
		http_transport.Flow_SetDeserializedBody(in, decodedRes.Expect())
		return result.Success(in)
	})
}
