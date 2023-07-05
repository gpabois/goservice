package authz_links

import (
	auth_flow "github.com/gpabois/goservice/auth/flow"
	"github.com/gpabois/goservice/authz"
	authz_flow "github.com/gpabois/goservice/authz/flow"
	authz_services "github.com/gpabois/goservice/authz/services"
	"github.com/gpabois/goservice/chain"
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/gostd/option"
)

type HasPermissionArgs struct {
	SubjectName option.Option[string]
	Name        option.Option[string]
	Permission  string
	Service     authz_services.IAuthorizationService
}

// Get the ACL, and check if the subject has the right permission
func HasPermission(args HasPermissionArgs) chain.Link {
	name := args.Name.UnwrapOr(func() string { return "0" })
	subjectName := args.SubjectName.UnwrapOr(func() string { return "0" })

	return chain.ByFunc(func(flo flow.Flow, next chain.NextFunc) chain.Result {
		subjectOpt := auth_flow.GetSubject[any](flo, subjectName)
		if subjectOpt.IsNone() {
			return chain.Result{}.Failed(authz.NewNotAuthenticatedError())
		}

		objectOpt := authz_flow.GetObject[any](flo, name)
		acl := authz_services.ACL{Subject: subjectOpt.Expect(), Permission: args.Permission, Object: objectOpt}

		res := args.Service.Can(acl)
		if res.HasFailed() {
			return chain.Result{}.Failed(res.UnwrapError())
		}

		can := res.Expect()
		if !can {
			return chain.Result{}.Failed(authz.NewNotAuthorizedError(acl))
		}

		return next(flo)
	}, 302)
}
