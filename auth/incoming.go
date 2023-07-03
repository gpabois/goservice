package auth

import (
	"fmt"

	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/gostd/option"
)

const strategyPattern = "Auth.%s.Strategy"
const subjectPattern = "Auth.%s.Subject"

func Flow_SetAuthenticationStrategy(in flow.Flow, strategy AuthenticationStrategy, name string) flow.Flow {
	in[fmt.Sprintf(strategyPattern, name)] = strategy
	return in
}

func Flow_GetAuthenticationStrategy(in flow.Flow, name string) option.Option[AuthenticationStrategy] {
	return flow.Lookup[AuthenticationStrategy](fmt.Sprintf(strategyPattern, name), in)

}

func Flow_SetSubject(in flow.Flow, subject any, name string) flow.Flow {
	in[fmt.Sprintf(subjectPattern, name)] = subject
	return in
}

func Flow_GetSubject[Subject any](in flow.Flow, name string) option.Option[Subject] {
	return flow.Lookup[Subject](fmt.Sprintf(subjectPattern, name), in)
}
