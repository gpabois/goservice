package auth_middlewares

import (
	"errors"

	"github.com/gpabois/goservice/auth"
	"github.com/gpabois/goservice/endpoint"
	endpoint_flow "github.com/gpabois/goservice/endpoint/flow"
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/goservice/middlewares"
	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/result"
)

type InjectSubjectArgs struct {
	FieldName string
	Name      option.Option[string]
}

// Inject authenticated subject into the endpoint request
func InjectAuthentication[EndpointRequest any, Subject any](args InjectSubjectArgs) middlewares.Middleware {
	name := args.Name.UnwrapOr(func() string { return "0" })
	return middlewares.ByFunc(func(in flow.Flow) result.Result[flow.Flow] {
		endpointRequestRes := endpoint_flow.Flow_GetEndpointRequest[EndpointRequest](in)
		if endpointRequestRes.IsNone() {
			return result.Failed[flow.Flow](errors.New("missing endpoint request"))
		}
		endpointRequest := endpointRequestRes.Expect()

		subjectOpt := auth.Flow_GetSubject[Subject](in, name)
		if subjectOpt.IsNone() {
			return result.Success(in)
		}

		subject := subjectOpt.Expect()
		if res := endpoint.Inject(&endpointRequest, map[string]any{args.FieldName: subject}); res.HasFailed() {
			return result.Result[flow.Flow]{}.Failed(res.UnwrapError())
		}

		in = endpoint_flow.Flow_SetEndpointRequest(in, endpointRequest)
		return result.Success(in)
	}, 102)
}
