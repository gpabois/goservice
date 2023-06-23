package bridge

import "github.com/gpabois/gostd/option"

type Bridge = map[string]any

func Lookup[T any, B ~Bridge](key string, bridge B) option.Option[*T] {
	anyValue, ok := bridge[key]

	if !ok {
		return option.None[*T]()
	}

}
