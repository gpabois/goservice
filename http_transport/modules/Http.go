package http_modules

import (
	auth_modules "github.com/gpabois/goservice/auth/modules"
	"github.com/gpabois/goservice/chain"
	control_flow_modules "github.com/gpabois/goservice/control_flow/modules"
	http_links "github.com/gpabois/goservice/http_transport/links"
	monitoring_modules "github.com/gpabois/goservice/monitoring/modules"
	"github.com/gpabois/gostd/option"
)

type AuthenticationArgs struct {
	http_links.GetAuthenticationStrategyArgs
	auth_modules.AuthenticationArgs
}

type HttpModuleArgs struct {
	DeserializeBody      option.Option[http_links.ReflectDeserializeBodyArgs]
	InjectRouteParams    bool
	EnableMonitoring     option.Option[monitoring_modules.MonitoringModuleArgs]
	EnableAuthentication option.Option[AuthenticationArgs]
}

func NewHttpModule(args HttpModuleArgs) chain.Module {
	return HttpModule{
		HttpModuleArgs: args,
	}
}

// Install http module, and control flow module
type HttpModule struct {
	HttpModuleArgs
}

func (mod HttpModule) Install(ch chain.Chain) chain.Chain {
	ch = ch.
		Install(control_flow_modules.ControlFlowModule{}).
		Link(http_links.Lifecycle())

	if mod.DeserializeBody.IsSome() {
		args := mod.DeserializeBody.Expect()
		ch = ch.Link(http_links.Reflect_DeserializedBody(args))
	}

	if mod.InjectRouteParams {
		ch = ch.Link(http_links.Reflect_InjectRouteParams())
	}

	if mod.EnableMonitoring.IsSome() {
		args := mod.EnableMonitoring.Expect()
		ch = ch.Install(monitoring_modules.NewMonitoringModule(args))
	}

	if mod.EnableAuthentication.IsSome() {
		args := mod.EnableAuthentication.Expect()
		ch = ch.Link(http_links.GetAuthenticationStrategy(args.GetAuthenticationStrategyArgs))
		ch = ch.Install(auth_modules.NewAuthenticationModule(args.AuthenticationArgs))
	}

	return ch
}
