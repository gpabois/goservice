package auth_modules

import (
	auth_links "github.com/gpabois/goservice/auth/links"
	"github.com/gpabois/goservice/chain"
	"github.com/gpabois/gostd/option"
)

type SubjectArgs struct {
	auth_links.ExtractSubjectArgs
	Inject option.Option[auth_links.InjectSubjectArgs]
}

type subjectModule struct {
	SubjectArgs
}

func newSubjectModule(args SubjectArgs) chain.Module {
	return subjectModule{SubjectArgs: args}
}

func (mod subjectModule) Install(ch chain.Chain) chain.Chain {
	ch = ch.Link(auth_links.ExtractSubject(mod.ExtractSubjectArgs))

	// Inject the subject into the endpoint request
	if mod.Inject.IsSome() {
		args := mod.Inject.Expect()
		ch = ch.Link(auth_links.Reflect_InjectSubject(args))
	}

	return ch
}
