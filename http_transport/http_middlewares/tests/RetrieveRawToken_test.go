package http_middlewares_tests

import (
	"errors"
	"fmt"
	"testing"

	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/goservice/http_transport"
	"github.com/gpabois/goservice/http_transport/http_middlewares"
	"github.com/gpabois/goservice/token"
	"github.com/stretchr/testify/assert"
)

func Test_RetrieveRawToken_Bearer(t *testing.T) {
	expectedValue := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	r := NewHttpRequestFixture()
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", expectedValue))

	m := http_middlewares.RetrieveRawToken(http_middlewares.TokenRetrieveArgs{})

	in := flow.Flow{}
	in = http_transport.Flow_SetHttpRequest(in, r)

	inRes := m.Intercept(in)
	assert.True(t, inRes.IsSuccess(), inRes.UnwrapError())

	valueRes := token.Flow_GetRawToken(in, "0").IntoResult(errors.New("missing raw token"))
	assert.True(t, valueRes.IsSuccess(), inRes.UnwrapError())

	assert.Equal(t, expectedValue, valueRes.Expect())
}
