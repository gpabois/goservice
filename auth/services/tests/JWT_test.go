package auth_services_tests

import (
	"testing"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/gpabois/goservice/auth"
	auth_services "github.com/gpabois/goservice/auth/services"
	"github.com/stretchr/testify/assert"
)

type claims struct {
	jwt.RegisteredClaims
	SubjectId int
}

type subject struct {
	Id int `serde:"SubjectId"`
}

func Test_JWT(t *testing.T) {
	key := []byte("mySigningKey")
	expectedValue := subject{Id: 10}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{SubjectId: 10}).SignedString(key)
	assert.Nil(t, err, err)

	strategy := auth.NewBearer(token)
	s := auth_services.NewJWT[subject](auth_services.JWTArgs{
		KeyFunc: func(_ *jwt.Token) (any, error) { return key, nil },
	})

	subRes := s.Authenticate(strategy)

	assert.True(t, subRes.IsSuccess(), subRes.UnwrapError())
	assert.Equal(t, expectedValue, subRes.Expect())
}
