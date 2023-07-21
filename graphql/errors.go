package graphql

import "errors"

func NewMissingRequest() error {
	return errors.New("missing graphql request")
}

func NewExecutionErrors(errs ...error) error {
	return errors.Join(errs...)
}
