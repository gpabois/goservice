package http_scaffolding

import (
	http_links "github.com/gpabois/goservice/http/links"
	http_modules "github.com/gpabois/goservice/http/modules"
	"github.com/gpabois/goservice/utils"
	"github.com/gpabois/gostd/option"
)

func ScaffoldHttpModule(options ...utils.Configurator[http_modules.HttpModuleArgs]) http_modules.HttpModule {
	args := http_modules.HttpModuleArgs{}
	utils.Configure(&args, options)
	return http_modules.NewHttpModule(args)
}

func EnableBodyDeserialization(options ...utils.Configurator[http_links.ReflectDeserializeBodyArgs]) utils.Configurator[http_modules.HttpModuleArgs] {
	return func(a *http_modules.HttpModuleArgs) {
		args := http_links.ReflectDeserializeBodyArgs{}
		utils.Configure(&args, options)
		a.EnableDeserializeBody = option.Some(args)
	}
}

func EnableRouteParamInjection() utils.Configurator[http_modules.HttpModuleArgs] {
	return func(args *http_modules.HttpModuleArgs) {
		args.EnableInjectRouteParams = true
	}
}
