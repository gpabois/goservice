package endpoint

import (
	"context"

	"github.com/gpabois/gostd/result"
)

type Endpoint[Request any, Response any] interface {
	Process(ctx context.Context, request Request) result.Result[Response]
}
