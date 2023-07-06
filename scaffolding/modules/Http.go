package modules_scaffolding

import (
	http_links "github.com/gpabois/goservice/http/links"
	http_modules "github.com/gpabois/goservice/http/modules"
	"github.com/gpabois/goservice/utils"
	"github.com/gpabois/gostd/option"
)

func Http(options ...utils.Configurator[http_modules.HttpArgs]) http_modules.Http {
	args := http_modules.HttpArgs{}
	utils.Configure(&args, options)
	return http_modules.NewHttpModule(args)
}

func EnableBodyDeserialization(options ...utils.Configurator[http_links.ReflectDeserializeBodyArgs]) utils.Configurator[http_modules.HttpArgs] {
	return func(a *http_modules.HttpArgs) {
		args := http_links.ReflectDeserializeBodyArgs{}
		utils.Configure(&args, options)
		a.EnableDeserializeBody = option.Some(args)
	}
}

func WithContentTypeFallback(fallback string) utils.Configurator[http_links.ReflectDeserializeBodyArgs] {
	return func(s *http_links.ReflectDeserializeBodyArgs) {
		s.FallbackContentType = option.Some(fallback)
	}
}

func EnableRouteParamInjection() utils.Configurator[http_modules.HttpArgs] {
	return func(args *http_modules.HttpArgs) {
		args.EnableInjectRouteParams = true
	}
}

func ExtractAuthenticationStrategy(name string, strategy http_links.GetAuthenticationStrategyFunc, options ...utils.Configurator[http_links.GetAuthenticationStrategyArgs]) utils.Configurator[http_modules.HttpArgs] {
	return func(a *http_modules.HttpArgs) {
		args := http_links.GetAuthenticationStrategyArgs{}
		args.Name = option.Some(name)
		args.Strategy = strategy
		utils.Configure(&args, options)
		a.ExtractAuthenticationStrategies = append(a.ExtractAuthenticationStrategies, args)
	}
}

func WitHeaderBasedAuthentication(options ...utils.Configurator[http_links.GetAuthenticationStrategyByHeaderArgs]) http_links.GetAuthenticationStrategyFunc {
	args := http_links.GetAuthenticationStrategyByHeaderArgs{}
	utils.Configure(&args, options)
	return http_links.GetAuthenticationStrategyByHeader(args)
}

func WithAuthenticationHeaderName(name string) utils.Configurator[http_links.GetAuthenticationStrategyByHeaderArgs] {
	return func(args *http_links.GetAuthenticationStrategyByHeaderArgs) {
		args.Header = option.Some(name)
	}
}
