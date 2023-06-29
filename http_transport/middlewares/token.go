package middlewares

import (
	"regexp"

	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/goservice/http_transport"
	"github.com/gpabois/goservice/middlewares"
	"github.com/gpabois/goservice/token"
	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/result"
)

const (
	BearerStrategy = iota
)

type TokenRetrieveArgs struct {
	HeaderName option.Option[string]
	Strategy   option.Option[int]
	Name       option.Option[string]
}

// Retrieve raw token from http request
func RetrieveRawToken(args TokenRetrieveArgs) middlewares.FlowMiddleware {
	bearerRegex := regexp.MustCompile("^Bearer[ ]+(?P<raw_token>[[:alnum:]]|.)+$")
	headerName := args.HeaderName.UnwrapOr(func() string { return "Authorization" })
	name := args.Name.UnwrapOr(func() string { return "0" })
	args.Strategy.UnwrapOr(func() int { return BearerStrategy })

	return middlewares.ByFunc(func(in flow.Flow) result.Result[flow.Flow] {
		r := http_transport.Flow_GetHttpRequest(in)

		var header string

		if header = r.Header.Get(headerName); header == "" {
			return result.Success(in)
		}

		// Only strategy is Bearer
		if m := bearerRegex.FindStringSubmatch(header); len(m) == 2 {
			rawToken := m[1]
			in = token.Flow_SetRawToken(in, rawToken, name)
		}

		return result.Success(in)
	})
}
