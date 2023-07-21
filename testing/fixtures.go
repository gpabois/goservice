package tests

import (
	"github.com/gpabois/goservice/chain"
	"github.com/gpabois/goservice/flow"
)

// Create a panicking link
func NewPanickingLink(err error, order int) chain.Link {
	return chain.ByFunc(func(flo flow.Flow, next chain.NextFunc) chain.Result {
		panic(err)
	}, order)
}
