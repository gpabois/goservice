package auth

import (
	"regexp"

	"github.com/gpabois/gostd/result"
)

type AuthenticationStrategy struct {
	Scheme      string
	Credentials string
}

const Bearer = "Bearer"
const Basic = "Basic"
const Digest = "Digest"

func NewBearer(token string) AuthenticationStrategy {
	return AuthenticationStrategy{
		Scheme:      Bearer,
		Credentials: token,
	}
}

// Get the authentication mode (Bearer <token>, Basic <base64 encoded credentials>, ....)
func GetAuthenticationStrategy(raw string) result.Result[AuthenticationStrategy] {
	authzRegex := regexp.MustCompile("^(<?P<Scheme>Basic|Bearer|Digest|DPoP|HOBA|Mutual|Negotiate|OAuth|SCRAM-SHA-1|SCRAM-SHA-256|vapid)[ ]+(?P<Credentials>[-A-Za-z0-9_|.]+)$")
	if m := authzRegex.FindStringSubmatch(raw); len(m) == 3 {
		return result.Success(AuthenticationStrategy{
			Scheme:      m[1],
			Credentials: m[2],
		})
	}

	return result.Result[AuthenticationStrategy]{}.Failed(NewUnsupportedAuthentication())
}
