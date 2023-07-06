package http_middlewares_tests

import (
	"errors"
	"fmt"
	"testing"

	"github.com/gpabois/goservice/auth"
	auth_flow "github.com/gpabois/goservice/auth/flow"
	"github.com/gpabois/goservice/chain"
	"github.com/gpabois/goservice/flow"
	http_flow "github.com/gpabois/goservice/http/flow"
	http_links "github.com/gpabois/goservice/http/links"
	"github.com/gpabois/gostd/option"
	"github.com/stretchr/testify/assert"
)

func Test_GetAuthenticationStrategy_Bearer(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	expectedValue := auth.AuthenticationStrategy{
		Scheme:      "Bearer",
		Credentials: token,
	}

	r := NewHttpRequestFixture()
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	flo := flow.Flow{}
	flo = http_flow.SetHttpRequest(flo, r)

	ch := chain.NewChain().Link(http_links.GetAuthenticationStrategy(
		http_links.GetAuthenticationStrategyArgs{
			Name: option.Some("0"),
			Strategy: http_links.GetAuthenticationStrategyByHeader(http_links.GetAuthenticationStrategyByHeaderArgs{
				Header: option.Some("Authorization"),
			}),
		},
	))

	res := ch.Call(flo)
	assert.True(t, res.IsSuccess(), res.UnwrapError())

	valueRes := auth_flow.GetAuthenticationStrategy(flo, "0").IntoResult(errors.New("missing authentication strategy"))

	assert.True(t, valueRes.IsSuccess(), valueRes.UnwrapError())
	assert.Equal(t, expectedValue, valueRes.Expect())
}
