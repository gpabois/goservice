package flow

import "github.com/gpabois/gostd/option"

type Flow map[string]any

func Lookup[T any](key string, flow Flow) option.Option[T] {
	anyValue, ok := flow[key]

	if !ok {
		return option.None[T]()
	}

	val, ok := anyValue.(T)

	if !ok {
		return option.None[T]()
	}

	return option.Some(val)
}
