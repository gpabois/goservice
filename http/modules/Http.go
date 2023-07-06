package http_modules

import (
	"github.com/gpabois/goservice/chain"
	control_flow_modules "github.com/gpabois/goservice/control_flow/modules"
	http_links "github.com/gpabois/goservice/http/links"
	"github.com/gpabois/gostd/option"
)

type HttpArgs struct {
	EnableDeserializeBody           option.Option[http_links.ReflectDeserializeBodyArgs]
	EnableInjectRouteParams         bool
	ExtractAuthenticationStrategies []http_links.GetAuthenticationStrategyArgs
}

func NewHttpModule(args HttpArgs) Http {
	return Http{
		HttpArgs: args,
	}
}

// Install http module, and control flow module
type Http struct {
	HttpArgs
}

func (mod Http) Install(ch chain.Chain) chain.Chain {
	ch = ch.
		Install(control_flow_modules.ControlFlowModule{}).
		Link(http_links.Lifecycle())

	if mod.EnableDeserializeBody.IsSome() {
		args := mod.EnableDeserializeBody.Expect()
		ch = ch.Link(http_links.Reflect_DeserializedBody(args))
	}

	if mod.EnableInjectRouteParams {
		ch = ch.Link(http_links.Reflect_InjectRouteParams())
	}

	for _, args := range mod.ExtractAuthenticationStrategies {
		ch = ch.Link(http_links.GetAuthenticationStrategy(args))
	}

	return ch
}
