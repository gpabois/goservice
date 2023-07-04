package middlewares

import (
	"sort"

	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/gostd/result"
)

const LastOrder = int(^uint(0) >> 1)
const FirstOrder = -LastOrder - 1

type Middleware interface {
	Intercept(flow flow.Flow) result.Result[flow.Flow]
	Order() int
}

type Middlewares []Middleware

func (ms Middlewares) Add(m Middleware) Middlewares {
	ms = append(ms, m)
	sort.SliceStable(ms, func(i, j int) bool { return ms[i].Order() < ms[j].Order() })
	return ms
}

func (ms Middlewares) Intercept(f flow.Flow) result.Result[flow.Flow] {
	for _, m := range ms {
		res := m.Intercept(f)
		if res.HasFailed() {
			return result.Result[flow.Flow]{}.Failed(res.UnwrapError())
		}

		f = res.Expect()
	}

	return result.Success(f)
}

type IO struct {
	Incoming  Middlewares
	Outcoming Middlewares
}

type Plugin interface {
	Install(io *IO)
}
