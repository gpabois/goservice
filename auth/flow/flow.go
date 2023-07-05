package auth

import (
	"fmt"

	"github.com/gpabois/goservice/auth"
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/gostd/option"
)

const strategyPattern = "Auth.%s.Strategy"
const subjectPattern = "Auth.%s.Subject"

func SetAuthenticationStrategy(in flow.Flow, strategy auth.AuthenticationStrategy, name string) flow.Flow {
	in[fmt.Sprintf(strategyPattern, name)] = strategy
	return in
}

func GetAuthenticationStrategy(in flow.Flow, name string) option.Option[auth.AuthenticationStrategy] {
	return flow.Lookup[auth.AuthenticationStrategy](fmt.Sprintf(strategyPattern, name), in)

}

func SetSubject(in flow.Flow, subject any, name string) flow.Flow {
	in[fmt.Sprintf(subjectPattern, name)] = subject
	return in
}

func GetSubject[Subject any](in flow.Flow, name string) option.Option[Subject] {
	return flow.Lookup[Subject](fmt.Sprintf(subjectPattern, name), in)
}
