package db

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UtilDBModule struct {
	Db *pgxpool.Pool
}

func NewUtilDb(db *pgxpool.Pool) *UtilDBModule {
	return &UtilDBModule{
		Db: db,
	}
}

type UtilDb interface {
	BeginTransaction(ctx context.Context, opt ...pgx.TxOptions) (pgx.Tx, error)
}
