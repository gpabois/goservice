package authz_flow

import (
	"fmt"

	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/gostd/option"
)

const objectPattern = "Authz.%s.Object"

func SetObject(in flow.Flow, subject any, name string) flow.Flow {
	in[fmt.Sprintf(objectPattern, name)] = subject
	return in
}

func GetObject[Object any](in flow.Flow, name string) option.Option[Object] {
	return flow.Lookup[Object](fmt.Sprintf(objectPattern, name), in)
}
