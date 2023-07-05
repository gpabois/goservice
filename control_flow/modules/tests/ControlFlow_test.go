package control_flow_modules_tests

import (
	"errors"
	"testing"

	"github.com/gpabois/goservice/chain"
	control_flow_modules "github.com/gpabois/goservice/control_flow/modules"
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/goservice/tests"
	"github.com/stretchr/testify/assert"
)

func Test_ControlFlowModule_Panic(t *testing.T) {
	expectedError := errors.New("panic!")
	ch := chain.NewChain().
		Install(control_flow_modules.NewControlFlowModule()). // Control Flow should recover from any panic, and returns a Failed result.
		Link(tests.NewPanickingLink(expectedError, 1000))     // Chain a panicking link

	res := ch.Call(flow.Flow{})
	assert.True(t, res.HasFailed(), "should have failed")
	assert.Equal(t, expectedError, res.UnwrapError())
}
