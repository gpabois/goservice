package monitoring_modules

import (
	"github.com/gpabois/goservice/chain"
	monitoring_links "github.com/gpabois/goservice/monitoring/links"
)

type MonitoringModuleArgs struct {
	monitoring_links.MonitorArgs
}

type MonitoringModule struct {
	MonitoringModuleArgs
}

func NewMonitoringModule(args MonitoringModuleArgs) chain.Module {
	return MonitoringModule{
		MonitoringModuleArgs: args,
	}
}

func (mod MonitoringModule) Install(ch chain.Chain) chain.Chain {
	return ch.Link(monitoring_links.Monitor(mod.MonitorArgs))
}
