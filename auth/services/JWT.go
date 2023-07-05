package auth_services

import (
	"errors"

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
type JWT struct {
	JWTArgs
}

func NewJWT(args JWTArgs) IAuthenticationService {
	return &JWT{JWTArgs: args}
}

func ExtractSubjectFromClaims(claims any, subject any) result.Result[bool] {
	mapClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		return result.Result[bool]{}.Failed(errors.New("not claims"))
	}
	d := norm.NewDecoder(mapClaims)
	return decoder.Reflect_DecodeInto(d, subject)
}

func (s *JWT) Authenticate(strategy auth.AuthenticationStrategy) result.Result[any] {
	// Only support Bearer auth scheme
	if strategy.Scheme != auth.Bearer {
		return result.Failed[any](auth.NewUnexpectedAuthenticationStrategy(auth.Bearer, strategy.Scheme))
	}

	rawToken := strategy.Credentials

	tok, err := jwt.ParseWithClaims(rawToken, jwt.MapClaims{}, s.KeyFunc)
	if err != nil {
		return result.Result[any]{}.Failed(auth.NewFailedAuthenticationError(err))
	}

	mapClaims := tok.Claims.(jwt.MapClaims)

	return result.Success[any](mapClaims)
}
