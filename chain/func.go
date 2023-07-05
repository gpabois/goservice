package chain

import (
	"github.com/gpabois/goservice/flow"
)

type LinkFunc func(flo flow.Flow, next NextFunc) Result

func ByFunc(fn LinkFunc, order int) Link {
	return funcLink{fn, order}
}

type funcLink struct {
	fn    func(flo flow.Flow, next NextFunc) Result
	order int
}

func (link funcLink) Call(flo flow.Flow, next NextFunc) Result {
	return link.fn(flo, next)
}

func (link funcLink) Order() int {
	return link.order
}
