package pg

import (
	"context"

	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/result"
	pgx "github.com/jackc/pgx/v5"
)

type MigratorArgs struct{}

type Migrator struct {
	conn *pgx.Conn
}

func (mig *Migrator) createMigrationTableIfNot() result.Result[bool] {
	sql := `
		CREATE TABLE IF NOT EXISTS goservice_migrations (
			namespace VARCHAR(255),
			current VARCHAR(255)
		)
	`

	_, err := mig.conn.Exec(context.Background(), sql)
	if err != nil {
		return result.Result[bool]{}.Failed(err)
	}
	return result.Success(true)
}

func (mig *Migrator) getCurrentMigration(namespace string) result.Result[option.Option[string]] {
	sql := `
		SELECT current FROM goservice_migrations WHERE namespace=$1
	`
	rows, err := mig.conn.Query(context.Background(), sql, namespace)

	if err != nil {
		return result.Result[option.Option[string]]{}.Failed(err)
	}

	if !rows.Next() {
		return result.Success(option.None[string]())
	}

	var current string
	Scan(rows, &current)
	return result.Success(option.Some(current))
}

func (mig *Migrator) Setup(namespace string, options ...configurator.Configurator[MigratorArgs]) result.Result[bool] {
	var res result.Result[bool]

	if res = mig.createMigrationTableIfNot(); res.HasFailed() {
		return res
	}

	currentRes := mig.getCurrentMigration(namespace)
}
