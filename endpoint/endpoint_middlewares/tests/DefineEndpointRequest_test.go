package endpoint_middlewares_tests

import (
	"errors"
	"testing"

	"github.com/gpabois/goservice/endpoint"
	"github.com/gpabois/goservice/endpoint/endpoint_middlewares"
	"github.com/gpabois/goservice/flow"
	"github.com/stretchr/testify/assert"
)

func Test_DefineEndpointRequest(t *testing.T) {
	in := flow.Flow{}
	m := endpoint_middlewares.DefineEndpointRequest[endpointRequest]()

	// Call the middleware
	inRes := m.Intercept(in)
	assert.True(t, inRes.IsSuccess(), inRes.UnwrapError())

	in = inRes.Expect()
	endpointRequestRes := endpoint.Flow_GetEndpointRequest[endpointRequest](in).IntoResult(errors.New("missing endpoint request"))

	assert.True(t, endpointRequestRes.IsSuccess(), endpointRequestRes.UnwrapError())
}
