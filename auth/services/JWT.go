package auth_services

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/gpabois/goservice/auth"
	"github.com/gpabois/gostd/result"
)

type JWTArgs[Subject any, Claims jwt.Claims] struct {
	ExtractSubject func(claims Claims) Subject
	KeyFunc        jwt.Keyfunc
}

// Handles JWT-based authentication
type JWT[Subject any, Claims jwt.Claims] struct {
	JWTArgs[Subject, Claims]
}

func NewJWT[Subject any, Claims jwt.Claims](args JWTArgs[Subject, Claims]) IAuthenticationService[Subject] {
	return &JWT[Subject, Claims]{
		JWTArgs: args,
	}
}

func (s *JWT[Subject, Claims]) Authenticate(strategy auth.AuthenticationStrategy) result.Result[Subject] {
	var claims Claims

	// Only support Bearer auth scheme
	if strategy.Scheme != auth.Bearer {
		return result.Failed[Subject](auth.NewUnexpectedAuthenticationStrategy(auth.Bearer, strategy.Scheme))
	}

	rawToken := strategy.Credentials
	tokenResult := result.From(jwt.ParseWithClaims(rawToken, claims, s.KeyFunc))
	if tokenResult.HasFailed() {
		return result.Result[Subject]{}.Failed(auth.NewFailedAuthenticationError(tokenResult.UnwrapError()))
	}

	tok := tokenResult.Expect()
	claims = tok.Claims.(Claims)
	return result.Success(s.ExtractSubject(claims))
}
