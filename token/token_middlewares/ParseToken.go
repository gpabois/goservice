package token_middlewares

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/goservice/middlewares"
	"github.com/gpabois/goservice/token"
	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/result"
)

type ParseTokenArgs struct {
	Name option.Option[string] // Name of the token
	Key  any                   // Key to validate the jwt
}

// Parse the token
func ParseWithClaimsToken[Claims jwt.Claims](args ParseTokenArgs) middlewares.FlowMiddleware {
	name := args.Name.UnwrapOr(func() string { return "0" })
	key := args.Key

	return middlewares.ByFunc(func(in flow.Flow) result.Result[flow.Flow] {
		var claims Claims
		rawTokenOpt := token.Flow_GetRawToken(in, name)
		// No token, we don't do anything
		if rawTokenOpt.IsNone() {
			return result.Success(in)
		}

		rawToken := rawTokenOpt.Expect()

		tokenResult := result.From(jwt.ParseWithClaims(rawToken, claims, func(tok *jwt.Token) (any, error) {
			return key, nil
		}))

		if tokenResult.HasFailed() {
			return result.Result[flow.Flow]{}.Failed(tokenResult.UnwrapError())
		}

		tok := tokenResult.Expect()
		in = token.Flow_SetToken(in, tok, name)

		return result.Success(in)
	})
}
