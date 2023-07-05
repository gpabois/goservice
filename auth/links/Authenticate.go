package auth_links

import (
	auth_flow "github.com/gpabois/goservice/auth/flow"
	auth_services "github.com/gpabois/goservice/auth/services"
	"github.com/gpabois/goservice/chain"
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/result"
)

type AuthenticateArgs struct {
	Service auth_services.IAuthenticationService
	Name    option.Option[string]
}

// Perform the authentication based on the determined strategy
// Store the product of the process for further processing (get a subject...)
func Authenticate(args AuthenticateArgs) chain.Link {
	name := args.Name.UnwrapOr(func() string { return "0" })
	return chain.ByFunc(func(flo flow.Flow, next chain.NextFunc) chain.Result {
		strategyOpt := auth_flow.GetAuthenticationStrategy(flo, name)

		// Nothing, so we do nothing
		if strategyOpt.IsNone() {
			return next(flo)
		}

		strategy := strategyOpt.Expect()
		authRes := args.Service.Authenticate(strategy)
		if authRes.HasFailed() {
			return result.Result[flow.Flow]{}.Failed(authRes.UnwrapError())
		}

		product := authRes.Expect()
		flo = auth_flow.SetProduct(flo, product, name)
		return next(flo)
	}, 201)
}
