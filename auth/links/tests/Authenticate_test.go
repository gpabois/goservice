package auth_links_tests

import (
	"testing"

	"github.com/gpabois/goservice/auth"
	auth_flow "github.com/gpabois/goservice/auth/flow"
	auth_links "github.com/gpabois/goservice/auth/links"
	"github.com/gpabois/goservice/auth/services/mocks"
	"github.com/gpabois/goservice/chain"
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/gostd/result"
	"github.com/stretchr/testify/assert"
)

func Test_Authenticate_Success(t *testing.T) {
	expectedValue := subject{Id: 10}
	s := mocks.NewIAuthenticationService(t)
	strategy := auth.NewBearer("myToken")

	s.EXPECT().Authenticate(strategy).Return(result.Success[any](expectedValue))

	flo := flow.Flow{}
	flo = auth_flow.SetAuthenticationStrategy(flo, strategy, "0")

	ch := chain.NewChain().Link(auth_links.Authenticate(auth_links.AuthenticateArgs{
		Service: s,
	}))
	res := ch.Call(flo)

	assert.True(t, res.IsSuccess(), res.UnwrapError())
	valueOpt := auth_flow.GetProduct(flo, "0")

	assert.True(t, valueOpt.IsSome())
	assert.Equal(t, expectedValue, valueOpt.Expect())
}
