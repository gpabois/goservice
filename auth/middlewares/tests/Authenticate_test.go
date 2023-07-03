package auth_middlewares_tests

import (
	"testing"

	"github.com/gpabois/goservice/auth"
	auth_middlewares "github.com/gpabois/goservice/auth/middlewares"
	"github.com/gpabois/goservice/auth/services/mocks"
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/gostd/result"
	"github.com/stretchr/testify/assert"
)

type subject struct {
	Id string
}

func Test_Authenticate_Success(t *testing.T) {
	expectedValue := subject{Id: "me"}
	s := mocks.NewIAuthenticationService[subject](t)
	strategy := auth.NewBearer("myToken")

	s.EXPECT().Authenticate(strategy).Return(result.Success(expectedValue))

	in := flow.Flow{}
	in = auth.Flow_SetAuthenticationStrategy(in, strategy, "0")

	m := auth_middlewares.Authenticate(auth_middlewares.AuthenticateArgs[subject]{Service: s})
	res := m.Intercept(in)

	assert.True(t, res.IsSuccess(), res.UnwrapError())
	value := auth.Flow_GetSubject[subject](in, "0").Expect()
	assert.Equal(t, expectedValue, value)
}
