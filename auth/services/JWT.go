package auth_services

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/gpabois/goservice/auth"
	"github.com/gpabois/gostd/result"
	"github.com/gpabois/gostd/serde/decoder"
	"github.com/gpabois/gostd/serde/norm"
)

type JWTArgs struct {
	KeyFunc jwt.Keyfunc
}

// Handles JWT-based authentication
type JWT[Subject any] struct {
	JWTArgs
}

func NewJWT[Subject any](args JWTArgs) IAuthenticationService[Subject] {
	return &JWT[Subject]{
		JWTArgs: args,
	}
}

func (s *JWT[Subject]) Authenticate(strategy auth.AuthenticationStrategy) result.Result[Subject] {
	var subject Subject
	// Only support Bearer auth scheme
	if strategy.Scheme != auth.Bearer {
		return result.Failed[Subject](auth.NewUnexpectedAuthenticationStrategy(auth.Bearer, strategy.Scheme))
	}

	rawToken := strategy.Credentials

	tok, err := jwt.ParseWithClaims(rawToken, jwt.MapClaims{}, s.KeyFunc)
	if err != nil {
		return result.Result[Subject]{}.Failed(auth.NewFailedAuthenticationError(err))
	}

	// Decode claims map into the claims
	mapClaims := tok.Claims.(jwt.MapClaims)
	d := norm.NewDecoder(mapClaims)

	res := decoder.DecodeInto(d, &subject)
	if res.HasFailed() {
		return result.Result[Subject]{}.Failed(res.UnwrapError())
	}

	return result.Success(subject)
}
