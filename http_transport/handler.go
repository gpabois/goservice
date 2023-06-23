package http_transport

import (
	"net/http"

	"github.com/gpabois/goservice/endpoint"
	"github.com/gpabois/goservice/middlewares"
	"github.com/gpabois/gostd/result"
	"github.com/gpabois/gostd/serde"
)

// Basic handler
// Intercept incoming request
// Generate the endpoint request
// Call the endpoint function
// Get the endpoint response
// Intercept outcoming response
type Handler[EndpointRequest any, EndpointResponse any] struct {
	endpoint endpoint.Endpoint[EndpointRequest, EndpointResponse]
	incoming middlewares.IsoMiddleware[Incoming]
}

// Deserialize the http body based on the Content-Type header
// Store the result in "deserialized_body"
func DeserializeHttpBody[T any]() middlewares.Middleware[Incoming] {
	return middlewares.ByFunc(func(in Incoming) result.Result[Incoming] {
		httpRequest := in.GetHttpRequest()
		contentType := httpRequest.Header.Get("Content-Type")
		decodedRes := serde.DeserializeFromReader[T](httpRequest.Body)
		if decodedRes.HasFailed() {
			return result.Result[Incoming]{}.Failed(decodedRes.UnwrapError())
		}
		in.SetDeserializedBody(decodedRes.Expect())
		return result.Success(in)
	})
}

// Take the value stored in "deserialized_body" and set it as the endpoint request
func DeserializedBodyIsEndpointRequest() middlewares.Middleware[Incoming] {
	return middlewares.ByFunc(func(in Incoming) result.Result[Incoming] {
		endpointRequest := *in.GetDeserializedBody().Expect()

	})
}

func (h *SimpleHandler[Request, Response]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	interceptIncoming := h.incoming.Intercept(Incoming{
		"http_request": r,
	})

}
