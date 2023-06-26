package http_transport

import (
	"net/http"

	"github.com/gpabois/goservice/endpoint"
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/goservice/middlewares"
	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/serde"
)

// Basic handler
// Intercept incoming request
// Generate the endpoint request
// Call the endpoint function
// Get the endpoint response
// Intercept outcoming response
type Handler[EndpointRequest any, EndpointResponse any] struct {
	endpoint  endpoint.Endpoint[EndpointRequest, EndpointResponse]
	incoming  middlewares.FlowMiddleware
	outcoming middlewares.FlowMiddleware
}

type HttpResult[T any] struct {
	Data  option.Option[T]
	Error string
}

func (res HttpResult[T]) Failed(err error, code option.Option[int]) HttpResult[T] {
	return HttpResult[T]{}
}

// Write the error
func WriteError[Response any](err error, w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Accept")

	if contentType == "" {
		contentType = "application/json"
	}

	switch err.(type) {
	case serde.UnhandledContentType:
		w.WriteHeader(415)
		contentType = "application/json"
	case serde.DeserializeError:
		w.WriteHeader(400)
	default:
		w.WriteHeader(500)
	}

	encodedRes := serde.Serialize(HttpResult[Response]{Error: err.Error()}, contentType)
	w.Write(encodedRes.Expect())
}

func (h *Handler[Request, Response]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	in := flow.Flow{}
	in = Flow_SetHttpRequest(in, r)

	incomingResult := h.incoming.Intercept(in)

	// The interception has failed, we don't know why
	if incomingResult.HasFailed() {
		WriteError[Response](incomingResult.UnwrapError(), w, r)
		return
	} else {
		endpointRequestOpt := endpoint.Flow_GetEndpointRequest[Request](incomingResult.Expect())

	}

}
