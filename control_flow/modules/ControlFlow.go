package control_flow_modules

import (
	"github.com/gpabois/goservice/chain"
	control_flow_links "github.com/gpabois/goservice/control_flow/links"
)

func NewControlFlowModule() ControlFlowModule {
	return ControlFlowModule{}
}

type ControlFlowModule struct{}

func (ctr ControlFlowModule) Install(ch chain.Chain) chain.Chain {
	return ch.Link(control_flow_links.CatchPanic())
}
