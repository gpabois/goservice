package pg

import (
	"context"

	"github.com/gpabois/gostd/result"
	"github.com/jackc/pgx/v5"
)

// A postgres transaction
type Transaction struct {
	Inner   pgx.Tx
	Managed bool
}

func (pg Transaction) IsManaged() bool {
	return pg.Managed
}

func (pg Transaction) Commit() result.Result[bool] {
	err := pg.Inner.Commit(context.Background())
	if err != nil {
		return result.Failed[bool](err)
	}

	return result.Success(true)
}

func (pg Transaction) Rollback() result.Result[bool] {
	err := pg.Inner.Rollback(context.Background())
	if err != nil {
		return result.Failed[bool](err)
	}

	return result.Success(true)
}
