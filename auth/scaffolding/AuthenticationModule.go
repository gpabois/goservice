package auth_scaffolding

import (
	auth_links "github.com/gpabois/goservice/auth/links"
	auth_modules "github.com/gpabois/goservice/auth/modules"
	auth_services "github.com/gpabois/goservice/auth/services"
	"github.com/gpabois/goservice/utils"
	"github.com/gpabois/gostd/option"
)

func ScaffoldAuthentication(name string, service auth_services.IAuthenticationService, options ...utils.Configurator[auth_modules.AuthenticationArgs]) auth_modules.AuthenticationModule {
	args := auth_modules.AuthenticationArgs{}
	args.Authenticate.Name = option.Some(name)
	args.Authenticate.Service = service
	utils.Configure(&args, options)
	return auth_modules.NewAuthenticationModule(args)
}

func EnableSubject(options ...utils.Configurator[auth_modules.SubjectArgs]) utils.Configurator[auth_modules.AuthenticationArgs] {
	return func(s *auth_modules.AuthenticationArgs) {
		args := auth_modules.SubjectArgs{}
		args.Name = s.Authenticate.Name
		utils.Configure(&args, options)
		s.EnableSubject = option.Some(args)
	}
}

func EnableSubjectInjection(fieldName string) utils.Configurator[auth_modules.SubjectArgs] {
	return func(s *auth_modules.SubjectArgs) {
		args := auth_links.InjectSubjectArgs{}
		args.FieldName = fieldName
		args.Name = s.Name
		s.EnableInjection = option.Some(args)
	}
}
