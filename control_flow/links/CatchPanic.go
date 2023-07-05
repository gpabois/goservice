package control_flow_links

import (
	"errors"

	"github.com/gpabois/goservice/chain"
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/gostd/result"
)

// Recover from panic and returns a FailedResult
// Order : 2
func CatchPanic() chain.Link {
	return chain.ByFunc(func(flo flow.Flow, next chain.NextFunc) (ret chain.Result) {
		defer catchPanic(&ret)
		ret = next(flo)
		return ret
	}, 2)
}

func catchPanic(ret *chain.Result) {
	if r := recover(); r != nil {
		var err error
		switch v := r.(type) {
		case error:
			err = v
		case string:
			err = errors.New(v)
		}
		*ret = result.Failed[flow.Flow](err)
	}
}
