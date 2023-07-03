package http_middlewares

import (
	"github.com/gpabois/goservice/auth"
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/goservice/http_transport"
	"github.com/gpabois/goservice/middlewares"
	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/result"
)

const (
	BearerStrategy = iota
)

type GetAuthenticationStrategyArgs struct {
	Header option.Option[string]
	Name   option.Option[string]
}

// Get the authentication strategy from the http request
func GetAuthenticationStrategy(args GetAuthenticationStrategyArgs) middlewares.FlowMiddleware {
	headerName := args.Header.UnwrapOr(func() string { return "Authorization" })
	name := args.Name.UnwrapOr(func() string { return "0" })

	return middlewares.ByFunc(func(in flow.Flow) result.Result[flow.Flow] {
		r := http_transport.Flow_GetHttpRequest(in)

		var header string
		if header = r.Header.Get(headerName); header == "" {
			return result.Success(in)
		}

		getAuthStratRes := auth.GetAuthenticationStrategy(header)
		if getAuthStratRes.HasFailed() {
			return result.Result[flow.Flow]{}.Failed(getAuthStratRes.UnwrapError())
		}

		in = auth.Flow_SetAuthenticationStrategy(in, getAuthStratRes.Expect(), name)
		return result.Success(in)
	})
}
