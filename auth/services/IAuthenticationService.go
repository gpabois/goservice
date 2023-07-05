package auth_services

import (
	"github.com/gpabois/goservice/auth"
	"github.com/gpabois/gostd/result"
)

//go:generate mockery
type IAuthenticationService interface {
	// Authenticate only, returns a product of the authentication (claims if JWT for instance, ...)
	Authenticate(strategy auth.AuthenticationStrategy) result.Result[any]
}
