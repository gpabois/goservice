package pg

import (
	"context"

	"github.com/gpabois/goservice/utils"
	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/result"
	pgx "github.com/jackc/pgx/v5"
)

type MigratorArgs struct{}
type Migrator struct {
	conn    *pgx.Conn
	current option.Option[string]
}

type Migration struct {
	ID        string // ID of the migration
	Upgrade   string // Query to upgrade the database
	Downgrade string // Query to downgrade the database
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

func NewMigrator(conn *pgx.Conn, namespace string, options ...utils.Configurator[MigratorArgs]) result.Result[*Migrator] {
	mig := &Migrator{conn: conn}

	// Setup migration table, if it does not exist
	if res := mig.createMigrationTableIfNot(); res.HasFailed() {
		return result.Result[*Migrator]{}.Failed(res.UnwrapError())
	}

	// Retrieve current migration
	currentRes := mig.getCurrentMigration(namespace)
	if currentRes.HasFailed() {
		return result.Result[*Migrator]{}.Failed(currentRes.UnwrapError())
	}
	mig.current = currentRes.Expect()

	return result.Success(mig)
}
