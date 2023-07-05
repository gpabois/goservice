package monitoring_links

import (
	"errors"
	"time"

	"github.com/gpabois/goservice/chain"
	"github.com/gpabois/goservice/flow"
	monitoring_services "github.com/gpabois/goservice/monitoring/services"
	"github.com/gpabois/gostd/result"
)

type MonitorArgs struct {
	Name    string
	Measure bool
	Service monitoring_services.IMonitoringService
}

// Monitor the lifecycle of the process
// Order: 0
func Monitor(args MonitorArgs) chain.Link {
	return chain.ByFunc(func(flo flow.Flow, next chain.NextFunc) chain.Result {
		var ret *result.Result[flow.Flow]

		// Monitor any panic
		defer catchPanic(args, ret)

		if args.Measure {
			defer measure(args, time.Now())
		}

		res := next(flo)
		if res.HasFailed() {
			args.Service.ReportError(args.Name, res.UnwrapError())
		}

		*ret = result.Success(res.Expect())

		return *ret
	}, 0)
}

func measure(args MonitorArgs, t0 time.Time) {
	dt := time.Now().Sub(t0)
	args.Service.ReportMeasurement(args.Name, dt)
}

func catchPanic(args MonitorArgs, ret *chain.Result) {
	if r := recover(); r != nil {
		var err error
		switch v := r.(type) {
		case error:
			err = v
		case string:
			err = errors.New(v)
		}
		args.Service.ReportError(args.Name, err)
		*ret = result.Failed[flow.Flow](err)
	}
}
