package http_middlewares_tests

import (
	"errors"
	"testing"

	"github.com/gpabois/goservice/endpoint"
	"github.com/gpabois/goservice/endpoint/endpoint_middlewares"
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/goservice/http_transport"
	"github.com/gpabois/goservice/http_transport/http_middlewares"
	"github.com/gpabois/goservice/middlewares"
	"github.com/stretchr/testify/assert"
)

func Test_DeserializeBody(t *testing.T) {
	expectedValue := endpointRequest{Value: true}
	r := NewHttpRequestFixtureWithBody(expectedValue)

	in := flow.Flow{}
	in = http_transport.Flow_SetHttpRequest(in, r)

	m := middlewares.Chain[flow.Flow](
		endpoint_middlewares.DefineEndpointRequest[endpointRequest](),
		http_middlewares.DeserializeBody[endpointRequest](),
	)

	// Call the middleware
	inRes := m.Intercept(in)
	assert.True(t, inRes.IsSuccess(), inRes.UnwrapError())

	in = inRes.Expect()
	endpointRequestRes := endpoint.Flow_GetEndpointRequest[endpointRequest](in).IntoResult(errors.New("missing endpoint request"))

	assert.True(t, endpointRequestRes.IsSuccess(), endpointRequestRes.UnwrapError())
	value := endpointRequestRes.Expect()
	assert.Equal(t, expectedValue, value)
}
