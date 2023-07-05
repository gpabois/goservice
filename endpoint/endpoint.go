package endpoint

import (
	"github.com/gpabois/gostd/result"
)

//go:generate mockery
type Endpoint[Request any, Response any] interface {
	Process(request Request) result.Result[Response]
}
