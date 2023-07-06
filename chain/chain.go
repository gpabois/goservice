package chain

import (
	"sort"

	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/gostd/result"
)

type NextFunc func(flo flow.Flow) result.Result[flow.Flow]
type Result = result.Result[flow.Flow]

type Chain []Link

type Module interface {
	Install(chain Chain) Chain
}

func NewChain() Chain {
	return Chain{}
}

func (rem Chain) Call(flo flow.Flow) result.Result[flow.Flow] {
	if len(rem) == 0 {
		return result.Success(flo)
	}

	next := func(flo flow.Flow) result.Result[flow.Flow] {
		return rem[1:].Call(flo)
	}

	return rem[0].Call(flo, next)
}

func (chain Chain) Link(lnk Link) Chain {
	chain = append(chain, lnk)
	sort.SliceStable(chain, func(i, j int) bool { return chain[i].Order() < chain[j].Order() })
	return chain
}

func (chain Chain) Install(mods ...Module) Chain {
	for _, mod := range mods {
		chain = mod.Install(chain)
	}
	return chain
}
