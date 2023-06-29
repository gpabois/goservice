package endpoint

import (
	"context"

	"github.com/gpabois/gostd/result"
)

//go:generate mockery
type Endpoint[Request any, Response any] interface {
	Process(ctx context.Context, request Request) result.Result[Response]
}
