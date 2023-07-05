package monitoring_services

import "time"

type IMonitoringService interface {
	ReportError(name string, err error)
	ReportSuccess(name string)
	ReportMeasurement(name string, duration time.Duration)
}
