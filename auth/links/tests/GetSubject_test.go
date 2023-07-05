package auth_links_tests

import (
	"testing"

	auth_flow "github.com/gpabois/goservice/auth/flow"
	auth_links "github.com/gpabois/goservice/auth/links"
	"github.com/gpabois/goservice/chain"
	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/gostd/option"
	"github.com/stretchr/testify/assert"
)

func Test_GetSubject_Success(t *testing.T) {
	expectedValue := subject{Id: 10}

	flo := flow.Flow{}
	flo = auth_flow.SetProduct(flo, expectedValue, "0")

	ch := chain.NewChain().Link(
		auth_links.ExtractSubject(
			auth_links.NewExtractSubjectArgs[subject](
				auth_links.ExtractSubjectArgs{
					Name: option.Some("0"),
				},
			),
		),
	)

	res := ch.Call(flo)
	assert.True(t, res.IsSuccess(), res.UnwrapError())

	subjectOpt := auth_flow.GetSubject[subject](flo, "0")
	assert.True(t, subjectOpt.IsSome())
	assert.Equal(t, expectedValue, subjectOpt.Expect())

}
