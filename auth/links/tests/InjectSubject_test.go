package auth_links_tests

import (
	"errors"
	"testing"

	auth_flow "github.com/gpabois/goservice/auth/flow"
	auth_links "github.com/gpabois/goservice/auth/links"
	"github.com/gpabois/goservice/chain"
	endpoint_flow "github.com/gpabois/goservice/endpoint/flow"
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/gostd/option"
	"github.com/stretchr/testify/assert"
)

func Test_Reflect_InjectSubject(t *testing.T) {
	value := endpointRequest{Value: true}
	expectedValue := endpointRequest{
		Value:   true,
		Subject: option.Some(subject{Id: 10}),
	}

	flo := flow.Flow{}
	flo = endpoint_flow.SetEndpointRequest(flo, value)
	flo = auth_flow.SetSubject(flo, subject{Id: 10}, "0")

	ch := chain.NewChain().
		Link(auth_links.Reflect_InjectSubject(auth_links.InjectSubjectArgs{
			FieldName: "subject",
			Name:      option.Some("0"),
		}))

	res := ch.Call(flo)
	assert.True(t, res.IsSuccess(), res.UnwrapError())

	endpointRequestRes := endpoint_flow.GetEndpointRequest[endpointRequest](flo).IntoResult(errors.New("missing endpoint request"))

	assert.True(t, endpointRequestRes.IsSuccess(), endpointRequestRes.UnwrapError())
	value = endpointRequestRes.Expect()
	assert.Equal(t, expectedValue, value)
}
