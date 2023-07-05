package authz_services

import (
	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/result"
)

type ACL struct {
	Subject    any
	Permission string
	Object     option.Option[any]
}

type IAuthorizationService interface {
	Can(acl ACL) result.Result[bool]
}
