package http_middlewares_tests

import (
	"errors"
	"testing"

	"github.com/gorilla/mux"
	"github.com/gpabois/goservice/chain"
	endpoint_flow "github.com/gpabois/goservice/endpoint/flow"
	"github.com/gpabois/goservice/flow"
	http_flow "github.com/gpabois/goservice/http_transport/flow"
	http_links "github.com/gpabois/goservice/http_transport/links"
	"github.com/stretchr/testify/assert"
)

func Test_InjectParams(t *testing.T) {
	expectedValue := endpointRequest{Id: 10, Value: true}
	r := NewHttpRequestFixtureWithBody(expectedValue)
	r = mux.SetURLVars(r, map[string]string{"id": "10"})

	flo := flow.Flow{}
	flo = http_flow.SetHttpRequest(flo, r)
	flo = endpoint_flow.SetEndpointRequest(flo, expectedValue)

	ch := chain.NewChain().Link(http_links.Reflect_InjectRouteParams())
	res := ch.Call(flo)
	assert.True(t, res.IsSuccess(), res.UnwrapError())

	endpointRequestRes := endpoint_flow.GetEndpointRequest[endpointRequest](flo).IntoResult(errors.New("missing endpoint request"))

	assert.True(t, endpointRequestRes.IsSuccess(), endpointRequestRes.UnwrapError())
	value := endpointRequestRes.Expect()
	assert.Equal(t, expectedValue, value)
}
