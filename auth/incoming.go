package auth

import (
	"fmt"

	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/gostd/option"
)

const subjectPattern = "Subjet.%s"

func Flow_SetSubject(in flow.Flow, subject any, name string) flow.Flow {
	in[fmt.Sprintf(subjectPattern, name)] = subject
	return in
}

func Flow_GetSubject[Subject any](in flow.Flow, name string) option.Option[Subject] {
	return flow.Lookup[Subject](fmt.Sprintf(subjectPattern, name), in)
}
