package http_transport

import (
	"net/http"

	"github.com/gpabois/goservice/chain"
	"github.com/gpabois/goservice/flow"
	http_flow "github.com/gpabois/goservice/http_transport/flow"
	http_helpers "github.com/gpabois/goservice/http_transport/helpers"
	http_modules "github.com/gpabois/goservice/http_transport/modules"
)

// Basic handler
// Intercept incoming request
// Generate the endpoint request
// Call the endpoint function
// Get the endpoint response
// Intercept outcoming response
type Handler struct {
	chain chain.Chain
}

func NewHandler(ch chain.Chain, args http_modules.HttpModuleArgs) http.Handler {
	// Install http module
	ch = ch.Install(http_modules.NewHttpModule(args))
	return &Handler{chain: ch}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f := flow.Flow{}

	// Setup the http request/response writer
	f = http_flow.SetHttpRequest(f, r)
	f = http_flow.SetHttpResponseWriter(f, w)

	// Call the chain
	res := h.chain.Call(f)

	// Write the error directly, if the chaining did not catch the error
	if res.HasFailed() {
		http_helpers.WriteResult(res.ToAny(), w, r)
	}
}
