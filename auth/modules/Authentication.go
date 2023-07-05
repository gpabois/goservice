package auth_modules

import (
	auth_links "github.com/gpabois/goservice/auth/links"
	"github.com/gpabois/goservice/chain"
	"github.com/gpabois/gostd/option"
)

type AuthenticationArgs struct {
	Authenticate  auth_links.AuthenticateArgs
	EnableSubject option.Option[SubjectArgs]
}

type AuthenticationModule struct {
	AuthenticationArgs
	EnableSubject option.Option[SubjectArgs]
}

func NewAuthenticationModule(args AuthenticationArgs) chain.Module {
	return AuthenticationModule{AuthenticationArgs: args}
}

func (mod AuthenticationModule) Install(ch chain.Chain) chain.Chain {
	ch = ch.Link(auth_links.Authenticate(mod.Authenticate))

	// Generate the subject
	if mod.EnableSubject.IsSome() {
		args := mod.EnableSubject.Expect()
		ch = ch.Install(newSubjectModule(args))
	}

	return ch
}
