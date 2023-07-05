package auth_links

import (
	"errors"
	"reflect"

	auth_flow "github.com/gpabois/goservice/auth/flow"
	"github.com/gpabois/goservice/chain"
	"github.com/gpabois/goservice/endpoint"
	endpoint_flow "github.com/gpabois/goservice/endpoint/flow"
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/result"
)

type InjectSubjectArgs struct {
	FieldName string
	Name      option.Option[string]
}

// Inject authenticated subject into the endpoint request
func Reflect_InjectSubject(args InjectSubjectArgs) chain.Link {
	name := args.Name.UnwrapOr(func() string { return "0" })
	return chain.ByFunc(func(flo flow.Flow, next chain.NextFunc) chain.Result {
		subjectOpt := auth_flow.GetSubject[any](flo, name)
		if subjectOpt.IsNone() {
			return next(flo)
		}

		anySubject := subjectOpt.Expect()
		anyReq := endpoint_flow.GetEndpointRequest[any](flo)
		if anyReq.IsNone() {
			return result.Failed[flow.Flow](errors.New("missing endpoint request"))
		}
		req := reflect.New(reflect.TypeOf(anyReq.Expect()))
		req.Elem().Set(reflect.ValueOf(anyReq.Expect()))

		mp := map[string]any{args.FieldName: anySubject}
		if res := endpoint.Reflect_Inject(req.Interface(), mp); res.HasFailed() {
			return result.Result[flow.Flow]{}.Failed(res.UnwrapError())
		}

		flo = endpoint_flow.SetEndpointRequest(flo, req.Elem().Interface())
		return next(flo)
	}, 203)
}
