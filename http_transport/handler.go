package http_transport

import (
	"context"
	"errors"
	"net/http"

	"github.com/gpabois/goservice/endpoint"
	endpoint_flow "github.com/gpabois/goservice/endpoint/flow"
	"github.com/gpabois/goservice/flow"
	http_flow "github.com/gpabois/goservice/http_transport/flow"
	"github.com/gpabois/goservice/middlewares"
)

// Basic handler
// Intercept incoming request
// Generate the endpoint request
// Call the endpoint function
// Get the endpoint response
// Intercept outcoming response
type Handler[EndpointRequest any, EndpointResponse any] struct {
	endpoint endpoint.Endpoint[EndpointRequest, EndpointResponse]
	io       middlewares.IO
}

func NewHandler[EndpointRequest any, EndpointResponse any](e endpoint.Endpoint[EndpointRequest, EndpointResponse], io middlewares.IO) http.Handler {
	// Install endpoint plugin automatically
	endpoint.EndpointPlugin[EndpointRequest, EndpointResponse]{}.Install(&io)

	return &Handler[EndpointRequest, EndpointResponse]{
		endpoint: e,
		io:       io,
	}
}

func (h *Handler[Request, Response]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f := flow.Flow{}
	// Set the http request
	f = http_flow.Flow_SetHttpRequest(f, r)

	incomingResult := h.io.Incoming.Intercept(f)
	if incomingResult.HasFailed() {
		WriteResult(incomingResult.ToAny(), w, r)
		return
	}

	endpointRequestRes := endpoint_flow.Flow_GetEndpointRequest[Request](incomingResult.Expect()).IntoResult(NewInternalServerError(errors.New("missing endpoint request")))
	if endpointRequestRes.HasFailed() {
		WriteResult(endpointRequestRes.ToAny(), w, r)
		return
	}

	endpointRequest := endpointRequestRes.Expect()
	endpointRespResult := h.endpoint.Process(context.Background(), endpointRequest)

	f = http_flow.Flow_SetHttpResponseWriter(f, w)
	f = endpoint_flow.Flow_SetEndpointResult(f, endpointRespResult)

	res := h.io.Outcoming.Intercept(f)

	// Write the error directly, if the outcoming interception failed.
	if res.HasFailed() {
		WriteResult(res.ToAny(), w, r)
	}
}
