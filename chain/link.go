package chain

import (
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/gostd/result"
)

type Link interface {
	Call(flo flow.Flow, next NextFunc) result.Result[flow.Flow]
	Order() int
}
