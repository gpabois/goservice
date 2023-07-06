package http_middlewares_tests

import (
	"errors"
	"testing"

	"github.com/gpabois/goservice/chain"
	endpoint_flow "github.com/gpabois/goservice/endpoint/flow"
	"github.com/gpabois/goservice/flow"
	http_flow "github.com/gpabois/goservice/http/flow"
	http_links "github.com/gpabois/goservice/http/links"
	"github.com/stretchr/testify/assert"
)

func Test_DeserializeBody(t *testing.T) {
	expectedValue := endpointRequest{Value: true}
	r := NewHttpRequestFixtureWithBody(expectedValue)

	flo := flow.Flow{}
	flo = http_flow.SetHttpRequest(flo, r)
	flo = endpoint_flow.SetEndpointRequest(flo, expectedValue)

	ch := chain.NewChain().Link(http_links.Reflect_DeserializedBody(http_links.ReflectDeserializeBodyArgs{}))
	res := ch.Call(flo)
	assert.True(t, res.IsSuccess(), res.UnwrapError())

	endpointRequestRes := endpoint_flow.GetEndpointRequest[endpointRequest](flo).IntoResult(errors.New("missing endpoint request"))

	assert.True(t, endpointRequestRes.IsSuccess(), endpointRequestRes.UnwrapError())
	value := endpointRequestRes.Expect()
	assert.Equal(t, expectedValue, value)
}
