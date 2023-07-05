package auth_links_tests

import "github.com/gpabois/gostd/option"

type subject struct {
	Id int
}

type endpointRequest struct {
	Subject option.Option[subject] `serde:"subject"`
	Value   bool
}
