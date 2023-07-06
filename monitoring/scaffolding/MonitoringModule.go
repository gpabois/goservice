package monitoring_scaffolding

import (
	monitoring_modules "github.com/gpabois/goservice/monitoring/modules"
	monitoring_services "github.com/gpabois/goservice/monitoring/services"
	"github.com/gpabois/goservice/utils"
)

func ScaffoldMonitoring(name string, service monitoring_services.IMonitoringService, options ...utils.Configurator[monitoring_modules.MonitoringModuleArgs]) monitoring_modules.MonitoringModule {
	args := monitoring_modules.MonitoringModuleArgs{}
	args.Name = name
	args.Service = service
	utils.Configure(&args, options)
	return monitoring_modules.NewMonitoringModule(args)
}
