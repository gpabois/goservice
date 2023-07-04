package auth_middlewares

import (
	"github.com/gpabois/goservice/auth"
	auth_services "github.com/gpabois/goservice/auth/services"
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/goservice/middlewares"
	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/result"
)

type AuthenticateArgs[Subject any] struct {
	Service auth_services.IAuthenticationService[Subject]
	Name    option.Option[string]
}

// Authenticate
func Authenticate[Subject any](args AuthenticateArgs[Subject]) middlewares.Middleware {
	name := args.Name.UnwrapOr(func() string { return "0" })
	return middlewares.ByFunc(func(in flow.Flow) result.Result[flow.Flow] {
		strategyOpt := auth.Flow_GetAuthenticationStrategy(in, name)

		// Nothing, so we do nothing
		if strategyOpt.IsNone() {
			return result.Success(in)
		}

		strategy := strategyOpt.Expect()
		authRes := args.Service.Authenticate(strategy)
		if authRes.HasFailed() {
			return result.Result[flow.Flow]{}.Failed(authRes.UnwrapError())
		}

		subject := authRes.Expect()
		in = auth.Flow_SetSubject(in, subject, name)

		return result.Success(in)
	}, 101)
}
