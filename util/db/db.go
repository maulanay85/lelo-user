package db

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func (u *UtilDBModule) BeginTransaction(ctx context.Context, opt ...pgx.TxOptions) (pgx.Tx, error) {
	var tx pgx.Tx
	var err error

	if len(opt) > 0 {
		tx, err = u.Db.BeginTx(ctx, opt[0])
	} else {
		tx, err = u.Db.BeginTx(ctx, pgx.TxOptions{})
	}

	if err != nil {
		return nil, err
	}
	return tx, err
}
