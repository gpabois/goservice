package auth

import "fmt"

type UnsupportedAuthentication struct{}

func (err UnsupportedAuthentication) Error() string {
	return "unsupported authentication strategy"
}

func NewUnsupportedAuthentication() error {
	return UnsupportedAuthentication{}
}

type UnexpectedAuthenticationStrategyError struct {
	Expected string
	Got      string
}

func NewUnexpectedAuthenticationStrategy(expected string, got string) error {
	return UnexpectedAuthenticationStrategyError{
		Expected: expected,
		Got:      got,
	}
}

func (err UnexpectedAuthenticationStrategyError) Error() string {
	return fmt.Sprintf("unexpected authentication strategy, expecting %s, got %s", err.Expected, err.Got)
}

type FailedAuthenticationError struct {
	inner error
}

func NewFailedAuthenticationError(inner error) error {
	return FailedAuthenticationError{inner}
}

func (err FailedAuthenticationError) Error() string {
	return fmt.Sprintf("authentication failed")
}
