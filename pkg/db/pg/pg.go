package pg

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/kirillmc/platform_common/pkg/db"
	"github.com/kirillmc/platform_common/pkg/db/prettier"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type pg struct {
	dbc *pgxpool.Pool
}

func NewDB(dbc *pgxpool.Pool) db.DB {
	return &pg{
		dbc: dbc,
	}
}

func (p *pg) ScanOneContext(ctx context.Context, dest interface{}, q db.Query, logger zap.Logger, args ...interface{}) error {
	logQuery(ctx, q, logger, args...)

	row, err := p.QueryContext(ctx, q, logger, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanOne(dest, row)
}

func (p *pg) ScanAllContext(ctx context.Context, dest interface{}, q db.Query, logger zap.Logger, args ...interface{}) error {
	logQuery(ctx, q, logger, args...)

	rows, err := p.QueryContext(ctx, q, logger, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanAll(dest, rows)
}

func (p *pg) ExecContext(ctx context.Context, q db.Query, logger zap.Logger, args ...interface{}) (pgconn.CommandTag, error) {
	logQuery(ctx, q, logger, args...)

	return p.dbc.Exec(ctx, q.QueryRaw, args...)
}

func (p *pg) QueryContext(ctx context.Context, q db.Query, logger zap.Logger, args ...interface{}) (pgx.Rows, error) {
	logQuery(ctx, q, logger, args...)

	return p.dbc.Query(ctx, q.QueryRaw, args...)
}

func (p *pg) QueryRowContext(ctx context.Context, q db.Query, logger zap.Logger, args ...interface{}) pgx.Row {
	logQuery(ctx, q, logger, args...)

	return p.dbc.QueryRow(ctx, q.QueryRaw, args...)
}

func (p *pg) Ping(ctx context.Context) error {
	return p.dbc.Ping(ctx)
}

func (p pg) Close() {
	p.dbc.Close()
}

func logQuery(ctx context.Context, q db.Query, logger zap.Logger, args ...interface{}) {
	prettyQuery := prettier.Pretty(q.QueryRaw, prettier.PlaceholderDollar, args...)
	logger.Sugar().Debugln(ctx,
		fmt.Sprintf("sql: %s", q.Name),
		fmt.Sprintf("query: %s", prettyQuery))

}
