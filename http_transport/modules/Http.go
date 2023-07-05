package http_modules

import (
	"github.com/gpabois/goservice/chain"
	control_flow_modules "github.com/gpabois/goservice/control_flow/modules"
	http_links "github.com/gpabois/goservice/http_transport/links"
	"github.com/gpabois/gostd/option"
)

type HttpModuleArgs struct {
	DeserializeBody   bool
	InjectRouteParams bool
	EnableMonitoring  option.Option[monitoring_modules.MonitoringModuleArgs]
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

	if mod.DeserializeBody {
		ch = ch.Link(http_links.DeserializeBody())
	}

	if mod.InjectRouteParams {
		ch = ch.Link(http_links.InjectRouteParams())
	}

	if mod.Monitor {
		ch = ch.Install(monitoring_modules.NewMonitoringModule())
	}

	return ch
}
