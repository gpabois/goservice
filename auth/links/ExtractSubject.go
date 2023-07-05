package auth_links

import (
	"errors"
	"reflect"

	"github.com/golang-jwt/jwt/v5"
	auth_flow "github.com/gpabois/goservice/auth/flow"
	auth_services "github.com/gpabois/goservice/auth/services"
	"github.com/gpabois/goservice/chain"
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/reflectutil"
)

type ExtractSubjectArgs struct {
	Type reflect.Type
	Name option.Option[string]
}

func NewExtractSubjectArgs[Subject any](args ExtractSubjectArgs) ExtractSubjectArgs {
	args.Type = reflectutil.TypeOf[Subject]()
	return args
}

// Extract the subject from the authentication product
func ExtractSubject(args ExtractSubjectArgs) chain.Link {
	name := args.Name.UnwrapOr(func() string { return "0" })
	return chain.ByFunc(func(flo flow.Flow, next chain.NextFunc) chain.Result {
		prodOpt := auth_flow.GetProduct(flo, name)

		if prodOpt.IsNone() {
			return next(flo)
		}

		prod := prodOpt.Expect()

		// The product type is the same as the subject type, we don't do anything else.
		if reflect.TypeOf(prod) == args.Type {
			flo = auth_flow.SetSubject(flo, prod, name)
			return next(flo)
		}

		subject := reflect.New(args.Type)
		switch p := prod.(type) {
		case jwt.MapClaims:
			res := auth_services.ExtractSubjectFromClaims(p, subject.Elem().Interface())
			if res.HasFailed() {
				return chain.Result{}.Failed(res.UnwrapError())
			}
			auth_flow.SetSubject(flo, subject.Elem().Interface(), name)
		}

		return chain.Result{}.Failed(errors.New("Unknown authentication product"))
	}, 202)
}
