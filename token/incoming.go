package token

import (
	"fmt"

	"github.com/gpabois/goservice/flow"
	"github.com/gpabois/gostd/option"
)

const rawTokenName = "Token.%s.Raw"
const tokenName = "Token.%s"

func Flow_SetToken[Token any](in flow.Flow, token Token, name string) flow.Flow {
	in[fmt.Sprintf(tokenName, name)] = token
	return in
}

func Flow_GetToken[Token any](in flow.Flow, token Token, name string) option.Option[Token] {
	return flow.Lookup[Token](fmt.Sprintf(tokenName, name), in)

}
func Flow_SetRawToken(in flow.Flow, token string, name string) flow.Flow {
	in[fmt.Sprintf(rawTokenName, name)] = token
	return in
}

func Flow_GetRawToken(in flow.Flow, name string) option.Option[string] {
	return flow.Lookup[string](fmt.Sprintf(rawTokenName, name), in)
}
