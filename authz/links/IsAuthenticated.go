package authz_links

import (
	auth_flow "github.com/gpabois/goservice/auth/flow"
	"github.com/gpabois/goservice/authz"
	"github.com/gpabois/goservice/chain"
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/gostd/option"
)

type AuthenticatedArgs struct {
	Name option.Option[string]
}

// Ensure that the user is authenticated
// Impl: Check if authentication's product exists
func IsAuthenticated(args AuthenticatedArgs) chain.Link {
	name := args.Name.UnwrapOr(func() string { return "0" })
	return chain.ByFunc(func(flo flow.Flow, next chain.NextFunc) chain.Result {
		productOpt := auth_flow.GetProduct(flo, name)
		if productOpt.IsNone() {
			return chain.Result{}.Failed(authz.NewNotAuthenticatedError())
		}
		return next(flo)
	}, 302)
}
