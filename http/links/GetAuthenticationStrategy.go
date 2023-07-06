package http_links

import (
	"net/http"

	"github.com/gpabois/goservice/auth"
	auth_flow "github.com/gpabois/goservice/auth/flow"
	"github.com/gpabois/goservice/chain"
	"github.com/gpabois/goservice/flow"
	http_flow "github.com/gpabois/goservice/http/flow"
	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/result"
)

const (
	BearerStrategy = iota
)

type GetAuthenticationStrategyArgs struct {
	Strategy GetAuthenticationStrategyFunc
	Name     option.Option[string]
}

type GetAuthenticationStrategyFunc = func(r *http.Request) result.Result[option.Option[auth.AuthenticationStrategy]]

type GetAuthenticationStrategyByHeaderArgs struct {
	Header option.Option[string]
}

const DefaultAuthenticationHeader = "Authorization"

func GetAuthenticationStrategyByHeader(args GetAuthenticationStrategyByHeaderArgs) GetAuthenticationStrategyFunc {
	return func(r *http.Request) result.Result[option.Option[auth.AuthenticationStrategy]] {
		var header string
		// Nothing to see...
		if header = r.Header.Get(args.Header.UnwrapOr(func() string { return DefaultAuthenticationHeader })); header == "" {
			return result.Success(option.None[auth.AuthenticationStrategy]())
		}

		getAuthStratRes := auth.GetAuthenticationStrategy(header)
		if getAuthStratRes.HasFailed() {
			return result.Failed[option.Option[auth.AuthenticationStrategy]](getAuthStratRes.UnwrapError())
		}

		getAuthStrat := getAuthStratRes.Expect()
		return result.Success(option.Some(getAuthStrat))
	}
}

// Get the authentication strategy from the http request
// Order 200
func GetAuthenticationStrategy(args GetAuthenticationStrategyArgs) chain.Link {
	name := args.Name.UnwrapOr(func() string { return "0" })

	return chain.ByFunc(func(flo flow.Flow, next chain.NextFunc) chain.Result {
		r := http_flow.GetHttpRequest(flo)

		res := args.Strategy(r)
		if res.HasFailed() {
			return chain.Result{}.Failed(res.UnwrapError())
		}

		opt := res.Expect()
		if opt.IsSome() {
			flo = auth_flow.SetAuthenticationStrategy(flo, opt.Expect(), name)
		}

		return next(flo)
	}, 200)
}
