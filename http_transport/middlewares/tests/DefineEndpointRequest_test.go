package middlewares_tests

import (
	"errors"
	"testing"

	"github.com/gpabois/goservice/endpoint"
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/goservice/http_transport/middlewares"
	"github.com/stretchr/testify/assert"
)

func Test_DefineEndpointRequest(t *testing.T) {
	in := flow.Flow{}
	m := middlewares.DefineEndpointRequest[endpointRequest]()

	// Call the middleware
	inRes := m.Intercept(in)
	assert.True(t, inRes.IsSuccess(), inRes.UnwrapError())

	in = inRes.Expect()
	endpointRequestRes := endpoint.Flow_GetEndpointRequest[endpointRequest](in).IntoResult(errors.New("missing endpoint request"))

	assert.True(t, endpointRequestRes.IsSuccess(), endpointRequestRes.UnwrapError())
}
