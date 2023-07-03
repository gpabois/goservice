package http_middlewares_tests

import (
	"errors"
	"fmt"
	"testing"

	"github.com/gpabois/goservice/auth"
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/goservice/http_transport"
	"github.com/gpabois/goservice/http_transport/http_middlewares"
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

	m := http_middlewares.GetAuthenticationStrategy(http_middlewares.GetAuthenticationStrategyArgs{})

	in := flow.Flow{}
	in = http_transport.Flow_SetHttpRequest(in, r)

	inRes := m.Intercept(in)
	assert.True(t, inRes.IsSuccess(), inRes.UnwrapError())

	valueRes := auth.Flow_GetAuthenticationStrategy(in, "0").IntoResult(errors.New("missing authentication strategy"))
	assert.True(t, valueRes.IsSuccess(), inRes.UnwrapError())

	assert.Equal(t, expectedValue, valueRes.Expect())
}
