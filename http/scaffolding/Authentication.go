package http_scaffolding

import (
	http_links "github.com/gpabois/goservice/http/links"
	http_modules "github.com/gpabois/goservice/http/modules"
	utils "github.com/gpabois/goservice/utils"
	"github.com/gpabois/gostd/option"
)

func EnableAuthentication(options ...utils.Configurator[http_links.GetAuthenticationStrategyArgs]) utils.Configurator[http_modules.HttpModuleArgs] {
	return func(a *http_modules.HttpModuleArgs) {
		args := http_links.GetAuthenticationStrategyArgs{}
		utils.Configure(&args, options)
		a.EnableAuthentication = option.Some(args)
	}
}

type AuthenticationScaffolding struct {
	http_links.GetAuthenticationStrategyArgs
}

func WithAuthenticationStrategy(strategy http_links.GetAuthenticationStrategyFunc) utils.Configurator[http_links.GetAuthenticationStrategyArgs] {
	return func(s *http_links.GetAuthenticationStrategyArgs) {
		s.Strategy = strategy
	}
}

func WitHeaderBasedAuthentication(options ...utils.Configurator[http_links.GetAuthenticationStrategyByHeaderArgs]) utils.Configurator[http_links.GetAuthenticationStrategyArgs] {
	args := http_links.GetAuthenticationStrategyByHeaderArgs{}
	utils.Configure(&args, options)
	strategy := http_links.GetAuthenticationStrategyByHeader(args)
	return WithAuthenticationStrategy(strategy)
}

func WithAuthenticationHeaderName(name string) utils.Configurator[http_links.GetAuthenticationStrategyByHeaderArgs] {
	return func(args *http_links.GetAuthenticationStrategyByHeaderArgs) {
		args.Header = option.Some(name)
	}
}

func WithAuthenticationName(name string) utils.Configurator[http_links.GetAuthenticationStrategyArgs] {
	return func(s *http_links.GetAuthenticationStrategyArgs) {
		s.Name = option.Some(name)
	}
}
