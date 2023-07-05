package authz_links

import (
	auth_flow "github.com/gpabois/goservice/auth/flow"
	"github.com/gpabois/goservice/authz"
	"github.com/gpabois/goservice/chain"
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/gostd/result"
)

type ExtractObjectArgs struct {
	Name string
	Func func(flo flow.Flow) result.Result[any]
}

// Extract the object of the ACL
func ExtractObject(args AuthenticatedArgs) chain.Link {
	name := args.Name.UnwrapOr(func() string { return "0" })
	return chain.ByFunc(func(flo flow.Flow, next chain.NextFunc) chain.Result {
		productOpt := auth_flow.GetProduct(flo, name)
		if productOpt.IsNone() {
			return chain.Result{}.Failed(authz.NewNotAuthenticatedError())
		}
		return next(flo)
	}, 301)
}
