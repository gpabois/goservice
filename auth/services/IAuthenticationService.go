package auth_services

import (
	"github.com/gpabois/goservice/auth"
	"github.com/gpabois/gostd/result"
)

//go:generate mockery
type IAuthenticationService[Subject any] interface {
	Authenticate(strategy auth.AuthenticationStrategy) result.Result[Subject]
}
