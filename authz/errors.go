package authz

import authz_services "github.com/gpabois/goservice/authz/services"

type NotAuthenticatedError struct{}

func (err NotAuthenticatedError) Error() string {
	return "not authenticated"
}

func NewNotAuthenticatedError() error {
	return NotAuthenticatedError{}
}

type NotAuthorizedError struct{}

func (err NotAuthorizedError) Error() string {
	return "not authorized"
}

func NewNotAuthorizedError(acl authz_services.ACL) error {
	return NotAuthorizedError{}
}
